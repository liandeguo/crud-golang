package database

import "fmt"

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
type DB struct {
	Items  []Item
	NextID int
}

func (db *DB) GetItem(id int) {
	fmt.Println(id)
	for i := 0; i < len(db.Items); i++ {
	}
}
