package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Meeting struct {
	remoteip string
	duration int
	confid   string
}
type CmdParams struct {
	m    Meeting
	flag int
}

func flag_parse(p *CmdParams) {
	flag.StringVar(&p.m.remoteip, "ip", "127.0.0.1", "盒子Ip地址，默认127.0.0.1")
	flag.IntVar(&p.m.duration, "t", 5, "停留时间,默认5秒")
	flag.StringVar(&p.m.confid, "confid", "", "会议号，默认为空")
	flag.Parse()
}
func startMeeting(m *Meeting) {
	fullip := "http://" + m.remoteip + ":8282/op?"
	action := "action=start"
	confId := "&confid=" + m.confid
	resp, err := http.Get(fullip + action + confId)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func endMeeting(m *Meeting) {
	fullip := "http://" + m.remoteip + ":8282/op?"
	action := "action=stop"
	resp, err := http.Get(fullip + action)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {

	var params CmdParams
	flag_parse(&params)
	for {
		startMeeting(&params.m)
		time.Sleep(time.Duration(params.m.duration) * time.Second)

		endMeeting(&params.m)
		time.Sleep(time.Duration(2) * time.Second)
	}
}
