package main

import (
	"database/sql"
	"fmt"
	"log"
	
	"github.com/l0ng7h0r/golang/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DBDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("--- Roles table ---")
	rows, err := db.Query("SELECT id, name FROM roles")
	if err != nil {
		fmt.Println("Error querying roles:", err)
	} else {
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			fmt.Printf("Role ID: %d, Name: %s\n", id, name)
		}
		rows.Close()
	}

	fmt.Println("--- Users table ---")
	rows, err = db.Query("SELECT id, email FROM users")
	if err != nil {
		fmt.Println("Error querying users:", err)
	} else {
		for rows.Next() {
			var id string
			var email string
			rows.Scan(&id, &email)
			fmt.Printf("User ID: %s, Email: %s\n", id, email)
		}
		rows.Close()
	}

	fmt.Println("--- User_Roles table ---")
	rows, err = db.Query("SELECT user_id, role_id FROM user_roles")
	if err != nil {
		fmt.Println("Error querying user_roles:", err)
	} else {
		for rows.Next() {
			var uid string
			var rid int
			rows.Scan(&uid, &rid)
			fmt.Printf("UserID: %s, RoleID: %d\n", uid, rid)
		}
		rows.Close()
	}
}
