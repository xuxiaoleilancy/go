package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"strconv"
	//	"strings"
	"time"
)

type ServerCfg struct {
	SyslogServerAddr string `json:"addr"`
	Protocol         string `json:"protocol"`
	Port             int    `json:"port"`
	AppName          string `json:"appname"`
}

var (
	sysLog          *syslog.Writer
	SyslogServercfg = ServerCfg{
		SyslogServerAddr: "10.10.6.95",
		Protocol:         "udp",
		Port:             514,
		AppName:          "huijian"}
)

func LoadConfig(configfilename string) {

	log.Printf("starting load configure file %s", configfilename)
	data, err := ioutil.ReadFile(configfilename)
	if err != nil {
		log.Printf("ReadFile %s error:%v", configfilename, err)
	}
	log.Printf("configure file data %s", data)

	err = json.Unmarshal(data, &SyslogServercfg)
	if err != nil {
		log.Printf("json.Unmarshal error:%v", err)
	}
	log.Printf("get config json data:%v", SyslogServercfg)
}

// 处理GET请求
func handleGet(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	// 第一种方式
	// id := query["id"][0]

	// 第二种方式
	id := query.Get("id")

	fmt.Printf("GET: id=%s\n", id)

	fmt.Fprintf(writer, `{"code":0}`)
}

func handlePostJson(writer http.ResponseWriter, request *http.Request) {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(request.Body)

	// 用于存放参数key=value数据
	var params map[string]string

	// 解析参数 存入map
	decoder.Decode(&params)

	fmt.Printf("POST json: username=%s, password=%s\n", params["username"], params["password"])

	fmt.Fprintf(writer, `{"code":0}`)
}
func getMsgSuffix(request *http.Request) string {
	//	query := request.URL.Query()
	//	user := query.Get("user")
	//	module := query.Get("module")
	//	action := query.Get("action")
	//	msg := query.Get("msg")

	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	decoder.Decode(&params)

	//s, _ := ioutil.ReadAll(request.Body)
	bodyMsg := params["msg"]
	bodyuser := params["user"]
	bodymodule := params["module"]
	bodyaction := params["action"]
	return ", host=" + request.Host + ",user=" + bodyuser + ",module=" + bodymodule + ",action=" + bodyaction + ",bodymsg=" + bodyMsg + ",timestamp=" + strconv.FormatInt(time.Now().UnixNano(), 10)
}
func RLogResponse(rw http.ResponseWriter, request *http.Request) {
	//timeStr := time.Now().Format("2006-01-02 15:04:05:000")

	// //把body 内容读入字符串 select

	//sysLog.Emerg(" Emerg from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Alert(" Alert from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Crit(" Crit from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Err(" Err from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Warning(" Warning from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Notice(" Notice from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Info(" Info from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//sysLog.Debug(" Debug from 10.10.6.37 Hello world." + getMsgSuffix(request))
	//tmpLogMsg := "Info  long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message from 10.10.6.37 Hello worldInfo  long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message from 10.10.6.37 Hello worldInfo  long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message from 10.10.6.37 Hello worldInfo  long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message long message from 10.10.6.37 Hello world"
	//sysLog.Info(tmpLogMsg + getMsgSuffix(request) + " : count= " + strconv.Itoa(strings.Count(tmpLogMsg, "")))
	go sysLog.Info(getMsgSuffix(request))
	rw.WriteHeader(200)
	//fmt.Fprintf(rw, getMsgSuffix(request))
}

func main() {

	LoadConfig("conf.json")
	var err error
	addr := ""
	if len(SyslogServercfg.SyslogServerAddr) != 0 {
		addr = SyslogServercfg.SyslogServerAddr + ":" + strconv.Itoa(SyslogServercfg.Port)
	}

	sysLog, err = syslog.Dial(SyslogServercfg.Protocol, addr, syslog.LOG_WARNING|syslog.LOG_DAEMON, SyslogServercfg.AppName)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", RLogResponse)
	http.ListenAndServe(":3000", nil)
}
