package main

import (
	"fmt"
	"encoding/json"
)

// Data Structures

type Track struct {
	Top_traks	[]Track_info
}

type Track_info struct {
    Info       []Request_info
	Pag   	   []Pagination
	Data       []Offers
}

type Request_info struct {
	Status	string `json:"status"`
	Message	string `json:"message"`
}

type Pagination struct {
	Page		string `json:"page"`
	Size		string `json:"size"`
	Total_Size	string `json:"totalSize"`
	Total_Page	string `json:"totalPage"`
}

type Offers struct {
	Data_feed	[]Feed
}

type Feed struct {
	Id			string `json:"id"`
	Name		string `json:"name"`
	Category	string `json:"category"`
	Link		string `json:"link"`
	Thumbnail	string `json:"thumbnail"`
	Price		string `json:"price"`
	Discount	string `json:"discount"`
	Installment	string `json:"installment"`
	Store		string `json:"store"`
}

func main() {

	const jsonData = `{
		"requestInfo": {
			"status": "",
			"message": ""
		},
		"pagination": {
			"page": 0,
			"size": 0,
			"totalSize": 0,
			"totalPage": 0
		},
		"offers": [
			{
				"id": "234512",
				"name": "Celular",
				"category": {},
				"link": "",
				"thumbnail": "",
				"price": 0,
				"discount": 0,
				"installment": {},
				"store": {}
			}
		]

	}`

	data := []byte(jsonData)

	offers_day := Offers.Feed{}
	err := json.Unmarshal(data, &offers_day)
		if err != nil {
			fmt.Println(err)
			return
		}
	// fmt.Printf(offers_day.Id, offers_day.Name)
	fmt.Printf("Id: %+v, Name: %+v", offers_day.Id, offers_day.Name)

}