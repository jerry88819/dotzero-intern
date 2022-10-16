package main

import (
	_ "github.com/joho/godotenv/autoload"
	"fmt"
	"log"
	postgres "gitlab.com/dotzerotech/toolApi/database"
	"gitlab.com/dotzerotech/toolApi/model"
	router "gitlab.com/dotzerotech/toolApi/route"
)

func main() {
	fmt.Printf("tooling report api start\n")

	// connect to postgres
	db, err := postgres.Init()
	if err != nil {
		fmt.Printf("database error =>%v\n", err)
		return
	}
	fmt.Printf("database init success\n")

	// defer db.Close() // 全部做完就關掉

	defer func() {
		_db, err := db.DB()
		if err != nil {
			log.Println("Close db failed:", err)
			fmt.Printf("error =>%v\n", err)
			return
		}
		_db.Close()
	}()

	// set db in model package
	model.Init(db)

	fmt.Printf("model init success\n")
	// set router
	router.SetConfig()
	router.SetRouter()
	router.StartServer()

} // main()
