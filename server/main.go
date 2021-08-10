package main

import (
	"/interfaces"
	"fmt"
	"infra/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	listeningPort = "50051"
	dbDriver      = "mysql"
)

func main() {
	db, err := sqlConnect()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	server := interfaces.NewServer(interfaces.ServerParams{
		CardRepository: repository.NewCardRepository(db),
	})
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "mycard"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	return gorm.Open(DBMS, CONNECT)
}
