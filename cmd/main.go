package main

import (
	"errors"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	server "pvz-service/internal/app"
	"pvz-service/internal/config"
	"pvz-service/pkg/logger"
	"syscall"

	"github.com/joho/godotenv"
)

var buildmodes = map[string]struct{}{
	"dev":   {},
	"local": {},
}

func main() {
	buildmode := processFlags()
	if _, ok := buildmodes[*buildmode]; !ok {
		log.Fatalln("invalid buildmode provided")
	}

	configDir := filepath.Join("..", "configs")

	cfgPath := filepath.Join(configDir, *buildmode+".yaml")
	envPath := filepath.Join(configDir, ".env")

	godotenv.Load(envPath)

	cfg, err := config.New(cfgPath)
	reportOnError(err)

	logger, err := logger.New(os.Stdout, cfg.LogLevel)
	reportOnError(err)

	server, err := server.New(cfg, logger)
	reportOnError(err)

	go func() {
		err = server.Run()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("server error: %v", err)
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	slog.Info("starting server shutdown")
	if err := server.Shutdown(); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}
	slog.Info("server shutdown completed")
}

func reportOnError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func processFlags() *string {
	buildmode := flag.String("mode", "dev", "mode for application build")

	flag.Parse()

	return buildmode
}
