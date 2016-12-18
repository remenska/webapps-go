package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	file, err := os.Open("servers.json")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	var s Serverslice
	json.Unmarshal(data, &s)
	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)

	// type assertion?
	m := f.(map[string]interface{})
	fmt.Println("M=", m)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of type I don't know how to handle.")
		}
	}

	var ss Serverslice
	ss.Servers = append(ss.Servers, Server{ServerName: "Whatever", ServerIP: "25.143.123.52"})
	ss.Servers = append(ss.Servers, Server{"Blaaah", "234.255.244.23"})
	bb, err := json.Marshal(ss)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(bb))
}
