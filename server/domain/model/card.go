package model

type Card struct {
	No    int64 `gorm:"primary_key"`
	Word  string
	Imgid int64
}
