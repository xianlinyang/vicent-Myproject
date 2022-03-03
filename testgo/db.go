package main

import "fmt"

type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	fmt.Println(key)
	if value, err := db.Get(key); err == nil {
		return value
	} else {
		fmt.Println(value)
		fmt.Println(err)
	}

	return -1
}
