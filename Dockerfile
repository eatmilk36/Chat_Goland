# 使用 Go 的官方映像
FROM golang:1.22

# 設定工作目錄
WORKDIR /app

# 複製 go.mod 和 main.go 如果有外部依賴要有go.sum
COPY go.mod ./
RUN go mod download

# 複製所有 Go 源代碼並編譯
COPY . ./
RUN go build -o myapp .

# 確保 myapp 擁有執行權限
RUN chmod +x myapp

# 設定執行命令
CMD ["./myapp"]