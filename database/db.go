package database

import "github.com/jinzhu/gorm"

func DataBase() (*gorm.DB, error) {
	//db config
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dhlee dbname=kakao password=0546 sslmode=disable")
	//defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, nil

}
