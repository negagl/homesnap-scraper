package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type House struct {
	PropertyID        int    `json:"propertyid"`
	FullStreetAddress string `json:"address"`
	City              string `json:"city"`
	State             string `json:"state"`
	Zip               string `json:"zip"`
	Beds              string `json:"beds"`
	BathsFull         string `json:"bathsfull"`
	SqFt              string `json:"sqft"`
	Value             string `json:"value"`
	Rent              string `json:"rent"`
}

func main() {

	// Instance of House
	var house House

	address := "1616 E Cornell Avenue"
	city := "Fresno"
	state := "CA"

	house.PropertyID = GetPropertyID(address, city, state)

	fmt.Println("PropertyID: ", house.PropertyID)

	GetPropertyDetails(strconv.Itoa(house.PropertyID))
}

func GetPropertyID(address string, city string, state string) int {

	address = strings.ReplaceAll(address, " ", "-")

	// First we connect to the GetByURL endpoint to extract the PropertyID
	url := "https://www.homesnap.com/service/PropertyAddresses/GetByUrl"
	method := "POST"
	req_address := state + "/" + city + "/" + address
	post_url := fmt.Sprintf(`{"url":"%s"}`, req_address)

	// Body of the request should be a Reader type
	payload := strings.NewReader(post_url)

	// Creating the request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Adding headers to the request
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "es-419,es;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36 Edg/109.0.1518.61")

	// Sending Request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()

	// Get the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Parse the JSON response
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	property := result["d"].(map[string]any)
	var propertyId int

	for key, value := range property {
		// Each value is an `any` type, that is type asserted as a string
		if key == "PropertyID" {
			propertyIdFloat := value.(float64)
			propertyId = int(propertyIdFloat)
			// fmt.Println(key, propertyId)

			break
		}
	}

	return propertyId
}

func GetPropertyDetails(propertyId string) []any {

	// First we connect to the GetByURL endpoint to extract the PropertyID
	url := "https://www.homesnap.com/service/Properties/GetDetails"
	method := "POST"

	requestBody := fmt.Sprintf(`{"propertyID":"%s"}`, propertyId)

	// Body of the request should be a Reader type
	payload := strings.NewReader(requestBody)

	// Creating the request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Adding headers to the request
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "es-419,es;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36 Edg/109.0.1518.61")

	// Sending Request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	// Get the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Parse the JSON response
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	// property := result["d"].(map[string]any)
	// var propertyData []any

	// for key, value := range property {
	// 	// Each value is an `any` type, that is type asserted as a string
	// 	if key == "PropertyID" {
	// 		propertyIdFloat := value.(float64)
	// 		propertyId = int(propertyIdFloat)
	// 		// fmt.Println(key, propertyId)

	// 		break
	// 	}
	// }

	return []any{}
}

// fmt.Println(result)
// os.Exit(0)
