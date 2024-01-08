package monitor

import (
	"bytes"
	"encoding/json"
	"errors"
	"filscan-delay-monitor/prometh"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Result struct {
		Height    int64 `json:"height"`
		BlockTime int64 `json:"block_time"`
	} `json:"result"`
}

func fetchBlockTimeDiff(api string) (timeDiff int64, err error) {
	var jsonStr = []byte(`{}`)

	req, _ := http.NewRequest("POST", api, bytes.NewBuffer(jsonStr))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 999, errors.New(fmt.Sprintf("finalheight api request fail,err:%s", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 999, errors.New(fmt.Sprintf("finalheight read body fail,err:%s", err))
	}

	var res Response

	err = json.Unmarshal(body, &res)
	if err != nil {
		return 999, errors.New(fmt.Sprintf("finalheight json Unmarshal fail,err:%s", err))
	}

	currentTime := time.Now().Unix()
	timeDiff = currentTime - res.Result.BlockTime

	return timeDiff, nil
}

func RunMonitorLoop(url string, pushAddr string, net string, interval int) {
	for {
		timeDiff, err := fetchBlockTimeDiff(url)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Minute)
			continue
		}
		log.Printf("get %s delay success %d, will push to pushgateway", net, timeDiff)
		prometh.Push(pushAddr, timeDiff, net)

		time.Sleep(time.Duration(interval) * time.Minute)

	}
}
