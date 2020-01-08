PKGNAME = git.sr.ht/~porcellis/t
GOPATH = $(realpath .go)
PKGPATH = .go/src/$(PKGNAME)

PREFIX?=/usr/local
_INSTDIR=$(DESTDIR)$(PREFIX)
BINDIR?=$(_INSTDIR)/bin
MANDIR?=$(_INSTDIR)/share/man

all: t doc

install: all
	mkdir -m755 -p $(BINDIR) $(MANDIR)/man1
	install -m755 t $(BINDIR)/t
	install -m644 docs/t.1 $(MANDIR)/man1/t.1

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
	rm -rf t docs/*.gz

.PHONY:
	t get clean doc
