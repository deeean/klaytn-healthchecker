package rpc

import (
	"bytes"
	"encoding/json"
	"github.com/deeean/klaytn-healthchecker/util"
	"io/ioutil"
	"net/http"
)

type syncing struct {
	ID      int               `json:"id"`
	Jsonrpc string            `json:"jsonrpc"`
	Result  map[string]string `json:"result"`
}

var getSyncingBody = []byte(`{"id":1,"jsonrpc":"2.0","method":"klay_syncing","params":0}`)

func GetSyncing(url string) (int64, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(getSyncingBody))
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var syncingPayload syncing
	err = json.Unmarshal(buf, &syncingPayload)
	if err != nil {
		return 0, nil
	}

	currentBlock, err := util.HexToInt64(syncingPayload.Result["currentBlock"])
	if err != nil {
		return -1, nil
	}

	highestBlock, err := util.HexToInt64(syncingPayload.Result["highestBlock"])
	if err != nil {
		return -1, nil
	}

	return highestBlock - currentBlock, nil
}
