package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"example.com/web-service-gin/internal/interfaces"

	_ "github.com/lib/pq"
)

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

func GETListAllRestaurants() *[]byte {
	db := OpenConnection()

	sqlStatement := `
    SELECT r.id, r.name, r.owner, r.address, r.description, r.image
    FROM restaurants r
    WHERE 1 = 1
    ORDER BY r.name ASC, r.id ASC
    `

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var restaurants []interfaces.Restaurant

	for rows.Next() {
		var restaurant interfaces.Restaurant
		rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Owner, &restaurant.Address, &restaurant.Description, &restaurant.Image)
		restaurants = append(restaurants, restaurant)
	}

	restaurantsBytes, _ := json.MarshalIndent(restaurants, "", "\t")

	defer db.Close()
	return &restaurantsBytes
}

func FindRestaurant(term string) *[]byte {
	db := OpenConnection()

	sqlStatement := `
    SELECT r.id, r.name, r.owner, r.address, r.description, r.image
    FROM restaurants r
    WHERE name = $1;`

	rows, err := db.Query(sqlStatement, term)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var restaurants []interfaces.Restaurant

	for rows.Next() {
		var restaurant interfaces.Restaurant
		rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Owner, &restaurant.Address, &restaurant.Description, &restaurant.Image)
		restaurants = append(restaurants, restaurant)
	}

	restaurantsBytes, _ := json.MarshalIndent(restaurants, "", "\t")

	defer db.Close()
	return &restaurantsBytes
}

func DeleteRestaurant(restaurantId string) string {
	db := OpenConnection()

	sqlStatement := `
    DELETE FROM restaurants
    WHERE id = ` + restaurantId + `;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	defer db.Close()
	return restaurantId
}

func POSTHealth() {
	db := OpenConnection()

	defer db.Close()
}
