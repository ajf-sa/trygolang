package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	UserID   string  `sql:"user_id"`
	UserName string  `sql:"username"`
	Email    string  `sql:"email"`
	Password string  `sql:"password"`
	Mood     *string `sql:"mood"`
}

func main() {
	url := "postgres://admin:secret@localhost:5432/pgx-migrate-v1?sslmode=disable"
	ctx := context.Background()
	db, _ := pgxpool.Connect(ctx, url)
	_, err := migrate.New(
		"file://db/migrations",
		url,
	)
	if err != nil {
		log.Fatal(err)
	}
	// if err := m.Up(); err != nil {
	// 	log.Fatal(err)
	// }
	// var ah User
	// ah.UserName = "weww"
	// ah.Email = "s@wew.sdf"
	// ah.Password = "wiefj"
	// _, err = db.Exec(ctx, "insert into users(username,email,password) values($1,$2,$3)", ah.UserName, ah.Email, ah.Password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var u []*User
	rows, _ := db.Query(ctx, "Select * from users")
	defer rows.Close()
	for rows.Next() {
		var uu User
		err = rows.Scan(&uu.UserID, &uu.UserName, &uu.Email, &uu.Password, &uu.Mood)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uu)
		u = append(u, &uu)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
