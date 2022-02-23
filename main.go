package main

import (
	"github.com/deeean/klaytn-healthchecker/handler"
	"github.com/deeean/klaytn-healthchecker/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] ${remote_ip} ${status} ${method} ${host}${path} ${latency_human}\n",
		}),
	)

	rpcUrl := util.GetEnvOrDefault("RPC_URL", "http://localhost:8551")
	maxBlockDifference, err := util.GetEnvInt64OrDefault("MAX_BLOCK_DIFFERENCE", 30)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/healthz", handler.Healthz(rpcUrl, maxBlockDifference))
	e.Logger.Fatal(e.Start(":3000"))
}
