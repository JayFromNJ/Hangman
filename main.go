package main

import (
	"Hangman/gojaygo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameStatus struct {
	MissedLetters  []string
	CorrectLetters []string
	SecretWord     string
	GameFinished   bool
}

var HangmanPics = [...]string{`
     +---+
         |
         |
         |
        ===`, `
     +---+
     O   |
         |
         |
        ===`, `
     +---+
     O   |
     |   |
         |
        ===`, `
     +---+
     O   |
    /|   |
         |
        ===`, `
     +---+
     O   |
    /|\  |
         |
        ===`, `
     +---+
     O   |
    /|\  |
    /    |
        ===`, `
     +---+
     O   |
    /|\  |
    / \  |
        ===`,
}

func GetRandomWord() string {
	words := strings.Split("ant baboon badger bat bear beaver camel cat clam cobra cougar coyote crow deer dog "+
		"donkey duck eagle ferret fox frog goat goose hawk lion lizard llama mole monkey moose mouse mule newt otter "+
		"owl panda parrot pigeon python rabbit ram rat raven rhino salmon seal shark sheep skunk sloth snake spider "+
		"stork swan tiger toad trout turkey turtle weasel whale wolf wombat zebra", " ")

	random := gojaygo.NewRandomGenerator()
	return words[random.NextInt(0, len(words)-1)]
}

func StringContains(strArray []string, toFind string) bool {
	for _, a := range strArray {
		if a == toFind {
			return true
		}
	}

	return false
}

func DisplayBoard(missedLetters []string, correctLetters []string, secretWord string) {
	fmt.Println(HangmanPics[len(missedLetters)])

	fmt.Print("Missed letters: ")
	for i := 0; i < len(missedLetters); i++ {
		fmt.Printf("%v ", missedLetters[i])
	}
	fmt.Println()

	displayWord := ""
	lettersInSecretWord := strings.Split(secretWord, "")

	for i := 0; i < len(lettersInSecretWord); i++ {
		if StringContains(correctLetters, lettersInSecretWord[i]) {
			displayWord = displayWord + lettersInSecretWord[i]
		} else {
			displayWord = displayWord + "_"
		}
	}

	fmt.Println(displayWord)
}

func GetGuess(currentGuesses []string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Guess a letter.")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimRight(guess, "\r\n")
		guess = strings.ToLower(guess)

		validGuesses := "abcdefghijklmnopqrstuvwxyz"

		if len(guess) != 1 {
			fmt.Println("Not a valid entry.")
			continue
		} else if StringContains(currentGuesses, guess) {
			fmt.Println("You've already guessed that letter.")
			continue
		} else if strings.Contains(validGuesses, guess) == false {
			fmt.Println("You did not enter a letter.")
			continue
		} else {
			return guess
		}
	}
}

func ResetGame() GameStatus {
	return GameStatus{
		MissedLetters:  []string{},
		CorrectLetters: []string{},
		SecretWord:     GetRandomWord(),
		GameFinished:   false,
	}
}

func main() {
	fmt.Println("H A N G M A N")

	gameStatus := ResetGame()

	for {
		DisplayBoard(gameStatus.MissedLetters, gameStatus.CorrectLetters, gameStatus.SecretWord)

		guess := GetGuess(append(gameStatus.MissedLetters, gameStatus.CorrectLetters...))

		// Check if the player correctly guess a letter
		if StringContains(strings.Split(gameStatus.SecretWord, ""), guess) {
			gameStatus.CorrectLetters = append(gameStatus.CorrectLetters, guess)

			// Check if player has won
			// If all letters in secret word are contained in CorrectLetters, then player wins
			secretLetters := strings.Split(gameStatus.SecretWord, "")
			foundAllLetters := true
			for i := 0; i < len(secretLetters); i++ {
				if StringContains(gameStatus.CorrectLetters, secretLetters[i]) == false {
					foundAllLetters = false
					break
				}
			}

			if foundAllLetters {
				fmt.Printf("YOU WIN! The secret word was: %v\n", gameStatus.SecretWord)
				gameStatus.GameFinished = true
			}
		} else { // Player did not correctly guess a letter
			gameStatus.MissedLetters = append(gameStatus.MissedLetters, guess)

			// Check if the player has lost the game - if there are 6 missed letters, the player loses
			if len(gameStatus.MissedLetters) == 6 {
				DisplayBoard(gameStatus.MissedLetters, gameStatus.CorrectLetters, gameStatus.SecretWord)
				fmt.Println("You have run out of guesses.")
				fmt.Printf("The secret word was: %v\n", gameStatus.SecretWord)
				gameStatus.GameFinished = true
			}
		}

		if gameStatus.GameFinished {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("\nDo you want to play again? (Y or N) ")
			playAgain, _ := reader.ReadString('\n')
			playAgain = strings.TrimRight(playAgain, "\r\n")
			playAgain = strings.ToUpper(playAgain)

			if playAgain == "Y" {
				gameStatus = ResetGame()
			} else {
				break
			}
		}

	}
}
