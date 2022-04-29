package main

import "fmt"

type AuthID int
type StaffID int

type Auth struct {
	ID int
}

type Staff struct {
	ID     int
	AuthID AuthID
}

func main() {
	auth := new(Auth)
	auth.ID = 2
	staff := new(Staff)
	staff.ID = 1
	staff.AuthID = AuthID(auth.ID)
	fmt.Println("Hello, World!, ", staff.AuthID)
}
