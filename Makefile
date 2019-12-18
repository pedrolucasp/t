PKGNAME = t

GOPATH = $(realpath .go)

all: t

t:
	env go build -o bin/$(PKGNAME) src/main.go
