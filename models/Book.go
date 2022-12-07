package models

type Book struct {
	ID     int64 `db:"id, primarykey, autoincrement" json:"id"`
	Name   string
	Author string
	Genre  string
}
