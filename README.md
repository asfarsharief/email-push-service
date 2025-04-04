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
