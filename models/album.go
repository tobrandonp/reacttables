package models

type Album struct {
	ID     string  `bson:"ID"`
	Title  string  `bson:"Title"`
	Artist string  `bson:"Artist"`
	Price  float64 `bson:"Price"`
}
