package main

import (
	"encoding/json"
	"fmt"
)

type Child struct {
	Key      string
	ConstKey string
}

type Serverslice struct {
	Childs []Child
}

func main() {

	//str := `{"childs":[{"key":"rmessage","constKey":"RMESSAGE"},{"key":"rop","constKey":"ROP"}]}`

	//str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	var s Serverslice

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("json parse error", err)
		return
	}
	fmt.Println(s)

	fmt.Println(s.Childs[0].Key)
	fmt.Println(s.Childs[0].ConstKey)
}
