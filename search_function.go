package main

import (
	"fmt"
	"encoding/json"
	"log"
    "net/http"
)

// Data Structures

type Track struct {
	Top_traks	Track_info
}

type Track_info struct {
    Info       Request_info
	Pag   	   Pagination
	Data       Offers
}

type Request_info struct {
	Status	string `json:"status"`
	Message	string `json:"message"`
}

type Pagination struct {
	Page		int  `json:"page"`
	Size		int  `json:"size"`
	Total_Size	int  `json:"totalSize"`
	Total_Page	int  `json:"totalPage"`
}

type Offers struct {
	Data_feed	Feed
}

type Feed struct {
	Id			string       `json:"id"`
	Name		string       `json:"name"`
	Category	string       `json:"category"`
	Link		string 	     `json:"link"`
	Thumbnail	string       `json:"thumbnail"`
	Price		float32      `json:"price"`
	Discount	float32      `json:"discount"`
	Installment	string       `json:"installment"`
	Store		string       `json:"store"`
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

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
	offers_day := Track{}
	err := json.Unmarshal(data, &offers_day)
		if err != nil {
			fmt.Println(err)
			return
		}
	fmt.Printf("Id: %v", offers_day.Top_traks.Data.Data_feed.Id)

}

func main() {
    http.HandleFunc("/offer/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}