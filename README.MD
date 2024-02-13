
# ccwc - Coding Challenges Word Count

This is a simple command-line tool called ccwc (Coding Challenges Word Count) that takes the `-c` option and outputs the number of bytes in a file.

## Prerequisites

Before using ccwc, ensure that you have the following:

1. **Test File**: You need to have a test file named `test.txt` in the same directory as the `main.go` file.

2. **Go Installation**: Make sure you have Go installed on your system. You can download and install Go from [here](https://golang.org/dl/).

## Build

To build the ccwc binary, navigate to the directory containing the `main.go` file and run the following command:

```bash
go build -o ccwc main.go
```

This command will compile the `main.go` file and generate an executable binary named `ccwc`.

## Installation

To make ccwc available as a command from any directory, you can move the compiled binary to a directory that is included in your system's `PATH`. For example, you can use the following command to move it to `/usr/local/bin`:

```bash
sudo mv ccwc /usr/local/bin
```

## Usage

Once ccwc is installed, you can use it to count the number of bytes in a file. Here's the syntax:

```bash
ccwc -c test.txt
```

Replace `test.txt` with the name of the file you want to analyze.

