package cmd

import (
	"context"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"time"
)

func init() {
	reset.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	reset.PersistentFlags().StringVarP(&login, "login", "l", "", "login to reset")
	reset.PersistentFlags().StringVarP(&ip, "ip", "i", "", "ip to reset")

	rootCmd.AddCommand(reset)
}

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Reset bucket",
	Long:  `Reset bucket`,
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
		r, err := c.Reset(ctx, &apipb.ResetRequest{
			Login: login,
			Ip:    ip,
		})
		if err != nil {
			log.Fatalf("could not reset: %v", err)
		}
		log.Printf("Success: %t", r.Ok)
	},
}
