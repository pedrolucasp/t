PKGNAME = git.sr.ht/~porcellis/t
GOPATH = $(realpath .go)
PKGPATH = .go/src/$(PKGNAME)

all: t

.go:
		mkdir -p $(dir $(PKGPATH))
		ln -fTrs $(realpath .) $(PKGPATH)

t: .go
	env GOPATH=$(GOPATH) go build -o $@ ./main.go

clean:
	rm -rf .go t
