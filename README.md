# off_the_grid

RPC basic test implementation.
Client uses gorutines for each connection, dumps the connection results into a channel.
Client spawns a special gorutine that reads from the comunication channel and logs the results.

Run instructions:

1. add off_the_grid project to the
2. run server/main.go
3. run client/main.go
