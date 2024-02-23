FROM golang:1.22.0

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /mochi

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

# # COPYしないとNo such file or directoryエラーが出る
# COPY /startup.sh /startup.sh
# RUN chmod +x /startup.sh

# CMD ["/startup.sh"]

#  go mod init myapp
#  go get github.com/labstack/echo/v4
#  go get gorm.io/gorm
#  go get gorm.io/driver/mysql