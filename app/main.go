package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/iqbalrestu07/datting-apps-api/app/server"
	middlewr "github.com/iqbalrestu07/datting-apps-api/app/server/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {

}

func main() {
	config := server.GetAPPConfig()
	db, err := server.InitDB()
	if err != nil {
		log.Fatal("Error connection database:", err)
	}
	defer func() {
		dbConn, _ := db.DB()
		_ = dbConn.Close()
	}()
	e := echo.New()
	e.HideBanner = true
	splash()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
		middlewr.DBTransactionMiddleware(db),
	)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Invoice System API"})
	})
	e.Static("/public", "file")

	e.Pre(middleware.RemoveTrailingSlash())

	server.BuildRoutes(e, db)
	go func() {
		if err := e.Start(config.AppPort); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	waitForShutdown(e)
	log.Fatal(e.Start(config.AppPort)) //nolint
}

func waitForShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func splash() {
	colorReset := "\033[0m"

	splashText := `
	   .___       __  .__                                                                            .__ 
  __| _/____ _/  |_|__| ____    ____           _____  ______ ______  ______         _____  ______ |__|
 / __ |\__  \\   __\  |/    \  / ___\   ______ \__  \ \____ \\____ \/  ___/  ______ \__  \ \____ \|  |
/ /_/ | / __ \|  | |  |   |  \/ /_/  > /_____/  / __ \|  |_> >  |_> >___ \  /_____/  / __ \|  |_> >  |
\____ |(____  /__| |__|___|  /\___  /          (____  /   __/|   __/____  >         (____  /   __/|__|
     \/     \/             \//_____/                \/|__|   |__|       \/               \/|__|       

 `
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}
