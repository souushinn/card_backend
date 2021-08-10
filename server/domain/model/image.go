package model

type Image struct {
	Id      int64 `gorm:"primary_key"`
	Img     string
	ImgName string
	CardNo  int64
}
