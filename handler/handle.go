package handler

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
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

// 1. csv파일 데이터베이스 저장 하는 API
func MakeDB(db *gorm.DB) echo.HandlerFunc {
	// 1. csv파일 데이터베이스 저장 하는 API
	return func(c echo.Context) error {
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
	}
}

// 2. 지원 하는 지자체 목록 검색 API -> 전체 목록 출력인것처럼 보임
func MouList(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		moudata := []Moudata{}
		db.Find(&moudata)
		//fmt.Println("@@@@" , moudata)
		return c.JSON(http.StatusOK, moudata)
	}
}

//3. 지자체명을 입력 받아 해당 지자체의 지원정보를 출력 하는 API -> 입력이 지자체 입력, 출력이 지자체 정보
func MouData(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		moudata := Moudata{}
		Region := c.Param("Region")

		if Region == "" {
			return c.String(http.StatusBadRequest, "찾고자하는 지역을 입력하세요")
		}
		db.Where("region =?", Region).Find(&moudata)
		return c.JSON(http.StatusOK, moudata)
	}
}

//4. 지원하는 지자체 정보 수정 기능 API -> 지자체 명을 넣고 정보 입력하면 바뀌게 함
func MouUpdate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		moudata := c.Get("mouData")
		db.Save(&moudata)
		return c.String(http.StatusOK, "변경이 완료되었습니다.")
	}
}

//5. 지원한도 컬럼에서 지원금액으로 내림차순 정렬(지원금액이 동일하면 이차보전 평균 비율이 적은 순서)하여 특정갯수( 입력:k) 만큼 지자체명 출력
func MouLimit(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var moudata []Moudata
		k := c.Param("k")
		kint, _ := strconv.Atoi(k)
		regionName := make([]string, kint)

		db.Order("rate desc,  \"limit\" desc").Limit(kint).Find(&moudata)

		for i, _ := range moudata {
			regionName[i] = moudata[i].Region
		}
		return c.JSON(http.StatusOK, regionName)
	}
}

//6. 이차보전 컬럼에서 보전비율이 가장 작은 추천 기관명을 출력하는 API 개발
func MouRate(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var moudata []Moudata
		//regionName := make([]string,1)

		db.Order("rate").Limit(1).Find(&moudata)

		//for i , _ := range moudata{
		//	regionName[i] = moudata[i].Region
		//}
		return c.JSON(http.StatusOK, moudata[0].Region)
	}
}
