# kakao test1

### 개발 프레임워크, 문제해결 방법, 빌드 및 실행 방법

1. 개발프레임워크 `GoLang 의 Echo 웹 프레임 워크`
2. 문제해결 방법 `Gorm 을 이용한 ORM DB 연동, PostGreSql DB 사용`
3. 빌드 및 실행 방법 `개발 당시 Go version 1.13.6 / OS : Mac `


``` 
git clone https://github.com/kinggigo/donghatest1.git 
go get -u github.com/labstack/echo/...
go get -u github.com/jinzhu/gorm

go build -o bin/main main.go
-> bin/main 실행 파일 생성. 

```
