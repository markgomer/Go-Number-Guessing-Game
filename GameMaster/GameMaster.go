package GameMaster

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func RunGame() {
    fmt.Println("GUESS YAR NUMBAR")
    fmt.Println("I'm gonna pick a numbar from 0 to 100 and ya hav to guess!")
    fmt.Println("Other wise yer life's ovar!")

    const numbarOfTries = 10
    guesses := [numbarOfTries]int64{}

    // Create reader from std input
    scanner := bufio.NewScanner((os.Stdin))

    secret := rand.Int64N(101)

    // check if guess is alright
    // limit to the number of tries
    for i := range numbarOfTries {
        // read the guess
        fmt.Printf(
            "TELL YAR GUESS ALREADY!\n" +
            "YA STILL HAV %d GUESSES\n",
            numbarOfTries - i,
        )
        scanner.Scan()

        guess := scanner.Text()
        // clean up dumb spaces and convert to int
        guess = strings.TrimSpace(guess)
        guessInt, err := strconv.ParseInt(guess, 10, 64)

        // check if the conversion had an error
        if err != nil {
            fmt.Println("DA GAME IS ABOUT NUMBARS! ENTER NUMBERS")
            return
        }
        switch {
        case guessInt < secret:
            fmt.Println("WRONG. THE NUMBAR IS GREATER THAN", guessInt)
        case guessInt > secret:
            fmt.Println("WRONG. THE NUMBAR IS LESSAR THAN", guessInt)
        case guessInt == secret:
            fmt.Printf(
                "WAAA!? YA GOT IT AIGHT. TWAS %d\n" +
                "YA GOT DIS IN %d TRIES\n" +
                "THOSE WERE YER TRIES: %v\n",
                secret, i+1, guesses[:i],
            )
            return
        }
        guesses[i] = guessInt
    }
    fmt.Printf(
        "PFFFFFT YA LOSE DA GAME. THE NUMBAR WAS SO EEZ: %d\n" +
        "YA HAD %d TRIES AND EVEN WITH ALL THAT YA LOST PFFFFT" +
        "THOSE WERE YER TRIES: %v\n",
        secret, numbarOfTries, guesses,
    )
}
