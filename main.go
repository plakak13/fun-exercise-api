package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	g := e.Group("/api")

	handler := wallet.New(p)

	g.GET("/v1/wallets", handler.WalletsHandler)
	g.GET("/v1/users/:id/wallets", handler.WalletsUserHandler)
	g.POST("/v1/wallets", handler.CreateUserWalletHandler)
	g.PATCH("/v1/wallets", handler.UpdateUserWalletHabdler)

	e.Logger.Fatal(e.Start(":1323"))
}
