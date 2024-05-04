## Word Count Tool (wc)

This Go program provides functionality similar to the familiar `wc` command found in Unix/Linux systems. It analyzes a file or text input and outputs counts for lines, words, bytes, and characters.

### Usage

Run the program using the following syntax:

```bash
go run main.go [flags] filename
```

**Available Flags:**

- `-l`: Prints the number of newline characters (lines).
- `-c`: Prints the total byte count of the input.
- `-m`: Prints the total character count of the input.

If no flags are specified, the program will display all counts.

**Examples:**

1. Count lines in a file named `file.txt`:

```bash
go run main.go -l file.txt
```

2. Count characters in a string:

```bash
echo "This is a test string." | go run main.go -m
```

### Building

To build an executable named `wc` for your current directory, use the following command:

```bash
go build -o wc main.go
```

This allows you to run the program without the `go run` prefix:

```bash
./wc [flags] filename
```
