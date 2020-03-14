package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"kakaoTest1/database"
	"net/http"
	"os"
	"strconv"
)

type Moudata struct {
	Key       int `gorm:"primary_key"`
	Region    string
	Target    string
	Usage     string
	Limit     string
	Rate      string
	Institute string
	Mgmt      string
	Reception string
}

func main() {

	db, err := database.DataBase()
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	}

	e := echo.New()
	// 1. csv파일 데이터베이스 저장 하는 API
	e.GET("/makeDB", func(c echo.Context) error {
		//초기화
		db.DropTable(Moudata{})

		file, err := os.Open("./tete1.csv")
		if err != nil {
			fmt.Println("@@@@@", err)
		}

		rdr := csv.NewReader(bufio.NewReader(file))

		rows, err := rdr.ReadAll()
		if err != nil {
			fmt.Println("@@@@@2", err)
		}

		for i, row := range rows {

			moudata := Moudata{}
			for j := range row {
				fmt.Printf("%s", rows[i][j])
				switch j {
				case 0:
					a, _ := strconv.Atoi(rows[i][j])
					moudata.Key = a
					break
				case 1:
					moudata.Region = rows[i][j]
					break
				case 2:
					moudata.Target = rows[i][j]
					break
				case 3:
					moudata.Usage = rows[i][j]
					break
				case 4:
					moudata.Limit = rows[i][j]
					break
				case 5:
					moudata.Rate = rows[i][j]
					break
				case 6:
					moudata.Institute = rows[i][j]
					break
				case 7:
					moudata.Mgmt = rows[i][j]
					break
				case 8:
					moudata.Reception = rows[i][j]
					break
				}
			}
			if i == 0 {
				db.CreateTable(&moudata)
			} else {
				db.Create(&moudata)
			}

		}

		db.Exec("commit;")
		return c.String(http.StatusOK, "DB 생성이 완료되었습니다.")
	})

	// 2. 지원 하는 지자체 목록 검색 API -> 전체 목록 출력인것처럼 보임

	e.GET("/mouList", func(c echo.Context) error {
		moudata := []Moudata{}
		db.Find(&moudata)
		//fmt.Println("@@@@" , moudata)
		return c.JSON(http.StatusOK, moudata)
	})

	//3. 지자체명을 입력 받아 해당 지자체의 지원정보를 출력 하는 API -> 입력이 지자체 입력, 출력이 지자체 정보
	e.GET("/mouData/:Region", func(c echo.Context) error {
		moudata := Moudata{}
		Region := c.Param("Region")

		if Region == "" {
			return c.String(http.StatusBadRequest, "찾고자하는 지역을 입력하세요")
		}
		db.Where("region =?", Region).Find(&moudata)
		return c.JSON(http.StatusOK, moudata)
	})

	//4. 지원하는 지자체 정보 수정 기능 API -> 지자체 명을 넣고 정보 입력하면 바뀌게 함
	e.POST("/mouUpdate", func(c echo.Context) error {

		return c.String(http.StatusOK, "변경이 완료되었습니다.")
	})

	//5. 지언한도 컬럼에서 지원금액으로 내림차순 정렬(지원금액이 동일하면 이차보전 평균 비율이 적은 순서)하여 특정갯수( 입력:k) 만큼 지자체명 출력

	//6. 이차보전 컬럼에서 보전비율이 가장 작은 추천 기관명을 출력하는 API 개발

	e.Logger.Fatal(e.Start(":1323"))
}
