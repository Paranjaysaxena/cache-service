# Cache Service
This project contains code for a full-fledged RPC service which would allow to get and set data to a Redis cache.

### Steps to run
1. Clone this repository under `$GOPATH/src/`.
2. Run `./init/initialize_redis.sh` to start Redis server.
3. To run server, run the file `server/user.go`
4. To run gRPC client, run `client/user.go`.
