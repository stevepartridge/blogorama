package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ory/viper"
	"github.com/stevepartridge/blogorama/pkg/log"
	blog "github.com/stevepartridge/blogorama/service"
)

var (
	serviceName = "blog"

	version = "0.0.0"
	builtAt = ""
	build   = "dev"
	githash = ""

	defaultPort = 8000
	port        int

	service *blog.Service
)

// setup the log and version/build info for the service
func init() {
	log.Setup(fmt.Sprintf("%s:%s:%s", serviceName, version, build))

	if build == "dev" {
		build = "-1"
		builtAt = time.Now().UTC().Format(time.RFC3339Nano)
	}

	if builtAt == "" {
		builtAt = time.Now().Format(time.RFC3339Nano)
	}

	blog.Name = serviceName
	blog.Version = version
	blog.Build = build
	blog.BuiltAt = builtAt
	blog.GitHash = githash
}

func main() {

	defer shutdown()

	prepare()

	// Listen to Unix signals SIGINT, SIGTERM so we can attempt to exit and clean up gracefully
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		s := <-sig
		fmt.Println("Received signal", s.String())
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to serve")
		shutdown(err)
	}

	shutdown()

}

// serve handles the serving up of the service as well as handling the ideal shutdown scenario gravcefully
func serve(ctx context.Context) error {

	port = viper.GetInt("port")
	if port == 0 {
		port = defaultPort
	}

	// Take the service off the main thread
	go func() {

		var err error

		service, err = blog.New(port) // Initiate our blog service
		if err != nil {
			shutdown(err)
		}

		log.Info().Int("port", port).Msg("Listening")
		if err = service.Serve(); err != nil && err != http.ErrServerClosed { // Serve it up
			log.Error().Err(err).Msg("Failed to serve")
		}
	}()

	<-ctx.Done()

	// Create shutdown context with timeout for exiting regardless of graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := service.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("Service shutdown failed")
		return err
	}
	log.Info().Msg("Service has stopped successfully")

	return nil

}

// Prepare the configuration
func prepare() {

	viper.SetConfigName(serviceName)
	viper.AddConfigPath("/etc/services/")                      // Path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", serviceName)) // Check home dir
	viper.AddConfigPath(".")                                   // Check for config in the working directory
	viper.SetConfigType("yaml")                                // Config file type
	viper.AutomaticEnv()                                       // Leverage the environment as well

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn().Str("notice", err.Error()).Msg("Config file not found, assuming environment holds the knowledge.")
	}

	go viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

		log.Warn().Str("config", e.Name).Msg("Config file changed")

		go reload()
	})

}

// Handle any reload tasks necessary if something changed
func reload() {
	service.Reload()
}

// Shutdown helper to throttle excessive restart iterations
func shutdown(errs ...error) {
	fmt.Printf("Shutting down...")
	if len(errs) > 0 {
		for i := range errs {
			fmt.Printf("\nERROR: %s\n", errs[i].Error())
		}
	}
	time.Sleep(3 * time.Second)
	fmt.Println("bye, felica!")
	os.Exit(1)
}
