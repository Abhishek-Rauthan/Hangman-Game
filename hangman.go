package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"strings"

	_ "embed"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth  = 800
	windowHeight = 600
	hangmanLife  = 6 // Number of incorrect guesses before losing
)

var (
	//go:embed words.txt
	words       string
	word        = ""
	guessed     = []rune{}
	incorrect   []rune
	life        int
	guessInput  string
	guessLetter rune
)

func initGame() {
	guessed = []rune{}
	incorrect = []rune{}
	life = hangmanLife
	guessInput = ""

	contents := strings.Split(words, "\n")
	max := big.NewInt(int64(len(contents)))
	// Generate a random integer and assign it to the global variable
	randomInt, err := rand.Int(rand.Reader, max)

	if err != nil {
		fmt.Println("Error generating random number:", err)
		return
	}
	n := int(randomInt.Int64())
	word = contents[n]
}

func displayWord() string {
	var display []rune
	for _, letter := range word {
		if contains(guessed, letter) {
			display = append(display, letter)
		} else {
			display = append(display, '_')
		}
	}
	return string(display)
}

func contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func isGameOver() bool {
	return life == 0 || !strings.Contains(displayWord(), "_")
}

func drawHangman(life int) {
	switch life {
	case 5:
		raylib.DrawLine(300, 200, 300, 300, raylib.Black) // post
	case 4:
		raylib.DrawLine(300, 200, 350, 200, raylib.Black) // top beam
	case 3:
		raylib.DrawLine(350, 200, 350, 250, raylib.Black) // rope
	case 2:
		raylib.DrawCircle(350, 270, 20, raylib.Black) // head
	case 1:
		raylib.DrawLine(350, 290, 350, 330, raylib.Black) // body
	case 0:
		raylib.DrawLine(350, 300, 330, 320, raylib.Black) // left arm
		raylib.DrawLine(350, 300, 370, 320, raylib.Black) // right arm
		raylib.DrawLine(350, 330, 330, 350, raylib.Black) // left leg
		raylib.DrawLine(350, 330, 370, 350, raylib.Black) // right leg
	}
}

func main() {

	// Initialize the window
	raylib.InitWindow(windowWidth, windowHeight, "Hangman Game")
	raylib.SetTargetFPS(60)

	// Start the game
	initGame()

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update input
		if raylib.IsKeyPressed(raylib.KeyEnter) && len(guessInput) == 1 {
			guessLetter = rune(guessInput[0])
			if !contains(guessed, guessLetter) {
				guessed = append(guessed, guessLetter)
				if !strings.ContainsRune(word, guessLetter) {
					incorrect = append(incorrect, guessLetter)
					life--
				}
			}
			guessInput = "" // Clear input after guess
		} else if raylib.IsKeyPressed(raylib.KeyBackspace) && len(guessInput) > 0 {
			guessInput = guessInput[:len(guessInput)-1] // Delete last character
		}

		// Handle character input for guessing
		if len(guessInput) < 1 {
			for key := raylib.KeyA; key <= raylib.KeyZ; key++ {
				if raylib.IsKeyPressed(int32(key)) {
					guessInput = string('A' + (key - raylib.KeyA)) // Get input character
					break
				}
			}
		}

		// Draw
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.White)

		// Draw the hangman figure
		drawHangman(life)

		// Draw the word with blanks for unguessed letters
		wordDisplay := displayWord()
		raylib.DrawText(wordDisplay, 300, 150, 40, raylib.Black)

		// Draw incorrect guesses
		incorrectText := fmt.Sprintf("Incorrect guesses: %s", string(incorrect))
		raylib.DrawText(incorrectText, 300, 400, 20, raylib.Black)

		// Display the life count
		raylib.DrawText(fmt.Sprintf("Lives: %d", life), 20, 20, 20, raylib.Black)

		// Display message if game over or won
		if isGameOver() {
			if life == 0 {
				raylib.DrawText("Game Over!", 320, 450, 30, raylib.Red)
			} else {
				raylib.DrawText("You Won!", 320, 450, 30, raylib.Green)
			}
		}

		// Display input box for guessing letters
		raylib.DrawText("Guess a letter:", 300, 500, 20, raylib.Black)
		raylib.DrawText(guessInput, 450, 500, 30, raylib.Blue)

		raylib.EndDrawing()

		// Check if the game is over
		if isGameOver() {
			if life == 0 {
				raylib.DrawText("Press Enter to Restart", 300, 520, 20, raylib.Black)
			} else {
				raylib.DrawText("Press Enter to Restart", 300, 520, 20, raylib.Black)
			}
			if raylib.IsKeyPressed(raylib.KeyEnter) {
				initGame()
			}
		}
	}

	// Close the window and OpenGL context
	raylib.CloseWindow()
}
