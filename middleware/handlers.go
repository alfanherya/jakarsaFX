package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alfanherya/jakarsafx/jakarsafx-registrasi/models"
	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgresql db
func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	//return the connection
	return db
}

//GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//get all the users in the db
	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	//send all the users as response
	json.NewEncoder(w).Encode(users)
}

// get one user from the DB by its userid
func getAllUsers() ([]models.User, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var users []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Nama, &user.Email, &user.Alamat, &user.Password)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}
