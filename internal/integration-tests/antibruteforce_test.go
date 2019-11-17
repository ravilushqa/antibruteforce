// +build integration

package integration_tests

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	apipb "gitlab.com/otus_golang/antibruteforce/internal/antibruteforce/delivery/grpc/api"
	"google.golang.org/grpc"
	"os"
	"time"
)

var grpcService = os.Getenv("GRPC_SERVICE")

func init() {
	if grpcService == "" {
		grpcService = "localhost:50051"
	}
}

type apiTest struct {
	login         string
	password      string
	ip            string
	responseError error
}

func (a *apiTest) loginIs(login string) error {
	a.login = login
	return nil
}

func (a *apiTest) passwordIs(pass string) error {
	a.password = pass
	return nil
}

func (a *apiTest) ipIs(ip string) error {
	a.ip = ip
	return nil
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

func FeatureContext(s *godog.Suite) {
	var test apiTest
	s.BeforeScenario(func(i interface{}) {
		test.login = ""
		test.password = ""
		test.ip = ""
		test.responseError = nil
	})
	s.Step(`^login is "([^"]*)"$`, test.loginIs)
	s.Step(`^password is "([^"]*)"$`, test.passwordIs)
	s.Step(`^ip is "([^"]*)"$`, test.ipIs)

	s.Step(`^I call grpc method "([^"]*)"$`, test.iCallGrpcMethod)
	s.Step(`^response error should be "([^"]*)"$`, test.responseErrorShouldBe)
}
