# hikvision-sdk
hikvision sdk for Golang

Fork from https://github.com/zxbit2011/hikvisionOpenAPIGo

Hikvision OpenAPI Doc: [海康开放平台](https://open.hikvision.com/docs/docId?productId=5c67f1e2f05948198c909700&version=%2Ff95e951cefc54578b523d1738f65f0a1&tagPath=%E5%AF%B9%E6%8E%A5%E6%8C%87%E5%8D%97)

# Quick Start

```shell
go get github.com/YueY4n9/hikvision-sdk
```

# Code Test

```go
package hikvision_sdk

import (
	"encoding/json"
	"testing"
)

func TestSDK(t *testing.T) {
	hk := HKConfig{
		Ip:      "127.0.0.1",
		Port:    443,
		AppKey:  "28057000",
		Secret:  "dZztQSS0000kLpURG000",
		IsHttps: true,
	}

	body := map[string]interface{}{
		"pageNo":   "1",
		"pageSize": "100",
	}
	result, err := hk.HttpPost("/artemis/api/resource/v1/cameras", body, 15)
	if err != nil {
		t.Fatal(err)
		return
	}
	resJson, err := json.Marshal(result)
	t.Log("OK", string(resJson))
}
```