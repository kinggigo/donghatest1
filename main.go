package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dhlee dbname=kakao password=0546 sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
