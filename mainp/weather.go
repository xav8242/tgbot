package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const key = "118757fd-72f3-42e6-96e4-36a657392d88"

type Data struct {
	Now        uint64                 `json: now`
	Now_dt     string                 `json: now_dt`
	Info       map[string]interface{} `json: info`
	Geo_object Geo_object             `json: geo_object`
	Fact       map[string]interface{} `json: fact`
	// Forecasts  []map[string]interface{} `json: forecasts`
}
type Geo_object struct {
	Country  IDName `json:country`
	District IDName `json: district`
	Locality IDName `json: locality`
	Province IDName `json: province`
}
type IDName struct {
	Id   uint64 `json: id`
	Name string `json: name`
}

func GetWather(lat, lon float64) *Data {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%s&lon=%s&extra=true", lat, lon), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("X-Yandex-API-Key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data Data

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%v\n", data.Geo_object.Locality)
	// fmt.Printf("%v\n", data.Fact["temp"])
	return &data

}
