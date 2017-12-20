package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Rectangle string `json:"rectangle"`
}

func PrintJson(obj interface{}) {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s.", err)
	}
	fmt.Printf("%s\n", data)
}

func main() {
	url := "http://restapi.amap.com/v3/ip?ip=112.74.164.107&key=c9556ad40fc6811432b2973ed7c1782d&output=json"
	//url := "http://restapi.amap.com/v3/ip"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Resp StatusCode failed:%s", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", response)

	PrintJson(response)
}

func exec_method_map() {
	url := "http://restapi.amap.com/v3/ip?ip=112.74.164.107&key=c9556ad40fc6811432b2973ed7c1782d&output=json"
	//url := "http://restapi.amap.com/v3/ip"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(body))

	var response map[string]interface{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

	fmt.Println("-------------------")
	//fmt.Println(response["info"])
	for k, v := range response {
		fmt.Println(k, v)
	}
}
