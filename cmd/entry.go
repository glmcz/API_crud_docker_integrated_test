package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gopkg.in/yaml.v3"

	"simpleCloudService/cmd/config"
	"simpleCloudService/internal/api"
	"simpleCloudService/internal/repository"
)

func Run(ctx context.Context, configFile string, port int, templatePath string) error {
	// handle manual interrupt
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg := config.NewDefaultConfig(configFile)
	if err := load(configFile, &cfg); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	postgresRepository, err := repository.NewPostgresRepository(&cfg.PostgresConfig)
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	//create an empty table
	err = postgresRepository.EmptyAutoMigrate()
	if err != nil {
		return fmt.Errorf("failed to create init table Users: %w", err)
	}

	layer := api.NewAPI(postgresRepository) // TODO add template

	server := http.Server{
		Addr:    cfg.ServerConfig.Address,
		Handler: layer.Muxer(),
	}

	if port > 0 && templatePath != "" {
		server.Addr = strconv.Itoa(port)
	}

	serverErrors := make(chan error, 1)
	go func() {
		println("Server started and listening on", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.Done():
		// Gracefully shutdown the server
		log.Println("Shutting down server...")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shut down server: %w", err)
		}
		log.Println("Server gracefully stopped.")
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
	}

	return nil
}

func load(configFile string, cfg *config.Config) error {
	fileConfig, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}

	fileConfig = []byte(os.ExpandEnv(string(fileConfig)))

	if err := yaml.Unmarshal(fileConfig, cfg); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}
