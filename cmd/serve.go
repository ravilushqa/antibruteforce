package cmd

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // GRPC Proxy https://github.com/jnewmano/grpc-json-proxy
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ravilushqa/antibruteforce/config"
	dbInstance "github.com/ravilushqa/antibruteforce/db"
	grpcInstance "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/repository"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/usecase"
	bucketRepository "github.com/ravilushqa/antibruteforce/internal/bucket/repository"
	"github.com/ravilushqa/antibruteforce/logger"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs" // optimization for k8s
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func init() {
	rootCmd.AddCommand(serve)
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start antibruteforce server",
	Long:  `Start grpc antibruteforce server`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.GetConfig()
		if err != nil {
			log.Fatalf("unable to load config: %v", err)
		}

		l, err := logger.GetLogger(c)
		if err != nil {
			log.Fatalf("unable to load logger: %v", err)
		}
		defer func() {
			_ = l.Sync()
		}()

		db, err := dbInstance.GetDb(c)
		if err != nil {
			log.Fatalf("unable to load db: %v", err)
		}

		lis, err := net.Listen("tcp", c.URL)
		if err != nil {
			l.Fatal(fmt.Sprintf("failed to listen %v", err))
		}
		l.Info("server has started at " + c.URL)
		grpcServer := grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_prometheus.StreamServerInterceptor,
				grpc_zap.StreamServerInterceptor(l),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_prometheus.UnaryServerInterceptor,
				grpc_zap.UnaryServerInterceptor(l),
			)),
		)

		if c.IsDevelopment() {
			reflection.Register(grpcServer)
		}

		r := repository.NewPsqlAntibruteforceRepository(db, l)
		br := bucketRepository.NewMemoryBucketRepository(l)
		u := usecase.NewAntibruteforceUsecase(r, br, l, c)
		apipb.RegisterAntiBruteforceServiceServer(grpcServer, grpcInstance.NewServer(u, l))

		// starting monitoring
		grpc_prometheus.Register(grpcServer)
		grpc_prometheus.EnableHandlingTimeHistogram()
		l.Info(fmt.Sprintf("Monitoring export listen %s", c.PrometheusHost))
		go func() {
			err = http.ListenAndServe(c.PrometheusHost, promhttp.Handler())
			if err != nil {
				l.Error(err.Error())
			}
			http.Handle("/metrics", promhttp.Handler())
		}()

		l.Info("grpc server starting")
		// starting service
		if err = grpcServer.Serve(lis); err != nil {
			l.Error(err.Error())
			grpcServer.GracefulStop()
		}
	},
}
