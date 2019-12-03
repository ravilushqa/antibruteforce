package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/otus_golang/antibruteforce/config"
	dbInstance "gitlab.com/otus_golang/antibruteforce/db"
	grpcInstance "gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/delivery/grpc"
	apipb "gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/repository"
	"gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/usecase"
	bucketRepository "gitlab.com/otus_golang/antibruteforce/internal/bucket/repository"
	"gitlab.com/otus_golang/antibruteforce/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

		db, err := dbInstance.GetDb(c)
		if err != nil {
			log.Fatalf("unable to load db: %v", err)
		}

		lis, err := net.Listen("tcp", c.URL)
		if err != nil {
			l.Fatal(fmt.Sprintf("failed to listen %v", err))
		}
		l.Info("server has started at " + c.URL)
		grpcServer := grpc.NewServer()

		if c.IsDevelopment() {
			reflection.Register(grpcServer)
		}

		r := repository.NewPsqlAntibruteforceRepository(db, l)
		br := bucketRepository.NewMemoryBucketRepository(l)
		u := usecase.NewAntibruteforceUsecase(r, br, l, c)
		apipb.RegisterAntiBruteforceServiceServer(grpcServer, grpcInstance.NewServer(u, l))

		err = grpcServer.Serve(lis)

		if err != nil {
			l.Fatal(err.Error())
		}
	},
}
