/*
Run the following command in the terminal to get the package
Create a go.mod file inside the project folder
go mod init <Module-Name>
downloads the dependency package
go get -u github.com/go-sql-driver/mysql
*/

package main

import(
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("SQL", "username:password@(127.0.0.1:5432)/dbname?parseTime=true")
	if err != nil{
		log.Fatal(err)
	}

	if err := db.Ping(); err !=nil{
		log.Fatal(err)
	}
	{//Create a new table
		query := `
			CREATE TABLE IF NOT EXISTS users(
				id INT AUTO_INCREMENT,
				name VARCHAR(255),
				username TEXT NOT NULL,
				password TEXT NOT NULL,
				created_at DATETIME,
				PRIMARY KEY(id)
			);`
		if _, err := db.Exec(query); err != nil{
			log.Fatal(err)
		}
	}
	{//Insert a new record
		username := "johndoe"
		password := "secret"
		created_at := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?,?,?)`, username, password, created_at)
		if err != nil{
			log.Fatal(err)
		}
		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{//Query a single user
		var (
			id int
			username string
			password string
			createdAt time.Time
		)
		query := "SELECT id, username, password, crreated_at FROM users WHERE id=?"
		 if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
            log.Fatal(err)
        }

		fmt.Println(id, username, password, created_at)
	}
	{//Query all users
		type user struct{
			id int
			username string
			password string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil{
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next(){
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil{
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err !=nil{
			log.Fatal(err)
		}
		fmt.Printf("%#v", users)
	}
	{
		_, err := db.Exec(`DELETE FROM users WHERE id =?`,1)
		if err != nil{
			log.Fatal(err)
		}
	}
}