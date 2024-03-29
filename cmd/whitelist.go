package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
)

func init() {
	whitelistAdd.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	whitelistAdd.PersistentFlags().StringVarP(&subnet, "subnet", "s", "", "subnet")

	whitelistRemove.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	whitelistRemove.PersistentFlags().StringVarP(&subnet, "subnet", "s", "", "subnet")

	rootCmd.AddCommand(whitelist)
	whitelist.AddCommand(whitelistAdd, whitelistRemove)
}

var whitelist = &cobra.Command{
	Use:   "whitelist",
	Short: "whitelist actions",
	Long:  `whitelist actions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Use antibruteforce whitelist [command].\nRun 'antibruteforce whitelist --help' for usage.\n")
	},
}

var whitelistAdd = &cobra.Command{
	Use:   "add",
	Short: "add to whitelist",
	Long:  `add to whitelist`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := apipb.NewAntiBruteforceServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.WhitelistAdd(ctx, &apipb.WhitelistAddRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not add to whitelist: %v", err)
		}
		log.Printf("Success: %t", r.Ok)
	},
}

var whitelistRemove = &cobra.Command{
	Use:   "remove",
	Short: "remove from whitelist",
	Long:  `remove from whitelist`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := apipb.NewAntiBruteforceServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.WhitelistRemove(ctx, &apipb.WhitelistRemoveRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not remove from whitelist: %v", err)
		}
		log.Printf("Success: %t", r.Ok)
	},
}
