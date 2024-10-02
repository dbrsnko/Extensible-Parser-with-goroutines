# 📦 Extensible Parser with Goroutines
## 🚀 Getting Started
To run this application, navigate to its root directory and execute the following command:
**go run .**

## 💡 Idea
The goal is to make the solution as extensible and
flexible as possible, even beyond common sense. Here’s what it should be:

### 1. File Input
Read multiple files, each containing JSON objects of the following structure, separated by **newlines** like this:
{"message":"test","timestamp":"2023-03-15T21:54:42.123Z"}

{"message":"qwerty","timestamp":"2023-03-16T21:54:42.123Z"}

{"message":"ytrewq","timestamp":"2023-03-17T21:54:42.123Z"}

{"message":"asdfg","timestamp":"2023-03-18T21:54:42.123Z"}

{"message":"gfdsa","timestamp":"2023-03-19T21:54:42.123Z"}


### 2. Parsing and Output
Parse each line.
Wait for a random number of seconds (less than 5).
Write the parsed data to the console and the corresponding output file.

### 3. Error Handling
If parsing or writing fails, gracefully stop all goroutines and print the error to the terminal:
Important: Panic, os.Exit, etc., are forbidden.
All successfully processed lines prior to the error will be saved to the appropriate file and displayed in the terminal.

### 4. Concurrency
Execute steps 1, 2, and 3 concurrently to maximize efficiency.
