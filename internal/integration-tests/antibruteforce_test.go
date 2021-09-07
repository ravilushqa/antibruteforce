package tests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cucumber/godog"
	apipb "github.com/ravilushqa/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"google.golang.org/grpc"
)

var grpcService = os.Getenv("GRPC_SERVICE")

func init() {
	if grpcService == "" {
		grpcService = "localhost:50051"
	}
}

func (a *apiTest) iCallGrpcMethod(method string) (err error) {
	cc, err := grpc.Dial(grpcService, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %v", err)
	}
	defer cc.Close()

	c := apipb.NewAntiBruteforceServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	switch method {
	case "Check":
		_, err = c.Check(ctx, &apipb.CheckRequest{
			Login:    a.login,
			Password: a.password,
			Ip:       a.ip,
		})
		a.responseError = err
	case "Reset":
		_, err = c.Reset(ctx, &apipb.ResetRequest{
			Login: a.login,
			Ip:    a.ip,
		})
		a.responseError = err
	case "BlacklistAdd":
		_, err = c.BlacklistAdd(ctx, &apipb.BlacklistAddRequest{
			Subnet: a.subnet,
		})
		a.responseError = err
	case "BlacklistRemove":
		_, err = c.BlacklistRemove(ctx, &apipb.BlacklistRemoveRequest{
			Subnet: a.subnet,
		})
		a.responseError = err
	case "WhitelistAdd":
		_, err = c.WhitelistAdd(ctx, &apipb.WhitelistAddRequest{
			Subnet: a.subnet,
		})
		a.responseError = err
	case "WhitelistRemove":
		_, err = c.WhitelistRemove(ctx, &apipb.WhitelistRemoveRequest{
			Subnet: a.subnet,
		})
		a.responseError = err
	default:
		return errors.New("unexpected method: " + method)
	}

	return nil
}

func (a *apiTest) responseErrorShouldBe(error string) error {
	if error != "nil" {
		error = "rpc error: code = Unknown desc = " + error
	}
	if error == "nil" && a.responseError != nil {
		return fmt.Errorf("unexpected error, expected %s, got %v", error, a.responseError)
	}
	if error != "nil" && a.responseError == nil {
		return fmt.Errorf("unexpected error, expected %s, got %v", error, nil)
	}
	if a.responseError != nil && error != a.responseError.Error() {
		return fmt.Errorf("unexpected error, expected %s, got %v", error, a.responseError.Error())
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	var test apiTest
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		test.login = ""
		test.password = ""
		test.ip = ""
		test.responseError = nil

		return ctx, nil
	})

	ctx.Step(`^login is "([^"]*)"$`, test.loginIs)
	ctx.Step(`^password is "([^"]*)"$`, test.passwordIs)
	ctx.Step(`^ip is "([^"]*)"$`, test.ipIs)
	ctx.Step(`^subnet is "([^"]*)"$`, test.subnetIs)

	ctx.Step(`^I call grpc method "([^"]*)"$`, test.iCallGrpcMethod)
	ctx.Step(`^response error should be "([^"]*)"$`, test.responseErrorShouldBe)
}
