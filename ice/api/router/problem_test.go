package router_test

import (
	"testing"

	. "github.com/ice-judge/ICE/ice/api/router"
	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, true)
	//ProblemHandler(nil, nil)
	HelloWorld()
}
