package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Restaurant struct {
	ID        	string	`json:"id"`
	Name        string	`json:"name"`
	Owner 		string  `json:"owner"`
	Address     string	`json:"address"`
	Description string	`json:"description"`
	Image		string	`json:"image`
}
type Product struct {
	ID        		int		`json:"id"`
	ID_Restaurant	int		`json:"name"`
	Name			string	`json:"name"`
	Description		string	`json:"description"`
	Price			float32	`json:"price"`
	Image			string	`json:"image"`
}
type Additional struct {
	Name	string	`json:"name"`
	Price	float32	`json:"price"`
}

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "admin"
	PASSWORD = "admin"
	DBNAME   = "postgres"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func GETListAllRestaurants(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	sqlStatement := `
	SELECT name
	FROM restaurants;
	`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var restaurants []Restaurant

	for rows.Next() {
		var restaurant Restaurant
		rows.Scan(&restaurant.Name)
		restaurants = append(restaurants, restaurant)
	}

	filmsBytes, _ := json.MarshalIndent(restaurants, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(filmsBytes)

	defer db.Close()
}

func POSTHealth(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	sqlStatement := `SELECT NOW() as now`

	rows, err = db.Query(sqlStatement)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	return, _ := json.MarshalIndent(rows[0], "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(return)

	defer db.Close()
}

func main() {
	http.HandleFunc("/health", POSTHealth)
	// http.HandleFunc("/restaurants", POSTHandler)
	// http.HandleFunc("/products", POSTHandler)

	http.HandleFunc("/restaurants", GETListAllRestaurants)
	//http.HandleFunc("/restaurants/:id", GETListAllRestaurants)
	//http.HandleFunc("/restaurants/:id/products", GETListAllRestaurants)
	//http.HandleFunc("/products/:id", GETListAllRestaurants)
	//http.HandleFunc("/products/:id/additionals", GETListAllRestaurants)

	log.Fatal(http.ListenAndServe(":3009", nil))
}