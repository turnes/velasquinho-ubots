# velasquinho-ubots

## Endpoints
### 1 - Liste os clientes ordenados pelo maior valor total em compras.
https://ubots.turnes.com.br/api/v1/report/orders

### 2 - Mostre o cliente com maior compra única no último ano (2016).
https://ubots.turnes.com.br/api/v1/report/orders/year/2016


## Setting Golang project

```
git clone https://github.com/turnes/velasquinho-ubots.git
go mod init github.com/turnes/velasquinho-ubots
go get -u github.com/gorilla/mux 
```

## Setting the repository
### Branch
![Alt text](images/settingsbranch.png?raw=true "Main branch")

## [Docker](https://hub.docker.com/repository/docker/turnes/velasquinho-ubots)
Dockerfile -> multi-stage build, so a tiny docker image
### Tags  
main or development 
### Run
`docker run -it --rm  -p 8080:5000 --name velasquinho-ubots-api turnes/velasquinho-ubots:[tag]`


# Development env
```
git clone git clone https://github.com/turnes/velasquinho-ubots.git
docker build -t velasquinho-ubots . && docker run -it --rm  -p 8080:5000 --name velasquinho-ubots-api velasquinho-ubots
```



