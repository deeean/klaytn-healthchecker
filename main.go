package main

import (
	"github.com/deeean/klaytn-healthchecker/routes"
	"github.com/deeean/klaytn-healthchecker/utils"
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

	rpcUrl := utils.GetEnvOrDefault("RPC_URL", "http://localhost:8551")
	maxBlockDifference, err := utils.GetEnvInt64OrDefault("MAX_BLOCK_DIFFERENCE", 30)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/healthz", routes.Healthz(rpcUrl, maxBlockDifference))
	e.Logger.Fatal(e.Start(":3000"))
}
