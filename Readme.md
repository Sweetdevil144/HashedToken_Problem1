# Readme.md

## Problem :

**Go Application : Encrypt / Decrypt fields in transit**

- Build a Go webserver with two goroutines (G1, G2) and two endpoints (E1, E1) where E1 performs actions in G1, and E2
  in G2 respectively in a non-blocking fashion i.e. E1 operations and E2 operations are asynchronous and independent.
- E1 action : Make a request having at least three fields (string, UTC-timestamp, integer) which are written to a Redis
  database after the fields are encrypted. Choose some Redis HKEY to mark each of the requests
- E2 action : When a request (from: UTC-timestamp, to: UTC-timestamp) is made it retrieves all the requests from Redis
  database, decrypts them and shows them at stdout.

*Solution - Implementation*

**Note : Run Redis locally via Docker. Choose an appropriate Go based encrypt / decrypt library
to achieve the same.**