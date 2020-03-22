# t


`t` is an minimalistic tool to take notes.

## Philosophy

Instead of using a web editor or some native app that under the hood is
a web page, or even a fully bloated app to take your notes and store 
it on someone elses server with their proprietary software and APIs we 
choose to use something way more simple, a combination of two very 
well documented, free and trusted tools: *git* and *markdown*.

## Design

`t` is designed to be completely simple and out of your way. Especially
when taking notes, sometimes all you want is somewhere to drop a bunch
of thoughts.

When writing a note, `t` will open your `EDITOR` (we recommend vim). As
soon that you write all you need and quit the app, `t` will push your
note to the configured git server. That's it.

## Usage

`t` is POSIX compliant, meaning that it uses a letter to handle it's actions.

- `l` will list your notes by it's modification date
- `c` will create a new note using the current date. You can supply
a title in quoted strings
- `e` will edit a note, by default it will edit the last one. You can
supply a position of the index you want to edit.

You can check the man page `man t` if you need any help.

## Installing

For now, you can only install by cloning this repository and running the
installation script.

So open your favorite terminal and fire these commands:
- `git clone https://git.sr.ht/~porcellis/t`
- `make && sudo make install`

If you have any interest in packaging this software to your distro of
choice, please let me know, so I can update the steps here.

## Contributing

The plain, good and old git-way of contributing.

The only dependency is `go`, which you can install trough your package
manager (apk on Alpine, pacman on Arch, apt on Debian-based etc). After
that, you can use `make get` to fetch all dependencies and start
hacking!

Send me your patches to my email at
[pedrolucasporcellis@gmail.com](mailto:pedrolucasporcellis@gmail.com)

(Remember to use [plain text](https://useplaintext.email))!

## Roadmap

There's a lot of things I want to implement on `t`, here is some of them:

- Add a improved visualization when displaying notes
- Support OS Pager when listing notes
- Delete a note
- Make setup through `t` (`t init` would create the notes repository, set a remote, etc).
- Add PGP support for notes (this is a must)
- Allow editing by title of note
- Companion app [tw](https://git.sr.ht/~porcellis/tw)

## License

GNU GPL v3 License. Check
[LICENSE](https://git.sr.ht/~porcellis/t/tree/master/LICENSE).
