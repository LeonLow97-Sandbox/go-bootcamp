```go
func main() {
	db, err := newDB()
	defer func() {
		db.Close()
		fmt.Println("Closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	row := db.QueryRowContext(context.Background(), "SELECT birth_year FROM users WHERE name = 'Leon'")
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	var birth_year int64
	if err := row.Scan(&birth_year); err != nil {
		fmt.Println("row.Scan", err)
		return
	}

	fmt.Println("birth_year", birth_year)

	name := "Hanni Pham"
	birth_year = 2004

	// Inserting records to the database
	_, err = db.ExecContext(context.Background(), "INSERT INTO users(name, birth_year) VALUES ($1, $2)", name, birth_year)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}

	rows, err := db.QueryContext(context.Background(), "SELECT name, birth_year FROM users")
	if err != nil {
		fmt.Println("row.Scan", err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	if row.Err() != nil {
		fmt.Println("row.Err()", err)
		return
	}

	for rows.Next() {
		var name string
		var birth_year int64

		if err := rows.Scan(&name, &birth_year); err != nil {
			fmt.Println("rows.Scan", err)
			return
		}

		fmt.Println("name", name, "birth_year", birth_year)
	}
}
```
