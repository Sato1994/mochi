package main

import (
	"github.com/labstack/echo/v4"
	"mochi/db"
	"mochi/routes"

)

func main() {
 db.Init()

 server := echo.New()
 routes.RegisterRoutes(server)
 server.Logger.Fatal(server.Start(":1323"))
}

