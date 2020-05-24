package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	blog "github.com/stevepartridge/blogorama"
)

var (
	defaultTestHost = "https://localhost"
	defaultTestPort = 12345
	defaultBaseURL  = fmt.Sprintf("%s:%d", defaultTestHost, defaultTestPort)

	svc *Service

	testClient *http.Client
)

func init() {

	var err error

	testClient = http.DefaultClient
	testClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	blog.SetDatastore(blog.DatastoreConfig{
		Type: blog.DatastoreTypeMySQLMock,
	})

	svc, err = New(defaultTestPort)
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	go svc.Serve()
}
