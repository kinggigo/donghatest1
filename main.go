package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kakaoTest1/database"
	"kakaoTest1/handler"
)

func main() {

	db, err := database.DataBase()
	defer db.Close()

	if err != nil {
		fmt.Print(err)
	}

	e := echo.New()
	// 1. csv파일 데이터베이스 저장 하는 API
	e.GET("/makeDB", handler.MakeDB(db))

	// 2. 지원 하는 지자체 목록 검색 API -> 전체 목록 출력인것처럼 보임
	e.GET("/mouList", handler.MouList(db))

	//3. 지자체명을 입력 받아 해당 지자체의 지원정보를 출력 하는 API -> 입력이 지자체 입력, 출력이 지자체 정보
	e.GET("/mouData/:Region", handler.MouData(db))

	//4. 지원하는 지자체 정보 수정 기능 API -> 지자체 명을 넣고 정보 입력하면 바뀌게 함
	e.POST("/mouUpdate", handler.MouUpdate(db))

	//5. 지원한도 컬럼에서 지원금액으로 내림차순 정렬(지원금액이 동일하면 이차보전 평균 비율이 적은 순서)하여 특정갯수( 입력:k) 만큼 지자체명 출력
	e.GET("/mouLimit/:k", handler.MouLimit(db))

	//6. 이차보전 컬럼에서 보전비율이 가장 작은 추천 기관명을 출력하는 API 개발
	e.GET("/mouRate", handler.MouRate(db))

	e.Logger.Fatal(e.Start(":1323"))
}
