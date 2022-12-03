# Profiler
This microservice will deals with the user profiles 
It will be the only service that can edit, get the user profile data.

It might be called by the API to create profile data when a 
new user register himself, or to get some users profile data.


# Docker
    $ docker-compose up (--build)
```--build only if you edit the code```


# Install module with

    $ go get go.mongodb.org/mongo-driver/mongo

    User docker ;) 


# Errors 

errors : ../../../../../../go/src/google.golang.org/api/storage/v1/storage-gen.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/grpc/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/http/configure_http2_go116.go:1:1:

FIX : go env -> export GO111MODULE="on"
