# Benchmarks

This repository contains the benchmarks tests and results.

## Results Round 1

date: 2017-09-18

server version: <a href="https://github.com/gost/server/commit/085b21972f9151b559059aa50f7c5ce48930602c">085b21</a>

database version: <a href="https://github.com/gost/gost-db/commit/afe835f003af3f022b420c92493de16d95a189e0">afe835</a>

| Name     | Test 1 (rps)  |  Test 2 (rps)   | Test 3 (rps)      | Test 4 (rps)  |
|----------|---------------|-----------------|-------------------|---------------|
| GOST     | 5000          | 224             | 37                | 284           |
| ref      | 3085          | 87              | 175               | 5773          |




## Setup

```
$ sudo apt-get update
$ sudo apt-get install docker.io
$ sudo apt-get install apache2-utils
$ sudo apt install docker-compose
$ wget https://raw.githubusercontent.com/gost/benchmarks/master/docker-compose.yml
$ sudo docker-compose up -d 
$ git clone https://github.com/gost/benchmarks.git
$ cd tests
$ sh tests.sh
```

## Environment

Used environment: 

- Azure Standard D2s v3 Virtual Machine with Ubuntu 17 (2 Core, 8 GB memory, 30GB disk)
- Client test tool on same machine

## Description Tests

1] Root test

Description: Request the SensorThings server root document

```
$ ab -n 2000 -k -c 50 http://localhost:8080/v1.0
```

2] Deep insert configuration

Description: Deep insert thing, location, sensor, observedproperty, datastream

```
$ ab -n 2000 -p 2_post_metadata.json -k -c 50  -T 'content-type-:application/json' http://localhost:8080/v1.0/Things
```

3] Post Observations

Description: insert observations

```
$ ab -n 2000 -p 3_post_observation.json -k -c 50  -T 'content-type-:application/json' http://localhost:8080/v1.0/Observations
```

4] Get Observations

Description: post observations

```
$ ab -n 2000 -k -c 50 http://localhost:8080/v1.0/Observations
```
