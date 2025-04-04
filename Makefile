#Makefile to run the application
#=======================================================================

dep:
	@go get ./...
	
run:
	@go run cmd/app/main.go listen -l ${LISTNER} -t ${TOPIC}

publish: 
	@go run publisher/main.go

build-image:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo -o bin/app  cmd/app/main.go