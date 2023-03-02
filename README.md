# TCP client-server application on Golang

Implementation:
- TCP server with secured (TLS) connection on Golang
- Proto API to communicate between client and server
- PostgreSQL as database
- Client to make API requests for the Demo
- Docker images for both the server and the client, managed with the Docker Compose

## Application Demo

This requires Docker and Docker Composer to be installed
https://docs.docker.com/compose/install/

Run cmd and see logs of the demo
```
make start
```

What's tested and added to Demo at this stage
```
server_1  | Listening: [::]:9992
server_1  | New client: 172.25.0.4:56044
client_1  | Client 172.25.0.4:56044 connected to: 172.25.0.3:9992
client_1  | --------------------------------------------
client_1  | Login with none existing user: qwerty
client_1  | Got error: not found: User not found
client_1  | --------------------------------------------
client_1  | Login with existing user: alex
client_1  | Got session ID: 13d5f506-cfef-4268-9f84-5523cdea4f45
client_1  | --------------------------------------------
client_1  | Logout with invalid session
client_1  | Got error: not authenticated: invalid session ID
client_1  | --------------------------------------------
client_1  | Logout with valid session
client_1  | Got succeed result
server_1  | Close client: 172.25.0.4:56044
```

- ...

## Testing

Run cmd run unit tests
```
make test
```

Some packages are covered with Unit Tests to show the testing approach.

- Mocks are generated for Interfaces using the [minimock](https://github.com/gojuno/minimock) tool.
Any change in Interface require the mock structures to be regenerated with cmd `make gen_code`.
The `minimock` tool must be installed for the mock code generation (see installation details in the tool source link).


- The [ginkgo](https://github.com/onsi/ginkgo) and [gomega](https://github.com/onsi/gomega) frameworks are used to write unit tests.
- The [easyjson](https://pkg.go.dev/github.com/tcolar/easyjson) is used to generate JSON marshaller

## TODO
- Implement workers pool for handling client connections on server side
- The TTL config for the session ID
- Implement brood force using Redis for the Login process
- Handle all proto errors properly
- Complete tests for the rest of the packages
