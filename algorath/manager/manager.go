package manager

import (
	"algorath/checkPrice"
	"algorath/checkWallet"
	"algorath/sendOrder"
	"fmt"
	"time"
)

type ManagerI interface {
	Launch() (error)
}

type Manager struct {

	cp checkPrice.PriceI
	cw checkWallet.WalletI
	so sendOrder.OrderI
}

func New(cp checkPrice.PriceI, cw checkWallet.WalletI, so sendOrder.OrderI) ManagerI{

	newManager := new(Manager)

	newManager.cp = cp
	newManager.cw = cw
	newManager.so = so

	return newManager

}

func (m *Manager) Launch() error{


	//Check if my wallet has at least 100 USD
	err := m.cw.CheckWallet()
	if err != nil{
		fmt.Errorf(err.Error())
		return err
	}

	//Check if fee is less than 0,25%
	err = m.cp.CheckPrice()
	if err != nil{
		fmt.Errorf(err.Error())
		return err
	}

	time.Sleep(3*time.Second)

	//Check again if fee is less than 0,25%
	err = m.cp.CheckPrice()
	if err != nil{
		fmt.Errorf(err.Error())
		return err
	}

	//Send order
	err = m.so.SendOrder()
	if err != nil{
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}