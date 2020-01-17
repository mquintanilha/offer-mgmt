package main

import (
	"fmt"
	"encoding/json"
	"log"
    "net/http"
)

// Data Structures

type Offers struct {
	Data_feed	Feed
}

type Feed struct {

	Category	Category
	Name		string       `json:"name"`
	Price_From	float32		 `json:"priceFrom"`
	Price		float32      `json:"price"`
	Installment	Installment
	Id			string       `json:"id"`
	Link		string 	     `json:"link"`
	Thumbnail	string       `json:"thumbnail"`
	Store		Store
}

type Category struct {
	Link	string
	Id		int
	Name	string
}

type Installment struct {
	Value 		float32
	Quantity 	int
}

type Store struct {
	Thumbnail	string		`json:"thumbnail"`
	Link 		string		`json:"link"`
	Id			int			`json:"id"`
	Name		string		`json:"name"`

}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	jsonData := `{
		"category": {
			"link": "http://api.lomadee.com/v3/1578677161277ae2a08a2/category/_id/0?sourceId=22830830",
			"id": 0,
			"name": "Geral"
		},
		"name": "Rele Auxiliar 24v 5 Terminais 0a 0a Vdo Cod.ref. D09308 Volvo /iveco",
		"priceFrom": 29.25,
		"price": 29.25,
		"installment": {
			"value": 29.25,
			"quantity": 1
		},
		"id": "62778060",
		"link": "https://developer.lomadee.com/redir/validation/?sourceId=22830830&appToken=1578677161277ae2a08a2",
		"thumbnail": "https://images-americanas.b2w.io/produtos/01/00/oferta/62778/0/62778063P1.jpg",
		"store": {
			"thumbnail": "https://www.lomadee.com/programas/BR/5632/imagemBox_80x60.png",
			"link": "https://developer.lomadee.com/redir/validation/?sourceId=22830830&appToken=1578677161277ae2a08a2",
			"id": 5632,
			"name": "Americanas.com"
		}
	}`

	offers_day := Feed{}
	err := json.Unmarshal([]byte(jsonData), &offers_day)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Nome:%s, Quantity:%+v, Value:%+v, Loja:%s", offers_day.Name, offers_day.Installment.Quantity, offers_day.Installment.Value, offers_day.Store.Name)

}

func main() {
    http.HandleFunc("/offer/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}