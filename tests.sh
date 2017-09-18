ab -n 2000 -k -c 50 http://localhost:8080/v1.0
ab -n 2000 -p 2_post_metadata.json -k -c 50  -T 'content-type-:application/json' http://localhost:8080/v1.0/Things
ab -n 2000 -p 3_post_observation.json -k -c 50  -T 'content-type-:application/json' http://localhost:8080/v1.0/Observations
ab -n 2000 -k -c 50 http://localhost:8080/v1.0/Observations


