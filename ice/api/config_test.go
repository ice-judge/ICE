package main

import (
	"os"
	"testing"

	"github.com/ice-judge/ICE/ice/api/server"
	"github.com/stretchr/testify/assert"
)

func prepareConfigTest() {
	os.Unsetenv("ADDR")
	os.Unsetenv("ASSETS_PATH")
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("GRACEFUL_TIMEOUT")
}

func TestParseDefaultConfig(t *testing.T) {
	prepareConfigTest()
	a := assert.New(t)

	config := parseConfig()
	solution := server.Config{
		Addr:            ":8080",
		AssetsPath:      "/dist",
		SecretKey:       "secret",
		GracefulTimeout: 10,
	}
	a.Equal(solution, config)
}

func TestParseEmptyConfig(t *testing.T) {
	prepareConfigTest()
	a := assert.New(t)

	os.Setenv("ADDR", "")
	os.Setenv("ASSETS_PATH", "")
	os.Setenv("SECRET_KEY", "")
	os.Setenv("GRACEFUL_TIMEOUT", "0")

	config := parseConfig()
	solution := server.Config{
		Addr:            "",
		AssetsPath:      "",
		SecretKey:       "",
		GracefulTimeout: 0,
	}
	a.Equal(solution, config)
}

func TestParseGracefulTimeout(t *testing.T) {
	prepareConfigTest()
	a := assert.New(t)

	os.Setenv("GRACEFUL_TIMEOUT", "5353")

	config := parseConfig()
	a.Equal(5353, int(config.GracefulTimeout))
}

func TestParseInvalidGracefulTimeout(t *testing.T) {
	prepareConfigTest()
	a := assert.New(t)

	os.Setenv("GRACEFUL_TIMEOUT", "")

	defer func() {
		if r := recover(); r == nil {
			a.Fail("parsing invaild graceful timeout should cause panic")
		}
	}()

	parseConfig()
}

func TestParseInvalidGracefulTimeout2(t *testing.T) {
	prepareConfigTest()
	a := assert.New(t)

	os.Setenv("GRACEFUL_TIMEOUT", "asdfsadf")

	defer func() {
		if r := recover(); r == nil {
			a.Fail("parsing invaild graceful timeout should cause panic")
		}
	}()

	parseConfig()
}
