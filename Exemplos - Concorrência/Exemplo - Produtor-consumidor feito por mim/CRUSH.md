# CRUSH: The Agent's Guide to this Go Producer-Consumer Codebase

This document provides essential information for an AI agent to effectively understand, modify, and contribute to this Go project.

## 1. Project Overview

This is a simple command-line application written in Go that demonstrates the **producer-consumer** concurrency pattern. It simulates a pizzeria where a "producer" (`pizzeria` goroutine) creates pizza orders and a "consumer" (in the `main` function) processes them.

The application uses the `github.com/charmbracelet/lipgloss` library for styled terminal output.

## 2. Essential Commands

As there are no Makefiles or script files, use standard Go commands:

- **Run the application:**
  ```bash
  go run main.go
  ```
  or
  ```bash
  go run .
  ```

- **Build the application:**
  ```bash
  go build
  ```

- **Run tests (no tests currently exist):**
  ```bash
  go test ./...
  ```

- **Tidy dependencies:**
  ```bash
  go mod tidy
  ```

## 3. Code Organization and Structure

The entire application is contained within `main.go`.

- **`main()` function:**
  - Initializes the random number generator.
  - Sets up the `Producer`.
  - Starts the `pizzeria` producer in a separate goroutine.
  - The consumer logic will be implemented here to receive `PizzaOrder`s from the producer.

- **`Producer` struct:**
  - `data chan PizzaOrder`: A channel to send pizza orders from the producer to the consumer.
  - `quit chan chan error`: A channel used for graceful shutdown of the producer goroutine.

- **`PizzaOrder` struct:**
  - Represents the data passed between producer and consumer. Contains details about the pizza order.

- **`pizzeria()` function:**
  - Acts as the producer.
  - It runs in an infinite loop, intended to create `PizzaOrder`s and send them to the `data` channel. The core production logic is not yet implemented inside the loop.

## 4. Key Patterns and Conventions

- **Concurrency:** The core pattern is producer-consumer using goroutines and channels. The `pizzeria` runs as a separate goroutine, and the `main` function's main goroutine acts as the consumer.
- **Channels for Communication:** All communication between the producer and consumer happens through the `data` channel (`chan PizzaOrder`). This is the idiomatic way to handle data transfer between goroutines in Go.
- **Graceful Shutdown:** The `Producer.Close()` method implements a graceful shutdown pattern. It uses a `quit` channel to signal the `pizzeria` goroutine to stop its work and return any final error. This is a robust pattern for managing the lifecycle of goroutines.

## 5. Gotchas and Important Notes

- **Incomplete Logic:** The `pizzeria` function contains an empty infinite `for` loop. The core logic for creating pizzas and handling the `quit` signal needs to be implemented there.
- **Missing Consumer:** The consumer logic in the `main` function is not yet implemented. An agent will need to implement a loop to receive from the `pizzaJob.data` channel.
- **Random Number Seeding:** The code uses `rand.Seed(time.Now().UnixNano())`. While this works, modern Go (1.20+) automatically seeds the global random number generator in the `math/rand` package. For this project, the existing implementation is sufficient, but be aware of modern best practices.
