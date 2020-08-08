package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ravilushqa/antibruteforce/config"
	"github.com/ravilushqa/antibruteforce/db"
	grpcinstance "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/repository"
	"github.com/ravilushqa/antibruteforce/internal/antibruteforce/usecase"
	bucketrepository "github.com/ravilushqa/antibruteforce/internal/bucket/repository"
	"github.com/ravilushqa/antibruteforce/logger"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "github.com/jnewmano/grpc-json-proxy/codec" // GRPC Proxy https://github.com/jnewmano/grpc-json-proxy
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs" // optimization for k8s
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

		db, err := db.GetDb(c)
		if err != nil {
			log.Fatalf("unable to load db: %v", err)
		}

		lis, err := net.Listen("tcp", c.URL)
		if err != nil {
			l.Fatal(fmt.Sprintf("failed to listen %v", err))
		}
		l.Info("server has started at " + c.URL)
		grpcServer := grpc.NewServer(
			grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
				grpcprometheus.StreamServerInterceptor,
				grpczap.StreamServerInterceptor(l),
			)),
			grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
				grpcprometheus.UnaryServerInterceptor,
				grpczap.UnaryServerInterceptor(l),
			)),
		)

		if c.IsDevelopment() {
			reflection.Register(grpcServer)
		}

		r := repository.NewPsqlAntibruteforceRepository(db, l.With(zap.String("module", "antibruteforce.repository")))
		br := bucketrepository.NewMemoryBucketRepository(l.With(zap.String("module", "bucket.repository")))
		u := usecase.NewAntibruteforceUsecase(r, br, l.With(zap.String("module", "antibruteforce.usecase")), c)
		apipb.RegisterAntiBruteforceServiceServer(grpcServer, grpcinstance.NewServer(u, l.With(zap.String("module", "grpc"))))

		// starting monitoring
		grpcprometheus.Register(grpcServer)
		grpcprometheus.EnableHandlingTimeHistogram()
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
