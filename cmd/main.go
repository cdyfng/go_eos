package main

import (
	"encoding/json"
	eos "github.com/eoscanada/eos-go"
	"fmt"
)

func main(){
	//fmt.Printf("main")
	api := eos.New("https://mainnet.eoscanada.com")

	info, err := api.GetInfo()
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}