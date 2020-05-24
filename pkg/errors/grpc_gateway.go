package errors

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorBody struct {
	RawError string `protobuf:"varint,100,name=error" json:"-"`
	Error    bool   `protobuf:"bool,101,name=error" json:"error,omitempty"`
	// This is to make the error more compatible with users that expect errors to be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	code int32  `json:"-"`
	Code string `protobuf:"bytes,2,name=code" json:"code"`
	// Key     string     `protobuf:"bytes,3,name=key" json:"key"`
	Message string     `protobuf:"bytes,4,name=message" json:"message"`
	Details []*any.Any `protobuf:"bytes,5,rep,name=details" json:"details,omitempty"`
}

// Clean removes the RPC specific style of handling error messages
// reducing them to be only the original message
func (e *errorBody) Clean() {
	if strings.Contains(e.RawError, " = ") {
		c, s := e.RawError, ""
		parts := strings.Split(c, " = ")
		if len(parts) > 2 {
			s = strings.Replace(parts[1], " desc", "", 1)
			c = parts[2]
		}
		if s != "" && e.code == 2 {
			switch s {
			case "Unauthenticated":
				e.code = 16
			}
		}

		e.Code = c
	}

	e.Message = humanizeCode(e.Code)

}

// HTTPError handles cleaning up a GRPC error to better reflect an expected REST response
// Additionally it prevents exposing that the service is actually a GRPC based server
func HTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "Unable to successfully marshal error message"}`
	fmt.Println(grpc.ErrorDesc(err))

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	contentType := marshaler.ContentType()
	// Check marshaler on run time in order to keep backwards compatability
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	body := &errorBody{
		Error:   true,
		Code:    s.Message(),
		Message: s.Message(),
		code:    int32(s.Code()),
		Details: s.Proto().GetDetails(),
	}

	body.Clean()

	buf, merr := marshaler.Marshal(body)
	if merr != nil {
		fmt.Printf("Failed to marshal error message %q: %v", body, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			fmt.Sprintf("Failed to write response: %v", err)
		}
		return
	}

	// md, ok := runtime.ServerMetadataFromContext(ctx)
	// if !ok {
	// 	fmt.Sprintf("Failed to extract ServerMetadata from context")
	// }

	// handleForwardResponseServerMetadata(w, mux, md)
	// handleForwardResponseTrailerHeader(w, md)
	st := runtime.HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		fmt.Sprintf("Failed to write response: %v", err)
	}
}

// Cheap way to sort of humanize a dot or underscore separated message
func humanizeCode(code string) string {

	result := strings.ReplaceAll(code, ".", " ")
	result = strings.ReplaceAll(result, "_", " ")

	return result
}
