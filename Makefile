##
## File description:
## Makefile
##

all: 		$(NAME)

$(NAME) : go run *.go

run: setup all
	go run *.go

setup:
	export GO111MODULE="on"
