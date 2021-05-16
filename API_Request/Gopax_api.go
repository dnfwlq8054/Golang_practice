package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	apiKey = "노트북에 있음"
	secret = "노트묵에 있음"
)

func call(
	needAuth bool, method string, path string,
	body map[string]interface{}, // can be nil
	recvWindow int, // set -1 not to assign
) *map[string]interface{} {
	method = strings.ToUpper(method)
	var bodyBytes []byte = nil
	if body != nil {
		bodyBytes, _ = json.Marshal(body)
	}

	req, _ := http.NewRequest(method, "https://api.gopax.co.kr"+path, bytes.NewBuffer(bodyBytes))

	if needAuth {
		timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
		msg := "t" + timestamp + method
		if method == "GET" && strings.HasPrefix(path, "/orders?") {
			msg += path
		} else {
			msg += strings.Split(path, "?")[0]
		}
		if recvWindow != -1 {
			msg += strconv.Itoa(recvWindow)
			req.Header.Set("receive-window", strconv.Itoa(recvWindow))
		}
		if bodyBytes != nil {
			msg += string(bodyBytes)
		}

		rawSecret, _ := base64.StdEncoding.DecodeString(secret)
		mac := hmac.New(sha512.New, rawSecret)
		mac.Write([]byte(msg))
		rawSignature := mac.Sum(nil)
		signature := base64.StdEncoding.EncodeToString(rawSignature)

		req.Header.Set("api-key", apiKey)
		req.Header.Set("timestamp", timestamp)
		req.Header.Set("signature", signature)

	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	respBodyBytes, _ := ioutil.ReadAll(resp.Body)

	return &map[string]interface{}{
		"statusCode": resp.StatusCode,
		"body":       string(respBodyBytes),
		"header":     resp.Header,
	}
}

func main() {
	/*postOrderReqBody := map[string]interface{}{
		"side": "buy", "type": "limit", "amount": 1,
		"price": 10000, "tradingPairName": "BTC-KRW",
	}*/
	//log.Print(*call(true, "POST", "/orders", postOrderReqBody, 200))
	log.Print(*call(true, "GET", "/balances", nil, -1))
	log.Print(*call(true, "GET", "/balances/ETH", nil, -1))
	//log.Print(*call(true, "GET", "/orders?includePast=true", nil, -1))
	//log.Print(*call(true, "GET", "/trades?limit=1", nil, -1))
	//log.Print(*call(false, "GET", "/trading-pairs/BTC-KRW/book?level=1", nil, -1))
}
