package main

import (
	"errors"
	"time"
)

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
var TooBusyError = errors.New("Concurrency too high. Timed out waiting for worker.")

func GetUserName(id int) (string, error) {
	db, err := getDatabaseClient()
	if err != nil {
		return "", err
	}
	defer func() { workers <- db }()
	return db.GetUser(id).Name, nil
}
func getDatabaseClient() (Database, error) {
	select {
	case w := <-workers:
		return w, nil
	case <-time.After(5 * time.Second):
		return nil, TooBusyError
	}
}

//END OMIT
