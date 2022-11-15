package models

type User struct {
	Id       int    `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
	Nama     string `gorm:"varchar(100)" json:"nama"`
	Username string `gorm:"varchar(30)" json:"username"`
	Password string `gorm:"varchar(255)" json:"password"`
}
