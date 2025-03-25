package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fevse/banners_rotator/internal/app"
	"github.com/fevse/banners_rotator/internal/config"
	"github.com/fevse/banners_rotator/internal/logger"
	"github.com/fevse/banners_rotator/internal/queue"
	grpcserver "github.com/fevse/banners_rotator/internal/server/grpc"
	sqlstorage "github.com/fevse/banners_rotator/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "path to configuration file")
}

func main() {
	flag.Parse()

	config, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logg := logger.New()
	storage := sqlstorage.New()
	err = storage.Connect(config.DB.DSN)
	if err != nil {
		logg.Error("storage error: " + err.Error())
		os.Exit(1)
	}
	rabbit := queue.New(config.Rabbit.URI, config.Rabbit.Queue, config.Rabbit.Exchange, config.Rabbit.Kind)
	err = rabbit.Connect()
	if err != nil {
		logg.Error("rabbit error: " + err.Error())
		os.Exit(1)
	}

	rotator := app.New(storage, logg, rabbit)

	err = storage.Migrate(config.DB.Migration)
	if err != nil {
		logg.Error("migration error: " + err.Error())
		os.Exit(1)
	}

	grpcserver := grpcserver.NewServer(rotator)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()
		_, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		grpcserver.Stop()
		storage.Close()
		rabbit.Close()
		logg.Info("service stopped")
	}()

	logg.Info("service is running...")

	if err := grpcserver.Start(config.GRPCServer.Network, config.GRPCServer.Address); err != nil {
		logg.Error("failed to stop grpc server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
