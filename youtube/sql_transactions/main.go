package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
)

type User struct {
	Name      *string
	BirthYear *int64
}

func main() {
	conn, err := newConn()
	if err != nil {
		fmt.Println("newConn", err)
		return
	}

	defer func() {
		conn.Close(context.Background())
		fmt.Println("Closed")
	}()

	newName := func(s string) *string {
		return &s
	}

	newBirthYear := func(b int64) *int64 {
		return &b
	}

	if err := insertUsers(conn, []User{
		{
			Name:      newName("Hanni Pham"),
			BirthYear: newBirthYear(2004),
		},
		{
			Name:      newName("Haerin"),
			BirthYear: newBirthYear(2006),
		},
	}); err != nil {
		fmt.Println("insertUsers", err)
	}
}

func insertUsers(conn *pgx.Conn, users []User) error {
	if err := conn.BeginFunc(context.Background(), func(tx pgx.Tx) error {
		for _, user := range users {
			_, err := tx.Exec(context.Background(), "INSERT INTO users(name, birth_year) VALUES($1, $2)", user.Name, user.BirthYear)
			if err != nil {
				return fmt.Errorf("tx.Exec %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("conn.BeginFunc %w", err)
	}

	// tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	// if err != nil {
	// 	return fmt.Errorf("db.BeginTx %w", err)
	// }
	// defer func() {
	// 	_ = tx.Rollback(context.Background())
	// }()

	// for _, user := range users {
	// 	_, err := tx.Exec(context.Background(), "INSERT INTO users(name, birth_year) VALUES($1, $2)", user.Name, user.BirthYear)
	// 	if err != nil {
	// 		return fmt.Errorf("tx.Exec %w", err)
	// 	}
	// }

	// if err := tx.Commit(context.Background()); err != nil {
	// 	return fmt.Errorf("tx.Commit %w", err)
	// }

	return nil
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5433",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect %w", err)
	}

	return conn, nil
}
