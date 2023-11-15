# Backend

Step 1: Make sure to check whether the Golang is installed on your system by checking the version of Go.

```go version```

Step 2: Set the GOPATH by using the following command.

```export GOPATH=$HOME/go```

Step 3: Now, set the PATH variable with the help of the following command.

```export PATH=$PATH:$GOPATH/bin```

Step 4: Download and install project packages and dependencies.

```go get```

Step 5: Start the Mongo database service.

```docker-compose up mongodb```

Step 6: Run the app. (Standing in the "src" folder)

```go run main.go```