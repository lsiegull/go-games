package main

import ("fmt"
"math/rand"
"time"
)

func main() {
  var nextRound string
  var difficulty int
  score := 0
  fmt.Println("Welcome to the secret number guessing game!")
  fmt.Printf("Please enter a number that corresponds to the difficulty level you prefer-- Easy (1), Medium (2), Hard(3): ")
  fmt.Scanln(&difficulty)
  for nextRound != "q" {
    rand.Seed(time.Now().UnixNano())
    var min int
    var max int
    fmt.Printf("Please select a minimum value for this round's secret number: ")
    fmt.Scanln(&min)
    fmt.Printf("Please select a maximum value for this round's secret number: ")
    fmt.Scanln(&max)
    secretNum := rand.Intn(max - min) + min
    var guess int
    var guesses int
    if difficulty == 1 {
      guesses = (max - min)/5
    } else if difficulty == 2 {
      guesses = (max - min)/10
    } else {
      guesses = (max - min)/20
    }
    for (guess != secretNum) && (guesses > 0){
      fmt.Printf("You have %d guesses left\n", guesses)
      fmt.Printf("Please enter a guess: ")
      fmt.Scanln(&guess)
      if guess > secretNum {
        fmt.Println("Too high!")
      }
      if guess < secretNum {
        fmt.Println("Too low!")
      }
      guesses--
    }
    if guess == secretNum {
      fmt.Println("You won!")
      score++
    } else {
      fmt.Println("You lost!")
    }
    if score == 1 {
      fmt.Printf("You have won %d round.\n", score)
    } else {
      fmt.Printf("You have won %d rounds.\n", score)
    }
    fmt.Printf("Type q to quit, or press any other key to play another round: ")
    fmt.Scanln(&nextRound)
  }
  fmt.Println("Thanks for playing!")
}
