PKGNAME = git.sr.ht/~porcellis/t
GOPATH = $(realpath .go)
PKGPATH = .go/src/$(PKGNAME)

all: t

.go:
		mkdir -p $(dir $(PKGPATH))
		ln -fTrs $(realpath .) $(PKGPATH)

doc:
	gzip -c docs/t.1 > docs/t.1.gz

t: .go
	env GOPATH=$(GOPATH) go build -o $@ ./main.go

get: .go
	env GOPATH=$(GOPATH) go get -d ./...

clean:
	rm -rf t

.PHONY:
	t get clean doc
