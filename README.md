# mai-labs-url-shortner
A Golang library to listen and publish emails

## Usage

### Run the application
```
make run LISTNER=lister TOPIC=topic
```
- Supported listners can be checked from the pkg/constants/constants.go file 
### Build the Go Binaries
```
make build
```

## Testing

### Development Testing
- `make run` run the server locally. A publisher script is also attacked to trigger a sample event. 
- To begin testing, first you need to login with provider by hitting localhost:8080/login?provider={provider}
- To run the sample queue (Nats) locally, just run `docker run --name nats -p 4222:4222 -d nats:latest`

### Steps To Local Test

- Install Docker and run the queue `docker run --name nats -p 4222:4222 -d nats:latest`
- Run the service `make run LISTNER=nats TOPIC=jobs`
- If token file has not been created or is expired, call `localhost:8080/login?provider=gmail`
- Use the sample publisher script to push event into queue `make publish`

### Components
- Server which serves auth related API's
- Listner to any given queue
- Data is store in Sqlite
