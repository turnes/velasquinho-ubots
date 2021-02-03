package report

import (
	"github.com/turnes/velasquinho-ubots/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

type OrderReportAlltime struct {
	Client   string   `json:"id"`
	Nome     string   `json:"nome"`
	Codes    []string `json:"codes"`
	SubTotal float64  `json:"ValorTotal"`
}

type OrderReportByYear struct {
	Client   string  `json:"id"`
	Nome     string  `json:"nome"`
	Code     string  `json:"codigo"`
	SubTotal float64 `json:"ValorTotal"`
}

func AllTime(clients util.Clients, orders util.Orders) []OrderReportAlltime {

	sort.Slice(orders.Orders, func(i, j int) bool {
		return orders.Orders[i].Client < orders.Orders[j].Client
	})

	var ReportAllTime []OrderReportAlltime
	lastId := ""
	reportPosition := -1
	for _, c := range orders.Orders {
		if lastId != c.Client {
			client := OrderReportAlltime{
				Client:   c.Client,
				Nome:     getName(clients, c.Client),
				Codes:    []string{c.Code},
				SubTotal: c.SubTotal,
			}
			ReportAllTime = append(ReportAllTime, client)
			lastId = c.Client
			reportPosition++
		} else {
			ReportAllTime[reportPosition].Codes = append(ReportAllTime[reportPosition].Codes, c.Code)
			ReportAllTime[reportPosition].SubTotal += c.SubTotal
		}
	}

	sort.Slice(ReportAllTime, func(i, j int) bool {
		return ReportAllTime[i].SubTotal > ReportAllTime[j].SubTotal
	})

	return ReportAllTime
}

func ByYear(clients util.Clients, orders util.Orders, year string) OrderReportByYear {
	var ordersYear []util.Order
	for _, c := range orders.Orders {
		y := strings.Split(c.Date, "-")
		if y[2] == year {
			ordersYear = append(ordersYear, c)
		}
	}
	sort.Slice(ordersYear, func(i, j int) bool {
		return ordersYear[i].SubTotal > ordersYear[j].SubTotal
	})
	clientOrder := ordersYear[0]

	return OrderReportByYear{
		Client:   clientOrder.Client,
		Nome:     getName(clients, clientOrder.Client),
		Code:     clientOrder.Code,
		SubTotal: ordersYear[0].SubTotal,
	}

}

func getName(clients util.Clients, id string) string {
	//idClientStr := strings.Replace(id, ".", "",-1)
	//idClientStr = strings.TrimLeft(idClientStr, "0")
	idClient, _ := strconv.Atoi(id)
	name, err := clients.GetNameByCode(idClient)
	if err != nil {
		log.Println(err)
	}
	return name

}
