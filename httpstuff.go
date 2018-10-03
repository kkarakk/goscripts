package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	// message := map[string]interface{}{
	// 	"user":     "Yogesh2",
	// 	"key":      "1c23cfa983XX",
	// 	"mobile":   +916263532853,
	// 	"message":  "++test",
	// 	"senderid": "SRISAL",
	// 	"accusage": 1,
	// }
	// bytesRepresentation, err := json.Marshal(message)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//resp, err := http.Get("https://httpbin.org/get")
	//http://send.smshost.co.in/submitsms.jsp?user=yalla2&key=1c23cfa983XX&mobile=+916269904310&message=++getgps&senderid=SRISAL&accusage=1
	resp, err := http.PostForm("http://send.smshost.co.in/submitsms.jsp",
		url.Values{"user": {"yalla2"}, "key": {"1c23cfa983XX"}, "mobile": {"+918269177508"}, "message": {"  test"}, "senderid": {"SRISAL"}, "accusage": {"1"}})

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(resp)
	}

	log.Println(string(body))
}
