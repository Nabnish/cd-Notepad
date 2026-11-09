package main

import (
	"context"
	"fmt"
	"github.com/cd-Notepad/Backend/db"
	"log" 
	"net/http"
	"time"
	"github.com/cd-Notepad/Backend/handlers"

	"github.com/cd-Notepad/Backend/services"
)

type Application struct {
	Models services.Models
}



func main() {
	
	Connection()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Connection() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
		fmt.Printf("server running")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	services.New(mongoClient)
	fmt.Printf("server running")
	log.Fatal(http.ListenAndServe(":8080",handlers.CreateRouter()))

}
