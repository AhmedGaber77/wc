## Build Your Own wc Tool

This repository contains my implementation of the Unix command line tool `wc` in [insert programming language]. Through this project, I learned about writing simple, modular programs with clean interfaces and how to design programs to be easily connected to other tools.

### What I Learned

- **Unix Philosophy**: Understanding the importance of writing simple parts connected by clean interfaces.
- **Functional Requirements**: Implementing word, line, character, and byte count functionalities.
- **Handling Command Line Options**: Supporting various command line options (-c, -l, -w, -m) and the default behavior.
- **Reading from Standard Input**: Enabling the tool to read from standard input if no filename is specified.

### Challenges Faced

- **Modular Design**: Structuring the code to ensure each function handles one task efficiently.
- **Option Parsing**: Parsing command line options and determining the appropriate action for each.
- **Locale Considerations**: Handling multibyte characters for accurate character count (-m option).
- **Error Handling**: Ensuring robust error handling to handle invalid inputs and edge cases.

### Conclusion

Building this `wc` tool was a rewarding challenge that deepened my understanding of software engineering principles and command line tool development. It provided valuable insights into modular design, option parsing, and handling different input sources.



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
