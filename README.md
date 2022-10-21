# Authentificator
This microservice will deals with authentification and data management in the db.
It will be the only service with open acces to the database.

It might be called by the API to register or logged in a user.
it might be called by the microservice dealing with messages to update conversation



# Install module with

    $ go get go.mongodb.org/mongo-driver/mongo


# Errors 

errors : ../../../../../../go/src/google.golang.org/api/storage/v1/storage-gen.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/grpc/dial.go:1:1: expected 'package', found 'EOF'
../../../../../../go/src/google.golang.org/api/transport/http/configure_http2_go116.go:1:1:

FIX : go env -> export GO111MODULE="on"
