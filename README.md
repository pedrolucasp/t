# t

`t` is an minimalistic tool to take notes.

> A version of this README is [available in portuguese][readme-pt-br].

[readme-pt-br]: https://git.sr.ht/~porcellis/t/tree/master/README.pt-br.md

## Philosophy

Instead of using a web editor or some web browser disguised as a native
app, or even a giant and sluggish app to simple write text notes and
store it on someone else "cloud" with their proprietary software and
APIs we rather choose to use something way more simple, a combination of
two very well documented, free and trusted tools: *git* and *markdown*.

## Design

`t` is designed to be completely simple and out of your way. Especially
when taking notes, sometimes all you want is somewhere to drop a bunch
of thoughts.

When writing a note, `t` will open your `EDITOR` (or fallback to vim).
As soon that you write all you need and quit the app, `t` will push your
note to the configured git server. That's it.

## Usage

`t` has a small set of commands:

- `list` or `l` will list your notes indexed by it's modification date,
while piping to the default `PAGER` on your system (usually `less (1)`).

- `create` or `c` will create a new note using the current date. You can supply
a title if you want

- `edit` or `e` will reopen the note on your `EDITOR`, by default it
will edit the last one. You can supply the index of the note you want to edit.

- `show` or `s` will pipe the note to [glow][glow]

- `share` or `sh` will throw the note on the configured shared service,
it defaults to a public instance of [rascunho][rascunho] hosted at
[eletrotupi.com][eletrotupi], but you can customize to your paste
service.

[rascunho]: https://sr.ht/~porcellis/rascunho
[glow]: https://github.com/charmbracelet/glow
[eletrotupi]: https://eletrotupi.com 

You can check the man page `man t` if you need any help.

## Installing

For now, you can only install by cloning this repository and running the
installation script.

So open your favorite terminal and fire these commands:
- `git clone https://git.sr.ht/~porcellis/t`
- `make && sudo make install`

If you have any interest in packaging this software to your distro of
choice, please let me know, so I can update the steps here.

**Note**: By now, the `show` commands depends on 
[glow][glow] to be installed.

## Contributing

The plain, good and old git-way of contributing.

There's a couple of dependencies, namely `git`, `go`, `less`, and
[glow][glow]. All probably are available to install trough your package
manager (apk on Alpine, pacman on Arch, apt on Debian-based distros
etc). After that, you can use `make get` to fetch all dependencies and
start hacking!

[Send your patches][git-send-email] to my public-inbox at
[~porcellis/public-inbox@lists.sr.ht][public-inbox]

(Remember to use [plain text][plain-text])!

[public-inbox]: mailto:~porcellis/public-inbox@lists.sr.ht
[plain-text]: https://useplaintext.email
[git-send-email]: https://git-send-email.io

## Roadmap

There's a lot of things I want to implement on `t`, here is some of them:

- Add a improved visualization when displaying notes, probably using
ncurses
- ~~Support OS Pager when listing notes (this can be achieved by `t list
| less -r`)~~
- Delete a note
- Make setup through `t` (`t init` would create the notes repository, set a remote, etc).
- Add PGP support for notes (this is a must)
- Allow editing by title of note
- A Companion UI app [tw](https://git.sr.ht/~porcellis/tw)

## License

Under terms of the GNU GPL-3.0 License. Check
[LICENSE](https://git.sr.ht/~porcellis/t/tree/master/LICENSE).
