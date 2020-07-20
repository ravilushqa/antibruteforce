package cmd

import (
	"context"
	"fmt"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

func init() {
	blacklistAdd.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	blacklistAdd.PersistentFlags().StringVarP(&subnet, "subnet", "s", "", "subnet")

	blacklistRemove.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	blacklistRemove.PersistentFlags().StringVarP(&subnet, "subnet", "s", "", "subnet")

	rootCmd.AddCommand(blacklist)
	blacklist.AddCommand(blacklistAdd, blacklistRemove)
}

var blacklist = &cobra.Command{
	Use:   "blacklist",
	Short: "blacklist actions",
	Long:  `blacklist actions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Use antibruteforce blacklist [command].\nRun 'antibruteforce blacklist --help' for usage.\n")
	},
}

var blacklistAdd = &cobra.Command{
	Use:   "add",
	Short: "add to blacklist",
	Long:  `add to blacklist`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := apipb.NewAntiBruteforceServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.BlacklistAdd(ctx, &apipb.BlacklistAddRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not add to blacklist: %v", err)
		}
		log.Printf("Success: %t", r.Ok)
	},
}

var blacklistRemove = &cobra.Command{
	Use:   "remove",
	Short: "remove from blacklist",
	Long:  `remove from blacklist`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := apipb.NewAntiBruteforceServiceClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.BlacklistRemove(ctx, &apipb.BlacklistRemoveRequest{
			Subnet: subnet,
		})
		if err != nil {
			log.Fatalf("could not remove from blacklist: %v", err)
		}
		log.Printf("Success: %t", r.Ok)
	},
}
