package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type log struct {
	Lang       string      `json:"lang"`
	App        string      `json:"app"`
	Level      string      `json:"level"`
	Controller string      `json:"controller"`
	Action     string      `json:"action"`
	Category   string      `json:"category"`
	Uri        string      `json:"uri"`
	Caller     string      `json:"caller"`
	IP         string      `json:"ip"`
	Pid        int         `json:"pid"`
	Method     string      `json:"method"`
	Rid        string      `json:"rid"`
	Runmode    string      `json:"runmod"`
	T          string      `json:"t"`
	TS         float64     `json:"ts,string"`
	Msg        string      `json:"msg"`
	Detail     interface{} `json:"detail"`
	Trace      interface{} `json:"trace"`
}

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	js := new(log)
	fmt.Println("==================== LOG-FMT ==================")
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		line := consolescanner.Bytes()
		err = json.Unmarshal(line, js)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}

		fmt.Printf("%s | %s | %s \n", Green(js.T), js.Rid, js.IP)
		fmt.Printf("%s | %s \n", Green(js.Method), js.Uri)
		fmt.Printf("%s | %s | %s | %s | %s \n", Green(js.Level), js.Controller, js.Action, js.Category, js.Caller)
		fmt.Printf("%s %s \n", Green(js.Msg), js.Detail)

		fmt.Println("===============================================")
	}
}
