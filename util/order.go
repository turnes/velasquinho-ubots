package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Item struct {
	Product  string  `json:"produto"`
	Type     string  `json:"variedade"`
	Country  string  `json:"pais"`
	Category string  `json:"categoria"`
	Vintage  string  `json:"safra"`
	Price    float64 `json:"preco"`
}

type Order struct {
	Code     string  `json:"codigo"`
	Date     string  `json:"data"`
	Client   string  `json:"cliente"`
	Items    []Item  `json:"itens"`
	SubTotal float64 `json:"valorTotal"`
}

type Orders struct {
	URL    string
	Orders []Order
}

func (o *Orders) FetchOrders() error {
	apiClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, o.URL, nil)

	res, err := apiClient.Do(req)
	if handleError(err) {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if handleError(err) {
		return err
	}

	var orders []Order

	err = json.Unmarshal(body, &orders)
	if handleError(err) {
		return err
	}
	o.Orders = orders
	o.convertId()
	return nil

}

func (o *Orders) PrintOrders() {
	for _, x := range o.Orders {
		fmt.Printf("Codigo: %s, data: %s, Cliente: %s, Valor: %.2f\n", x.Code, x.Date, x.Client, x.SubTotal)
		for _, i := range x.Items {
			fmt.Printf("Produto: %s, Pais: %s, Preco: %.2f \n", i.Product, i.Country, i.Price)
		}
		fmt.Println("")
	}

}

func (o *Orders) convertId() {
	for i, x := range o.Orders {
		idClientStr := strings.Replace(x.Client, ".", "", -1)
		idClientStr = strings.TrimLeft(idClientStr, "0")
		o.Orders[i].Client = idClientStr
	}
}
