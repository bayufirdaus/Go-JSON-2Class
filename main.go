package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Image struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"image"`

	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
}

type Client struct {
	Clients struct {
		Address struct {
			City   string `json:"city"`
			State  string `json:"state"`
			Street string `json:"street"`
			Zip    string `json:"zip"`
		} `json:"address"`

		Age      int64  `json:"age"`
		Company  string `json:"company"`
		Email    string `json:"email"`
		Gender   string `json:"gender"`
		ID       string `json:"id"`
		IsActive string `json:"isActive"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
	} `json:"clients"`
}

func getProduct(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Product

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	ID := request.ID
	Name := request.Name
	Type := request.Type
	Imghgt := request.Image.Height
	Imgurl := request.Image.URL
	Imgwdh := request.Image.Width
	Thmhgt := request.Thumbnail.Height
	Thmurl := request.Thumbnail.URL
	Thmwdh := request.Thumbnail.Width

	stmt, err := db.Prepare("INSERT INTO product (id, name, type, img_hgt, img_url, img_wth, thm_hgt, thm_url, thm_wth) VALUES(?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(ID, Name, Type, Imghgt, Imgurl, Imgwdh, Thmhgt, Thmurl, Thmwdh)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Product Created")
}

func getClient(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request Client

	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println("Failed decoding json message")
	}

	City := request.Clients.Address.City
	State := request.Clients.Address.State
	Street := request.Clients.Address.Street
	Zip := request.Clients.Address.Zip
	Age := request.Clients.Age
	Company := request.Clients.Company
	Email := request.Clients.Email
	Gender := request.Clients.Gender
	ID := request.Clients.ID
	IsActive := request.Clients.IsActive
	Name := request.Clients.Name
	Phone := request.Clients.Phone

	stmt, err := db.Prepare("INSERT INTO clients (city, state, street, zip, age, company, email, gender, id, isActive, name, phone) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(City, State, Street, Zip, Age, Company, Email, Gender, ID, IsActive, Name, Phone)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Client Created")
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_json")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/product", getProduct).Methods("POST")
	r.HandleFunc("/client", getClient).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
