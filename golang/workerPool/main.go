package main

type User struct {
	Name string
}
type Database interface {
	GetUser(id int) User
}
type MyDb struct{}

func (d *MyDb) GetUser(id int) User {
	return User{"Joe"}
}

//START OMIT
var workers chan Database

func init() {
	workers = make(chan Database, 10)
	for i := 0; i < 10; i++ {
		workers <- NewDatabaseObject()
	}
}

func GetUserName(id int) string {
	db := <-workers
	defer func() { workers <- db }()
	return db.GetUser(id).Name
}

//END OMIT
