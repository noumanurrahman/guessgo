package main

import (
    "fmt"
    "math/rand"
    "log"
    "os"
)

const rules = "Rules:\n1. a random number is picked\n2. your goal is to find that number\n3. try to minimise the tries you take\n4. you'll know if the number is larger or\nsmaller than the number you've guessed\n5. whoever takes less tries wins the game\n6. press 0 if you want to exit the game\n\n"

func main() {
    fmt.Println("Welcome to Guess The Number game\n")
    fmt.Print(rules)

    p1, p2 := initializePlayers()

    p1Tries := game(p1)
    p2Tries := game(p2)

    if p1Tries < p2Tries {
        fmt.Printf("%v has won the game!\n", p1)
    } else if p1Tries > p2Tries {
        fmt.Printf("%v has won the game!\n", p2)
    } else {
        fmt.Printf("We have to call it a draw!\n")
    }
}

func game(playerName string) int {
    number := rand.Intn(100) + 1
    // fmt.Println(number) // -> For Testing
    var guess int
    tries := 1
    for true {
        message := fmt.Sprintf("[%v](%v) Enter your guess: ", playerName, tries)
        fmt.Print(message)
        _, err := fmt.Scanln(&guess)
        if err != nil { log.Fatal(err) }

        if guess > 100 || guess < 0 {
            fmt.Println("Should be in the range [1-100]")
            continue
        } else if guess == 0 {
            os.Exit(0)
        }

        if guess > number {
            fmt.Println("Try guessing a smaller number")
        } else if guess < number {
            fmt.Println("Try guessing a larger number")
        } else {
            response := fmt.Sprintf("\n%v took %v %v.\n", playerName, tries, tryWord(tries))
            fmt.Println(response)
            break
        }

        tries += 1
    }
    return tries
}

func tryWord(tries int) string {
    if tries==1 {
        return "try"
    } else {
        return "tries"
    }
}

func initializePlayers() (string, string) {
    var player1, player2 string
    fmt.Print("Enter the name of Player 1: ")
    _, err := fmt.Scanln(&player1)
    if err != nil { log.Fatal(err) }
    fmt.Print("Enter the name of Player 2: ")
    _, err = fmt.Scanln(&player2)
    if err != nil { log.Fatal(err) }
    fmt.Println()
    return player1, player2
}