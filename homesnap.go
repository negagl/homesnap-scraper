package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type House struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Beds    string `json:"beds"`
	Baths   string `json:"baths"`
	Sqft    string `json:"sqft"`
	Price   string `json:"price"`
	Rent    string `json:"rent"`
}

func main() {

	address := "1616 E Cornell Avenue"
	city := "Fresno"
	state := "CA"

	address = strings.ReplaceAll(address, " ", "-")

	// First we connect to the GetByURL endpoint to extract the PropertyID
	url := "https://www.homesnap.com/service/PropertyAddresses/GetByUrl"
	method := "POST"
	req_address := state + "/" + city + "/" + address
	post_url := fmt.Sprintf(`{"url":"%s"}`, req_address)

	payload := strings.NewReader(post_url)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.homesnap.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "es-419,es;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "User=ID=998666696&Hash=0b7701ee4ca214832118018cb0dcd1933d8bd5da; ASP.NET_SessionId=mso5pppcidxlc3i0vjjpt0w0; ASP.NET_SessionId=0ongjuezjzusfhhtiqvpq4h1")
	req.Header.Add("origin", "https://www.homesnap.com")
	req.Header.Add("referer", "https://www.homesnap.com/"+req_address)
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"99\", \"Microsoft Edge\";v=\"109\", \"Chromium\";v=\"109\"")
	req.Header.Add("sec-ch-ua-mobile", "?1")
	req.Header.Add("sec-ch-ua-platform", "\"Android\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36 Edg/109.0.1518.61")
	req.Header.Add("x-costar-brand", "2")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Get the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON response
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	fmt.Println(result["d"].(map[string]interface{})["PropertyID"])
}
