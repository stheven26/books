package main

import (
	"runtime"
	"test/internal/interfaces"

	"github.com/labstack/echo/v4"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	e := echo.New()
	interfaces.SetupRouter(e)
	e.Start(":8080")
}
