package algorath

import (
	"algorath/checkPrice"
	"algorath/checkWallet"
	"algorath/endpoint"
	"algorath/manager"
	"algorath/repository"
	"algorath/sendOrder"
)

func main() {

	database := repository.New()
	cw := checkWallet.New(database)
	cp := checkPrice.New(database)
	so := sendOrder.New(database)

	manager  := manager.New(cp, cw, so)

	endpoint.New(database, manager)


}