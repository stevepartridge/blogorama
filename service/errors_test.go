package service

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func TestHandleErrorRequestIsNilAsInvalidArgument(t *testing.T) {
	code := codeFromError(ErrRequestIsNil)
	if codes.InvalidArgument != code {
		t.Errorf("Expected code %v but recieved %v", codes.InvalidArgument, code)
	}
}
