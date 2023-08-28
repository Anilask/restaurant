package databse

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = DBinstance()
var err error

func DBinstance() *sql.DB {
	db, err := sql.Open("mysql", "root:1208@tcp(127.0.0.1:3306)/Restaurant")
	if err != nil {
		fmt.Println(err)
	}
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set up the database connection with the context
	// db, err := sql.Open("mysql", db)
	// if err != nil {
	// 	fmt.Println("Error opening database:", err)
	// 	return
	// }
	defer db.Close()

	// Try to establish the connection with the context
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	fmt.Println("Connected to MySQL database successfully!")
	return db
}

func OpenCollection(db *sql.DB, collectionName string) *sql.DB {
	var collection *sql.DB = db.database("restaurant").collection(collectionName)
}
