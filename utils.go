package baiduMap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"errors"
)

func GetDefaultAK() string {
	ak := defaultAppKey // baidu's
	if value, has := os.LookupEnv("GOBAIDUMAP_AK"); has {
		ak = value
	}
	return ak
}

func requestBaidu(reqType, reqURL string) (interface{}, error) {

	res, err := getResStruct(reqType)
	if err != nil {
		return res, err
	}
	httpClient := http.Client{}
	resp, err := httpClient.Get(reqURL)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {

		err := json.Unmarshal(bytes, &res)

		if err != nil {
			return res, err
		}

	} else {
		return res, errors.New("error")
	}

	return res, nil
}

func getResStruct(reqType string) (interface{}, error) {
	var res interface{}
	if reqType == "GetAddressViaGEO" {
		return new(StructGEOToAddress), nil
	}
	if reqType == "GetRoute" {
		return new(StructRoute), nil
	}
	return res, errors.New("error for struct")
}
