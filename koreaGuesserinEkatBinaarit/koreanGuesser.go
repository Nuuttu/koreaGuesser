// Tuomo Miettinen
// http://terokarvinen.com/2020/go-programming-course-2020-w22/
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// https://www.geeksforgeeks.org/fmt-scanln-function-in-golang-with-examples/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var korean []string

func roundKorean() {
	file, err := os.Open("./roundKorean.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		korean = append(korean, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var vastaukset []string

func roundVastaukset() {
	file, err := os.Open("./roundVastaukset.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vastaukset = append(vastaukset, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func tulostaMerkit() {
	file, err := os.Open("./letters.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var oikeatVastaukset int
var kaikkiVastaukset int

func round() {

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(40)

	fmt.Println("\nEnter 'q' to stop and return")
	fmt.Println(" | How is this")
	fmt.Println(" ↓")
	fmt.Println(" ", korean[r])
	fmt.Println(" ↑")
	fmt.Println(" | pronounced? Write now...")

	var vastaus string
	fmt.Scanln(&vastaus)
	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")

	fmt.Println("\nYou answered: ", vastaus)

	switch vastaus {
	case vastaukset[r]:
		fmt.Println("\n----- You guessed right! -----\n")
		oikeatVastaukset++
	case "q":
		ending(oikeatVastaukset, kaikkiVastaukset)
		return
	default:
		fmt.Println("\n----- Wrong! -----")
		fmt.Println("The right answer for : '", korean[r], "' was: '", vastaukset[r], "'")
	}
	kaikkiVastaukset++
	round()
}

func start() {

	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")
	fmt.Println("What would you like to do?")
	fmt.Println("Enter 'start' to start the guessing game")
	fmt.Println("Enter 'print' to open pronunciation list")
	fmt.Println("Enter 'exit' to exit")

	var vastaus string
	fmt.Scanln(&vastaus)
	fmt.Println("You chose: ", vastaus)
	fmt.Println("____________________________________________")

	switch vastaus {
	case "start":
		oikeatVastaukset = 0
		kaikkiVastaukset = 0
		round()
	case "print":
		tulostaMerkit()
	case "exit":
		return
	default:
		fmt.Println("Not a valid argument.")
	}

	start()
}

func ending(oikeat int, kaikki int) {

	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")
	fmt.Println("Your score was", oikeat, " out of ", kaikki)
	fmt.Println("")
	fmt.Println("\nPress Enter to continue...")

	var vastaus string
	fmt.Scanln(&vastaus)

}

func main() {

	roundKorean()
	roundVastaukset()
	start()

	fmt.Println("Thanks for running")
}
