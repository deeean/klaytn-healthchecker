package main

import (
	"fmt"
	"github.com/deeean/klaytn-healthchecker/rpc"
	"github.com/deeean/klaytn-healthchecker/util"
	"github.com/gin-gonic/gin"
	"log"
)

var RPC_URL string
var MAX_BLOCK_DIFFERENCE int64

func healthz(c *gin.Context) {
	diff, err := rpc.GetSyncing(RPC_URL)
	if diff == -1 || diff > MAX_BLOCK_DIFFERENCE || err != nil {
		fmt.Println(diff)
		c.String(503, "unhealthy")
		return
	}

	c.String(200, "ok")
}

func main() {
	rpcUrl := util.GetString("RPC_URL")
	if rpcUrl == "" {
		rpcUrl = "http://localhost:8551"
	}

	maxBlockDifference, err := util.GetInt("MAX_BLOCK_DIFFERENCE")
	if err != nil {
		maxBlockDifference = 3
	}

	RPC_URL = rpcUrl
	MAX_BLOCK_DIFFERENCE = maxBlockDifference

	r := gin.Default()
	r.GET("/healthz", healthz)

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
