# asciimoji

Find relevant emojis from the command-line. Results are copied to your clipboard.

# Usage

```
$ asciimoji hurray
copied ＼(^o^)／

$ asciimoji -h
usage: asciimoji [<flags>] [<lookup>]

Look up ASCII emojis

Optional flags:
  -h, --help               Show context-sensitive help (also try --help-long and --help-man).
  -l, --list-all           list all available emojis
  -s, --search=SEARCH ...  search all available emojis
  -V, --version            Show application version.

Args:
  [<lookup>]  the keyword to lookup.

```

# Installation

```
go get github.com/jakewarren/asciimoji
```

# Customization

## Add a new emoji

Use the helper utility to easily add a new emoji and then run `make` to rebuild  & reinstall the binary. Multiple keywords can be specified.
```
cd $GOPATH/src/github.com/jakewarren/asciimoji
go run cmd/emojiadd/main.go "┌( ͝° ͜ʖ͡°)=ε/̵͇̿̿/’̿’̿ ̿" bond 007
make
```

## Add/Modify keywords

Edit the emojis.json file and then run `make` to rebuild  & reinstall the binary.
