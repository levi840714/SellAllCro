package cdc

import (
	"SellAllCro/config"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

func sign(_req RequestFormat) string {
	paramStr := sortMapAndSplice(_req.Params)
	signStr := fmt.Sprintf("%s%d%s%s%d", _req.Method, _req.Id, _req.ApiKey, paramStr, _req.Nonce)
	h := hmac.New(sha256.New, []byte(config.Config.SecretKey))
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

func milliTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func sortMapAndSplice(_m map[string]interface{}) string {
	var paramStr string
	keys := make([]string, 0, len(_m))

	for field := range _m {
		keys = append(keys, field)
	}
	sort.Strings(keys)

	for _, key := range keys {
		paramStr += key
		paramStr += fmt.Sprintf("%v", _m[key])
	}

	return paramStr
}

func getResponseJson(req *http.Request) (jsonByte []byte, err error) {
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	jsonByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}
