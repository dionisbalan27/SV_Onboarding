package main

import (
	"backend-api/app"
	"backend-api/config"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig()

	//telegram := notification.TelegramNotification(config.CONFIG["URL_MS_TELEGRAM"], config.CONFIG["TELEGRAM_CHANNEL_ID"], config.CONFIG["ELASTIC_APM_SERVICE_NAME"])

	postgresConn, Sql, err := config.ConnectPostgres()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	router := app.InitRouter(postgresConn)
	log.Println("routes Initialized")

	port := config.CONFIG["PORT"]
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server Initialized")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//go telegram.Send("Info", "Customer Hub Profile started")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
