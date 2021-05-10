package main

import (
	"algorath/algorath"
	"algorath/checkPrice"
	"algorath/checkWallet"
	"algorath/endpoint"
	"algorath/manager"
	"algorath/repository"
	"algorath/sendOrder"
	"fmt"
	"time"
)

func main() {

	database := repository.New()
	cw := checkWallet.New(database)
	cp := checkPrice.New(database)
	so := sendOrder.New(database)

	manager  := manager.New(cp, cw, so)

	api := endpoint.New(database, manager)

	api.HandleRequests()

	algorath.Running = true

	for algorath.Running {
		if algorath.Running==false{
			fmt.Println("aasdasd")
			return
		}

 		time.Sleep(5*time.Second) //Wait 5 second for another check to finish or not the procedure
	}
}