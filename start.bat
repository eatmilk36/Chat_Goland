@REM docker-compose down
@REM docker-compose up --build
@REM docker-compose up -d
docker-compose up --build -d

--------------------------
go test
go test -v      //詳細輸出
go test -cover

// 測試詳細報告
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

// 使用SQLite
set CGO_ENABLED=1