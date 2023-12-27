package hikvision_sdk

import (
	"encoding/json"
	"testing"
)

func TestSDK(t *testing.T) {
	hk := HKConfig{
		Ip:      "192.168.211.250",
		Port:    443,
		AppKey:  "20394016",
		Secret:  "LqOedmdU4hn2Uz9L9dPp",
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
