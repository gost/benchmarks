# Benchmarks

This repository contains the benchmarks tests and results.

## Setup

```
$ sudo apt-get install docker.io
$ sudo apt-get install apache2-utils
$ sudo apt install docker-compose
$ wget https://raw.githubusercontent.com/gost/benchmarks/master/docker-compose.yml
$ sudo docker-compose up -d 
$ git clone https://github.com/gost/benchmarks.git
$ cd tests
```

## Environment

Used environment: 

- Azure Standard A1 Virtual Machine with Ubuntu 16.10 (1 Core, 1.75 GB memory, 30GB disk)
- Client test tool on same machine

## Description Tests

1] Root test

Description: Request the SensorThings server root document

Amount: 100.000 times, concurrency 5

```
$ ab -n 100000 -k -c 50 http://localhost:8080/v1.0
```

2] Deep insert configuration

Description: Deep insert thing, location, sensor, observedproperty, datastream

Amount: 100.000 times, concurrency 5

```
$ ab -n 100000 -k -c 50 -X POST http://localhost:8080/v1.0
```

3] Insert Observations

Description: insert observations

Amount: 100.000 times, concurrency 5

4] Post Observation

Description: post observations

Amount: 100.000 times, concurrency 5

4] Get Observations by datastream

Description: post observations

Amount: 100.000 times, concurrency 5

5] Get observations by datastream with OData filter

todo

6] Create 100000 datastreams

todo

7] Create 100000 observations 

todo

## Results

date: 2017-09-14

server version: https://github.com/gost/server/commit/04171e4111917b5a16673559a4bd1dc97663914a

database version: https://github.com/gost/gost-db/commit/afe835f003af3f022b420c92493de16d95a189e0



| Test     | Duration (s)  |  Request/second | time per request (ms) |
|----------|---------------|-----------------|-----------------------|
| test 1   |  121          | 825             | 1.2                   |
| test 2   |               |                 |                       |

