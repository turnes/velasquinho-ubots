package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	ID   int    `json:"id"`
	Name string `json:"nome"`
	CPF  string `json:"cpf"`
}

type Clients struct {
	URL     string
	Clients []Client
}

func (c *Clients) FetchClients() error {

	apiClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, c.URL, nil)
	if handleError(err) {
		return err
	}

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

	var clients []Client

	err = json.Unmarshal(body, &clients)
	if handleError(err) {
		return err
	}
	c.Clients = clients

	return nil
}

func (c *Clients) PrintClients() {
	for _, x := range c.Clients {
		fmt.Printf("ID: %d\tNome: %s\tCPF: %s\n", x.ID, x.Name, x.CPF)
	}
}

func (c *Clients) GetNameByCode(id int) (string, error) {
	for _, x := range c.Clients {
		if x.ID == id {
			return x.Name, nil
		}
	}
	return "", errors.New("id not found")
}
