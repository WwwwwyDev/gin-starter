package main

import (
	"gin-starter/internal/initialize"
	"gin-starter/pkg/core"
)

func init() {
	core.Setup("dev")
}

func main() {
	initialize.Run()
}
