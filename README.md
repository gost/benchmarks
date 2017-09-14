# Benchmarks

This repository contains the benchmarks tests and results.

## Setup

```
$ sudo apt-get install apache2-utils
$ wget https://raw.githubusercontent.com/gost/benchmarks/master/docker-compose/docker-compose.yml
$ docker-compose up -d 
$ git clone https://github.com/gost/benchmarks.git
$ cd tests
```

## Environment

Used environment: 

- Azure Standard A1 Virtual Machine (1 Core, 1.75 GB memory)
- Client test tool on same machine

## Tests

1] Root test

Description: Request the SensorThings server root document

file: 1_get_root_document.sh

```
$ ab -n 5000 -c 5 http://localhost:8080/v1.0
```

## Results

date: 2017-09-14

server version: https://github.com/gost/server/commit/04171e4111917b5a16673559a4bd1dc97663914a
database version: https://github.com/gost/gost-db/commit/afe835f003af3f022b420c92493de16d95a189e0


Test 1: 825 request p/s
