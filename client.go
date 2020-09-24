// Created by @menduo @ 2019-01-24
package baiduMap

import (
	"fmt"
	"errors"
)

type BaiduMapClient struct {
	ak string
}

func NewBaiduMapClient(ak string) *BaiduMapClient {
	return &BaiduMapClient{ak: ak}
}

func (bc *BaiduMapClient) GetAk() string {
	return bc.ak
}

func (bc *BaiduMapClient) SetAk(ak string) {
	bc.ak = ak
}

func (bc *BaiduMapClient) GetRoute(lat1, lng1, lat2, lng2 string) (*StructRoute, error) {
	res := new(StructRoute)
	parameter := fmt.Sprintf("origin=%s,%s&destination=%s,%s&ak=%s", lat1, lng1, lat2, lng2, bc.GetAk())
	reqURL := fmt.Sprintf("%s%s", reqURLForRoute, parameter)

	res2, err := requestBaidu("GetRoute", reqURL)
	if err != nil {
		return res, err
	}
	if res2.(*StructRoute).Status != 0 {
		message := fmt.Sprintf("百度 API 报错：%s", res2.(*StructRoute).Message)
		return res, errors.New(message)
	}
	res3 := res2.(*StructRoute)
	return res3, nil
}

func (bc *BaiduMapClient) GetAddressViaGEO(lat, lng string) (*StructGEOToAddress, error) {
	res := new(StructGEOToAddress)

	parameter := fmt.Sprintf("%s&output=json&coordtype=wgs84ll&location=%s,%s", bc.GetAk(), lat, lng)
	reqURL := fmt.Sprintf("%s%s", reqURLForGEO, parameter)

	res2, err := requestBaidu("GetAddressViaGEO", reqURL)
	if err != nil {
		return res, err
	}
	if res2.(*StructGEOToAddress).Status != 0 {
		message := fmt.Sprintf("require failed")
		return res, errors.New(message)
	}
	res3 := res2.(*StructGEOToAddress)
	return res3, nil
}
