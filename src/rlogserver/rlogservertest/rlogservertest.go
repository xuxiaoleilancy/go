package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://10.10.6.37:3000/"

	payload := strings.NewReader("{   \"user\":\"xuxl\",\n    \"module\":\"meeting\",\n    \"action\":\"muteAll\",\n    \"msg\":\"this is a  body message\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "5856c7ca-e84e-d9b1-10b6-c8dda71e1d48")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
