package data

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/turnes/velasquinho-ubots/util"
	"log"
	"time"
)

type VelasquinhoData struct {
	HashOrders  string
	HashClients string
	Clients     util.Clients
	Orders      util.Orders
	ClientUrl   string
	OrdersURL   string
}

func (v *VelasquinhoData) Initialize(ClientUrl, OrdersURL string) {
	v.ClientUrl = ClientUrl
	v.Clients.URL = ClientUrl
	v.OrdersURL = OrdersURL
	v.Orders.URL = OrdersURL
	v.checkUpdates()
}

func (v *VelasquinhoData) checkUpdates() {
	fmt.Println(time.Now().String() + " - Start syncing")
	err := v.Clients.FetchClients()
	if err != nil {
		log.Println(err)
	}
	err = v.Orders.FetchOrders()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(time.Now().String() + " - Stop syncing")
	v.Clients.PrintClients()
	v.Orders.PrintOrders()
}

func (v *VelasquinhoData) Run() {
	c := cron.New()
	err := c.AddFunc("@every 5m", v.checkUpdates)
	if err != nil {
		log.Println(err)
	}
	go c.Start()
}
