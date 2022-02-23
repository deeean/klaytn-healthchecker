package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Syncing struct {
	ID      int               `json:"id"`
	Jsonrpc string            `json:"jsonrpc"`
	Result  map[string]string `json:"result"`
}

var getSyncingBody = []byte(`{"id":1,"jsonrpc":"2.0","method":"klay_syncing","params":0}`)

func GetSyncing(url string) (*Syncing, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(getSyncingBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var syncingPayload Syncing
	if err = json.Unmarshal(buf, &syncingPayload); err != nil {
		return nil, nil
	}

	return &syncingPayload, nil
}
