package handler

import (
	"github.com/deeean/klaytn-healthchecker/rpc"
	"github.com/deeean/klaytn-healthchecker/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Healthz(rpcUrl string, maxBlockDifference int64) echo.HandlerFunc {
	return func(c echo.Context) error {
		syncing, err := rpc.GetSyncing(rpcUrl)
		if err != nil {
			c.Logger().Error(err)
			return c.String(http.StatusServiceUnavailable, "unhealthy")
		}

		if syncing == nil {
			return c.String(http.StatusOK, "healthy")
		}

		currentBlock, err := util.HexToInt64(syncing.Result["currentBlock"])
		if err != nil {
			return c.String(http.StatusServiceUnavailable, "unhealthy")
		}

		highestBlock, err := util.HexToInt64(syncing.Result["highestBlock"])
		if err != nil {
			return c.String(http.StatusServiceUnavailable, "unhealthy")
		}

		if highestBlock-currentBlock > maxBlockDifference {
			c.Logger().Error(highestBlock - currentBlock)
			return c.String(http.StatusServiceUnavailable, "unhealthy")
		}

		return c.String(200, "healthy")
	}
}
