<h1 align="center">
  GOLANG-REDIS-MONGODB
</h1>


<p align="center">
  <a href=""><img src="https://img.shields.io/badge/golang-1.19-brightgreen"></a>
  <a href=""><img src="https://img.shields.io/badge/redis-io-red"></a>
  <a href=""><img src="https://img.shields.io/badge/mongo-DB-green"></a>
</p>


  
## How To Use

To clone and run this application, you'll need [Docker Desktop](https://www.docker.com/products/docker-desktop/), [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash

# 1) Run Redis and MongoDB                   
$ docker-compose -f docker-compose.yml up -d  

# 2) Install dependencies for GoApp             
$ go mod download

# 3) Build the GoApp                           
$ go build -o ./bin/ ./cmd/filesearch-worker/.

# 4) Run GoApp                              
$ ./bin/filesearch-worker 

```
