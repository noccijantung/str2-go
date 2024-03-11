package utils

import (
	"bytes"
	"encoding/json"
	"github.com/noccijantung/str2-go/config"
	"net/http"
	"os"
	"strconv"
	"time"
)

type reportBlock struct {
	Time       int64  `json:"time"`
	WorkerName string `json:"worker"`
	IpAddress  string `json:"ipAddress"`
	Miner      string `json:"miner"`
	Wallet     string `json:"wallet"`
	Block      string `json:"block"`
	Bluescore  uint64 `json:"bluescore"`
	Nonce      uint64 `json:"nonce"`
}

func rettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func makefiledump(f string, path string) {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	taskFile, errFile := os.Create(path + timestamp + ".json")
	if errFile != nil {
	}
	defer func(taskFile *os.File) {
		err := taskFile.Close()
		if err != nil {
		}
	}(taskFile)
	_, err := taskFile.WriteString(f)
	if err != nil {
	}
}

func Makepush(worker string, ip string, minerapp string, wallet string, block string, bl uint64, nonce uint64) error {
	vConfig := config.StrConfig

	// 1. preparedata
	statePush := vConfig.PushOnlyFile
	dataTask := reportBlock{Time: time.Now().Unix(), WorkerName: worker, IpAddress: ip, Miner: minerapp, Wallet: wallet, Block: block, Bluescore: bl, Nonce: nonce}
	pushString, errForm := rettyStruct(dataTask)
	if errForm != nil {
	}
	// 2. writedata to API or dump file
	if !statePush {
		jsonBody := []byte(pushString)
		bodyReader := bytes.NewReader(jsonBody)
		req, errReq := http.NewRequest(http.MethodPost, vConfig.ApiUrl, bodyReader)
		if errReq != nil {
			makefiledump(pushString, vConfig.Path)
		}
		req.Header.Set("User-Agent", "PoolCaller-Agent/1.0")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Access-Token", vConfig.ApiToken)
		clientHttp := http.Client{
			Timeout: 1 * time.Second,
		}
		res, errClient := clientHttp.Do(req)
		if errClient != nil {
			makefiledump(pushString, vConfig.Path)
		} else {
			statusCode := res.StatusCode
			if statusCode != 201 {
				makefiledump(pushString, vConfig.Path)
			}
		}
	} else {
		// dump to file for store later
		makefiledump(pushString, vConfig.Path)
	}
	return nil
}
