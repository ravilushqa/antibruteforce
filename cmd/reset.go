package cmd

import (
	"context"
	"github.com/spf13/cobra"
	apipb "gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	address string
	login   string
	ip      string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "localhost:50051", "antibruteforce address")
	rootCmd.PersistentFlags().StringVarP(&login, "login", "l", "", "login to reset")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "ip to reset")

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
