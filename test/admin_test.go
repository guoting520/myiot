package test

import (
	"encoding/json"
	"fmt"
	"myiot/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

func TestDeviceAddr(t *testing.T) {
	m := map[string]string{
		"token": "121212121111111111",
	}
	header, _ := json.Marshal(m)
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
