package models

type Address struct {
	Province string `json:"Province" bson:"Province"`
	District string `json:"District" bson:"District"`
	Pincode  int    `json:"Pincode" bson:"Pincode"`
}

type User struct {
	Name    string  `json:"Name" bson:"Name"`
	Age     int     `json:"Age" bson:"Age"`
	Address Address `json:"Address" bson:"Address"`
}
