# Hangman Game in Go using Raylib

A simple, interactive Hangman game built using Go (Golang) and Raylib. This game features an embedded word list, meaning the words are compiled directly into the binary, and there are no external customization options.

## Features

- Random word selection from an embedded list.
- Graphical user interface (GUI) using Raylib.
- Interactive gameplay with keyboard input.
- Limited number of incorrect guesses.
- Visual representation of the Hangman figure.

## Requirements

To run this project, you'll need:

- [Go](https://golang.org/dl/) (version 1.18 or later).
- [Raylib](https://github.com/gen2brain/raylib-go) for Go (raylib-go binding).
- [go:embed](https://golang.org/pkg/embed/) (Go 1.16+).

### Install Raylib for Go

1. Install Raylib using the Go package manager:

   ```bash
   go get -u github.com/gen2brain/raylib-go/raylib
   ```

2. Make sure to have Raylib properly installed. Check the [Raylib-Go installation guide](https://github.com/gen2brain/raylib-go) for platform-specific instructions.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Abhishek-Rauthan/hangman-go-raylib.git
   cd hangman-go-raylib
   ```

2. Install the necessary Go dependencies:

   ```bash
   go mod tidy
   ```

3. Run the game:

   ```bash
   go run main.go
   ```

## How to Play

- The game displays a word with hidden letters, represented by underscores (`_`).
- Guess letters by pressing corresponding keys on your keyboard.
- You are allowed a limited number of incorrect guesses (usually 6), after which the game is over, and the full hangman figure is drawn.
- The game continues until either the word is fully guessed or the number of wrong guesses reaches the maximum.

## Game Flow

1. A random word is selected from an embedded word list.
2. The word is displayed with underscores for each letter.
3. The player enters a letter by pressing a key.
4. If the letter is in the word, the underscores are replaced with the guessed letter.
5. If the letter is not in the word, the number of remaining guesses decreases, and a part of the hangman figure is drawn.
6. The game ends when the player either guesses the word or the hangman figure is fully drawn.

## Acknowledgements

- [Raylib](https://github.com/gen2brain/raylib-go) for the graphical rendering.
- Go programming language for its simplicity and efficiency.

## Contributing

If you'd like to contribute to the project, feel free to fork the repository and submit pull requests. Bug reports, feature requests, and general feedback are also welcome!
