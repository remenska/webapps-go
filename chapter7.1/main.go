package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
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
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println("----------------")
	vv := &Recurlyservers{Version: "1"}
	vv.Svs = append(vv.Svs, server{ServerName: "Shangai_VPN", ServerIP: "127.0.0.1"})
	vv.Svs = append(vv.Svs, server{ServerName: "Blaaah_VPN", ServerIP: "127.0.0.2"})
	output, err := xml.MarshalIndent(vv, "  ", "     ")
	if err != nil {
		fmt.Printf("error: %v, \n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
