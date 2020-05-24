package service

import (
	"context"
	"testing"

	pb "github.com/stevepartridge/blogorama/protos/go"
)

func TestServerShutdownAndRestart(t *testing.T) {
	var (
		ctx = context.Background()
		err error
	)

	svc.Shutdown(ctx)

	svc, err = New(defaultTestPort)
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

}

func TestServiceInfo(t *testing.T) {

	Version = "testing"

	svc := blogService{}
	resp, err := svc.Info(context.Background(), &pb.InfoRequest{})
	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	if resp.Version != "testing" {
		t.Errorf("Expected version '%s' but saw '%s'", "testing", resp.Version)
	}

}
