# GOLANG GOTO

A small experimental project inspired by [iridakos/goto](https://github.com/iridakos/goto/)
and [pass](https://www.passwordstore.org/), but much less useful because it offers
only the basic functionality.  It was created because I wanted to try out golang,
and nothing else. To use it, just do the standard `$GOPATH` magic, and then use
the `Makefile` to build/install.

Unfortunately it is not possible to change the working dir of a
parent shell from an app, so a bash function of some sort is always needed
to wrap the actual executable. For example, if you want to use the alias `g`
for faster jumping around, you need to add the following to your `~/.bashrc`:

```
function g() {
    _OUTPUT=$(goto --signal 5 "$@")

    if [ $? -eq 5 ]; then
        cd "${_OUTPUT}"
    else
        printf "${_OUTPUT}\n"
    fi
}
```

Now you should be able to use this:

```
$ g -h
Jump to specified dir alias

Usage:
  goto <alias> [flags]
  goto [command]

Available Commands:
  add         Add a new goto alias
  help        Help about any command
  init        Initialize a goto repository
  ls          List dir aliases

Flags:
  -h, --help   help for goto

Use "goto [command] --help" for more information about a command.

$ g init
Initialized new goto repo at /home/yourusername/.goto

$ g add vim ~/.vim/bundle
Added: vim => /home/yourusername/.vim/bundle

$ g ls
/home/yourusername/.goto
└── vim

$ g vim
# will actually cd into ~/.vim/bundle
```

Just like `pass`, you can init multiple goto "repos",
but that's usually not necessary. The idea can be extended further but I'm
disappointed in how golang works and feels, so this will probably be abandoned
until further notice.

Some of the projects used for snooping around golang syntax:

* [spf13/cobra](https://github.com/spf13/cobra)
* [spf13/viper](https://github.com/spf13/viper)
* [cayleygraph/cayley](https://github.com/cayleygraph/cayley)
