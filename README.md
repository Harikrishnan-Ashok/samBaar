# sambaar

`sambaar` is a Go-based application designed to run from the command line like a typical Linux utility.

---

## Prerequisites

Before you begin, make sure you have the following:

- [Go] programming language installed (version 1.20 or higher is recommended)
- A directory in your system `PATH` where the executable will be installed (e.g., `/usr/local/bin` or `$HOME/bin`)

---

## Building the Application

To compile the application into an executable named `sambaar`, run the following command from the project directory:

---

## Installing

### Option 1: System-wide install (requires sudo)

Move the compiled binary to a directory in your system `PATH`:

Then, ensure it has the correct permissions:

```bash
go build -o sambaar main.go
sudo mv sambaar /usr/local/bin/
chmod +x /usr/local/bin/sambaar
```

## Running the Application

Once installed, you can run `sambaar` from any terminal:

```bash
sambaar
```

---

## Development Workflow

To quickly test changes during development without building an executable, use:

```bash
go run main.go
```

This compiles and runs the application in one step.

---

## For Custom ShortCuts or alias

the default config file is read from ~/.config/samBaar/config.json file

refer the example_config.json file for reference
