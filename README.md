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

You can check the man page `man t`

## Contributing

The plain, good and old git-way of contributing.

The only dependency is `go`, which you can install trough your package
manager (apk on Alpine, pacman on Arch, apt on Debian-based etc). After
that, you can use `make get` to fetch all dependencies and start
hacking!

Send me your patches to my email at
[pedrolucasporcellis@gmail.com](mailto:pedrolucasporcellis@gmail.com)

(Remember to use [plain text](https://useplaintext.mail))!


## TODO

There's a lot of things I want to implement on `t`, heres some of them:

- Add a improved visualization when displaying notes
- Delete a note
- Make setup through `t` (`t init` and create the repository, etc).
- Add PGP support on notes
- Allow editing by title of note

## License

GNU GPL v3 License. Check
[LICENSE](https://git.sr.ht/~porcellis/t/tree/master/LICENSE).
