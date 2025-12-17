# CCWC

This repo contains John Crickett's coding challenge: [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc)

## Pre-requisites

- [Git](https://git-scm.com/)
- [Go](https://go.dev/) [Version used: `1.25.5`]
- [VS Code](https://code.visualstudio.com/) with Go extension installed

## External dependencies/packages

- [Cobra](https://github.com/spf13/cobra)

## Building the project

- In Terminal/Powershell, `cd` to project root directory `ccwc`
- Run: `go build`
- This will create the executable inside project root directory (`ccwc.exe` for Windows) 

## Running `ccwc` for a file

Run the commands inside the directory containing the file:

| Requirement | Command |
|-------------|---------|
| Count number of characters in a file | `ccwc --characters <FILE_NAME>` or `ccwc -c <FILE_NAME>` |
| Count number of words in a file | `ccwc --words <FILE_NAME>` or `ccwc -w <FILE_NAME>` |
| Count number of lines in a file | `ccwc --lines <FILE_NAME>` or `ccwc -l <FILE_NAME>` |
| Count number of bytes in a file | `ccwc --bytes <FILE_NAME>` or `ccwc -b <FILE_NAME>` |

> [!IMPORTANT]
> When no flags are supplied, by-default `-c -w -l` are passed to `ccwc` command

## Running `ccwc` for standard input stream

`ccwc` can work with input coming from standard input stream as well, examples being:

- `echo "This is a sentence" | ccwc -w`
- `cat some_text_file.txt | ccwc -c` 


