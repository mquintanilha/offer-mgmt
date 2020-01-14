package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://sandbox-api.lomadee.com/v3/{{ token }}/offer/_search?keyword={{ product }}&sourceId={{ sourd_id }}"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "sandbox-api.lomadee.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
