package main

import (
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/microgateway/activity/jwt/examples"
)

func main() {
	e, err := examples.Example()
	if err != nil {
		panic(err)
	}
	engine.RunEngine(e)
}
