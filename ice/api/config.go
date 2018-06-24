package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ice-judge/ICE/ice/api/server"
	"github.com/ice-judge/ICE/ice/pkg/utils"
)

func parseConfig() server.Config {
	var (
		Addr             = utils.Getenv("ADDR", ":8080")
		AssetsPath       = utils.Getenv("ASSETS_PATH", "/dist")
		SecretKey        = utils.Getenv("SECRET_KEY", "secret")
		GracefulTimeout_ = utils.Getenv("GRACEFUL_TIMEOUT", "10")
	)

	GracefulTimeout, err := strconv.Atoi(GracefulTimeout_)
	if err != nil {
		panic(fmt.Sprintf("Ivalid GRACEFUL_TIMEOUT: %s", GracefulTimeout_))
	}

	return server.Config{
		Addr:            Addr,
		AssetsPath:      AssetsPath,
		SecretKey:       SecretKey,
		GracefulTimeout: time.Duration(GracefulTimeout),
	}
}
