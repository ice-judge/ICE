package utils_test

import (
	"os"
	"testing"

	"github.com/ice-judge/ICE/ice/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetenv(t *testing.T) {
	a := assert.New(t)
	str := utils.RandString(20)

	val := utils.Getenv(str, "default")
	a.Equal("default", val)

	os.Setenv(str, "")

	val = utils.Getenv(str, "default")
	a.Equal("", val)

	os.Setenv(str, "asdf")

	val = utils.Getenv(str, "default")
	a.Equal("asdf", val)
}
