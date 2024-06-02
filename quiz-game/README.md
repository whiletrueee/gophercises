# Go Quiz Game with Timer

This program is a command-line quiz game developed in Go. The game reads a series of questions and answers from a CSV file, presents each question to the user, and tracks the number of correct answers given. The game also includes a timer, ensuring that the quiz must be completed within a specified time limit. This project demonstrates the use of Go routines and channels, key concepts in Go's concurrency model.

## Features

- **CSV Input**: The quiz questions and answers are read from a CSV file, allowing easy customization of the quiz content.
- **Timer**: The quiz includes a timer that limits the time a user has to complete the quiz. The game automatically ends when the time runs out.
- **Concurrency**: The program utilizes Go routines and channels to manage the quiz flow and handle user input concurrently with the timer.

## How It Works

1. **CSV Parsing**: The program uses the `flag` package to accept the CSV filename and time limit from the command line. It then reads and parses the CSV file using the `encoding/csv` package.

2. **Data Structuring**: The parsed CSV data is converted into a slice of `problem` structs, each containing a question and its corresponding answer.

3. **Quiz Execution**: The program iterates over the slice of problems, presenting each question to the user and collecting their answer. A correct answer increments the score.

4. **Timer Implementation**: A timer is started when the quiz begins. The `time.NewTimer` function is used to create a timer that sends the current time on its channel after the specified duration.

5. **Concurrency with Channels**: The program uses channels to handle user input and timer expiration concurrently. A Go routine captures user input and sends it to an answer channel. The `select` statement is used to wait for either the timer to expire or an answer to be provided.

6. **Result Display**: The quiz ends either when the timer expires or all questions have been answered. The user's score is then displayed.

## Usage


### Command-Line Flags

- `-csv`: Specifies the path to the CSV file containing quiz questions and answers (default is `problems.csv`).
- `-limit`: Specifies the time limit for the quiz in seconds (default is 30 seconds).

### Running the Program

1. Create a CSV file with your quiz questions and answers.
2. Run the program with the appropriate flags:

```go run main.go -csv=yourfile.csv -limit=60```

If default values are being used then no need to pass them 

```go run main.go```






