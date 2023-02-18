# trash - freedesktop-compliant "Trash can"
The trash project provides libraries and CLI commands that follow the [freedesktop.org trash specification](https://specifications.freedesktop.org/trash-spec/trashspec-1.0.html) in golang.  


# "trash" cli tool
## How to install
### Use "go install"
If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install).
```
$ go install github.com/nao1215/trash@latest
```

### Install from Package or Binary
[The release page](https://github.com/nao1215/trash/releases) contains packages in .deb, .rpm, and .apk formats.

## List of features
## How to use
### move file/directory to trash can
```
$ trash move PATH/TO/FILE
```
### list file/directory in trash can
```
$ trash list
```
### restore file/directory in trash can
```
$ trash restore FILE_NAME
```

### erase file/directory in trash can
```
$ trash erase FILE_NAME
```

# "trash" Library
## How to use
TODO:

# Contributing
First off, thanks for taking the time to contribute! ❤️  See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.
Contributions are not only related to development. For example, GitHub Star motivates me to develop!

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/trash/issues)

# LICENSE
The trash project is licensed under the terms of [MIT License](./LICENSE).