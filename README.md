# create

[![Release](https://github.com/bschaatsbergen/create/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/bschaatsbergen/create/actions/workflows/goreleaser.yaml) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bschaatsbergen/create) ![GitHub commits since latest release (by SemVer)](https://img.shields.io/github/commits-since/bschaatsbergen/create/latest) [![Go Reference](https://pkg.go.dev/badge/github.com/bschaatsbergen/create.svg)](https://pkg.go.dev/github.com/bschaatsbergen/create) ![GitHub all releases](https://img.shields.io/github/downloads/bschaatsbergen/create/total) 

A modern UNIX file generation tool

## Brew

To install create using brew, simply do the below.

```sh
brew tap bschaatsbergen/create
brew install create
```

## Binaries

You can download the [latest binary](https://github.com/bschaatsbergen/create/releases/latest) for Linux, MacOS, and Windows.

## Features

- Create files
- Create files with content
- Creates the directories if it doesn't exist
- Create files with permissions
- File overwrite protection

## Examples

Using `create` is very simple.

### Create a file

To create a file, simply run the below command.

```
$ create foo.txt
```

### Create a file with content

To create a file with content, you use the `-c` flag:

```
$ create foo.txt -c "example"
```

### Create directories if they don't exist

You simply add the path to the file you want to create:

```
$ create foo/bar.txt
```

### Create a file and set permissions

To create a file and set permissions, use the `-m` flag:

```
$ create foo.txt -m 0777
```

### Overwrite a file

To overwrite a file (`create` has file overwrite protection by default), use the `--force` flag:

```
$ create foo.txt --force
```

## Contributing

Contributions are highly appreciated and always welcome.
Have a look through existing [Issues](https://github.com/bschaatsbergen/create/issues) and [Pull Requests](https://github.com/bschaatsbergen/create/pulls) that you could help with.
