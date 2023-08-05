package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const pageSize = 10

var db *sql.DB

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://myuser:mypassword@postgres:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	createDummyData(db)
	return db
}

func main() {
	db = setupDB()
	defer db.Close()

	http.HandleFunc("/users", listUsers)
	http.ListenAndServe(":8080", nil)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * pageSize

	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Page %d:\n", page)
	for _, user := range users {
		fmt.Fprintf(w, "User ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}

func createDummyData(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			age INT
		);
		INSERT INTO users (name, age) VALUES
		('John', 25),
		('Jane', 30),
		('Michael', 28),
		('Emily', 22),
		('William', 35),
		('Olivia', 29),
		('James', 27),
		('Sophia', 24),
		('Alexander', 31),
		('Ava', 26),
		('Ethan', 29),
		('Isabella', 32),
		('Noah', 28),
		('Mia', 30),
		('Liam', 25),
		('Amelia', 27),
		('Benjamin', 29),
		('Charlotte', 30),
		('Mason', 32),
		('Harper', 28),
		('Elijah', 26),
		('Evelyn', 23),
		('Logan', 27),
		('Abigail', 29),
		('Lucas', 33),
		('Emily', 24),
		('Jackson', 28),
		('Elizabeth', 30),
		('Aiden', 25),
		('Sofia', 26),
		('Samuel', 29),
		('Avery', 31),
		('Daniel', 26),
		('Scarlett', 27),
		('Matthew', 32),
		('Victoria', 28),
		('David', 30),
		('Chloe', 24),
		('Joseph', 27),
		('Grace', 29),
		('Carter', 34),
		('Riley', 25),
		('Gabriel', 26),
		('Zoey', 31),
		('Dylan', 29),
		('Lily', 30),
		('Grayson', 32),
		('Layla', 28),
		('Oliver', 25),
		('Luna', 27),
		('Christopher', 29),
		('Emma', 22),
		('Alexander', 25),
		('Mia', 28),
		('James', 30),
		('Sophia', 26),
		('Michael', 29),
		('Ava', 32),
		('William', 28),
		('Olivia', 31),
		('Ethan', 26),
		('Isabella', 29),
		('Noah', 33),
		('Charlotte', 24),
		('Liam', 27),
		('Amelia', 30),
		('Benjamin', 25),
		('Harper', 26),
		('Elijah', 31),
		('Evelyn', 29),
		('Logan', 34),
		('Abigail', 25),
		('Lucas', 26),
		('Emily', 31),
		('Jackson', 29),
		('Elizabeth', 30),
		('Aiden', 32),
		('Sofia', 28),
		('Samuel', 30),
		('Avery', 27),
		('Daniel', 25),
		('Scarlett', 26),
		('Matthew', 31),
		('Victoria', 29),
		('David', 34),
		('Chloe', 25),
		('Joseph', 26),
		('Grace', 31),
		('Carter', 29),
		('Riley', 30),
		('Gabriel', 32),
		('Zoey', 28),
		('Dylan', 30),
		('Lily', 27),
		('Grayson', 26),
		('Layla', 31),
		('Oliver', 29),
		('Luna', 25),
		('Christopher', 26),
		('Eleanor', 30),
		('Owen', 29),
		('Penelope', 33),
		('Sebastian', 24),
		('Hannah', 27),
		('Jackson', 26),
		('Stella', 31),
		('Aiden', 28),
		('Nora', 30),
		('Matthew', 25),
		('Ella', 26),
		('Mason', 31),
		('Aria', 29),
		('Jacob', 34),
		('Aurora', 25),
		('Eli', 26),
		('Lucy', 31),
		('Levi', 29),
		('Anna', 30),
		('David', 28),
		('Hazel', 33),
		('Daniel', 25),
		('Savannah', 26),
		('Henry', 31),
		('Zoe', 29),
		('Joseph', 30),
		('Camilla', 25),
		('Alexander', 26),
		('Maya', 31),
		('Samuel', 29),
		('Lila', 30),
		('Benjamin', 32),
		('Violet', 28),
		('Ethan', 30),
		('Natalie', 27),
		('Logan', 26),
		('Lena', 31),
		('Caleb', 29),
		('Stella', 30),
		('Mason', 25),
		('Mila', 26),
		('James', 31),
		('Hannah', 29),
		('Elijah', 30),
		('Eleanor', 34),
		('Oliver', 25),
		('Lily', 26),
		('Daniel', 31),
		('Grace', 29),
		('William', 30),
		('Ella', 28),
		('Michael', 32),
		('Scarlett', 33),
		('Lucas', 25),
		('Sofia', 26),
		('Jackson', 31),
		('Victoria', 29),
		('Avery', 30),
		('Mia', 27),
		('Noah', 26),
		('Jie Wei', 26);
	`)
	if err != nil {
		log.Fatalf("Error creating dummy data: %v", err)
	}
}
