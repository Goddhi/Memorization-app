package main

import (
	"log"
	"net/http"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"


) 

func main() {
	log.Println("Starting Server .....")

	//setting up router

	router := gin.Default()

	router.GET("/api/account", func(c *gin.Context) {// gi.context contains both the request and  response 
		c.JSON(http.StatusOK, gin.H{
			"Hello": "your account is up",
		}) 
	})

	srv := &http.Server {
		Addr: ":8080",
		Handler: router,
	}

	/// Gracefully Server shutdown
	go func() {  // setting up the server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server %v\n", err)
		}
	}()

	log.Printf("Listening on port %v", srv.Addr)
	
	// wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// this blocks until a signal is passed into the quit channel
	<-quit

	// the context used to inform the server it has 5 seconds to finish the request it is currently handling 
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutting server
	log.Println("Shutting down server.....")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown ")
	}

}





/// wsl --set-version Ubuntu-22.04

