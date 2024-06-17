# Quiz-App

Quiz-App is a simple quiz application built in Go, featuring both a command-line interface (CLI) and a REST API. Users can fetch quiz questions, submit their answers, and compare their performance with others who have taken the quiz.

## Features

- **Command-Line Interface (CLI):**
  - `get`: Fetches the list of quiz questions and their options.
  - `submit`: Submits user answers and displays the number of correct answers and performance.

- **REST API:**
  - `/questions` (GET): Retrieves the list of quiz questions and their options.
  - `/submit` (POST): Submits user answers in JSON format and returns the number of correct answers and performance.

## How It Works

The application is structured with the following components:

- **Core Components:**
  - `quiz_service`: Handles the business logic for fetching questions, submitting answers, and calculating performance.
  - `quiz_repository`: Manages the data for questions, answers, and quiz results.

- **Interfaces:**
  - `QuizInput` and `QuizRepository`: Define interfaces for interacting with quiz data and operations.

- **CLI (Command-Line Interface):**
  - Uses Cobra framework for CLI commands (`get` and `submit`).
  - Retrieves questions from the quiz service and displays them to the user.
  - Accepts user input for answers and submits them, displaying results and performance metrics.

- **REST API:**
  - Built using Gorilla Mux for routing.
  - Provides endpoints (`/questions` and `/submit`) for fetching questions and submitting answers via HTTP requests.
  - Returns responses in JSON format.

## How to Run

To run the Quiz-App, follow these steps:

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/EN2U/quiz-app.git
   cd quiz-app
   ```

2. **Install Dependencies:**

    Ensure you have Go installed. Then, install necessary dependencies:

    ```bash
    go mod tidy
    ```

3. **Run the Command-Line Interface (CLI):**

    To use the CLI:

    ```bash
    go run main.go get
    go run main.go submit [1=1 2=2 ...]
    ```

    Also you can do it with:

    ```bash
    quiz-app get
    quiz-app submit [1=1 2=2 ...]
    ```

4. **Run the REST API Server:**
    ```bash
    go run cmd/quiz_server/main.go
    ```

    The server will start at http://localhost:8080. You can send requests to endpoints /questions and /submit.

    Example:

    ```bash
    curl -X GET http://localhost:8080/questions
    curl -X POST http://localhost:8080/submit -d '{"1":2,"2":1}' -H "Content-Type: application/json"
    ```

5. **Testing:**

    ```bash
    go test ./...
    ```

This will execute all tests in the project, including unit tests and any other tests defined.
