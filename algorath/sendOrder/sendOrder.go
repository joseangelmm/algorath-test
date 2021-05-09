package sendOrder


import (
	"algorath/repository"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/event"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/order"
	"github.com/bitfinexcom/bitfinex-api-go/pkg/mux"
	"fmt"
)

type OrderI interface {
	SendOrder() (error)
}

type Order struct {
	db repository.DatabaseI
}

func New(db repository.DatabaseI) OrderI{

	newManager := new(Order)

	newManager.db = db

	return newManager

}

func (w Order) SendOrder() error{

	cred, err := w.db.GetCredential()

	if err!= nil{
		fmt.Errorf(err.Error())
		return err
	}

	crash := make(chan error)
	auth := make(chan bool)

	ws := mux.New().TransformRaw().WithAPIKEY(cred.APIKey).WithAPISEC(cred.APISecret).Start()

	go func() {
		// if listener will fail, program will exit by passing error to crash channel
		crash <- ws.Listen(func(msg interface{}, err error) {
			if err != nil {
				fmt.Errorf("error received: %s\n", err)
			}

			switch v := msg.(type) {
			case event.Info:
				if v.Event == "auth" && v.Status == "OK" {
					auth <- true
				}
			case order.New:
				fmt.Errorf("%T: %+v\n", v, v)
				close(crash)
			}
		})
	}()

	//Makes operation to sell 100USD in BTC
	return nil
}
