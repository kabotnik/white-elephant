package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kabotnik/white-elephant/pkg/loader"
)

var gameExpressions = [...]string{
	"Will you take a safe option and steal? Or will you live on the edge?",
	"Steal a gift or choose from the pile!",
	"You got this!",
	"I believe in you!",
	"What's it gonna be, friend?",
	"Hey there, pal.",
	"Go get 'em, tiger!",
	"Aw man, this gonna be gud!",
	"I was wondering when you were going go show up!",
	"Welcome to the party!",
	"Now we're in trouble.",
	"(╯°□°）╯︵ ┻━┻",
	"┬─┬ノ( º _ ºノ)",
}

var (
	playCmd = &cobra.Command{
		Use:   "play",
		Short: "Play will run a game of white elephant",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			playerFile := viper.GetString("players")

			players, err := loader.LoadInitialPlayersFromFile(playerFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			var remainingPlayers []string = players
			var round int = 1

			callClear()

			for ok := true; ok; ok = (len(remainingPlayers) > 0) {
				fmt.Println("Round: ", round)
				fmt.Println("Players remaining:", len(remainingPlayers))
				fmt.Print(SSlice(remainingPlayers))

				_, err := stdinContinue("Hit <Enter> to choose the next player")
				if err != nil {
					fmt.Println(err)
				}

				s := rand.NewSource(time.Now().UnixNano())
				r := rand.New(s)

				p := r.Intn(len(remainingPlayers))
				fmt.Printf("\n%s it's your turn! %s\n", remainingPlayers[p], gameExpressions[r.Intn(len(gameExpressions))])

				remainingPlayers = takeTurn(remainingPlayers, p)

				_, err = stdinContinue("Hit <Enter> to continue")
				if err != nil {
					fmt.Println(err)
				}

				round++
				callClear()
			}

			fmt.Println("No more players!")
		},
	}
)

var clear map[string]func()

type SSlice []string

func init() {
	rootCmd.AddCommand(playCmd)
	playCmd.Flags().StringP("players", "p", "", "File containing all players for the game")
	viper.BindPFlag("players", playCmd.Flags().Lookup("players"))

	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func takeTurn(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func stdinContinue(prompt string) (string, error) {
	var b []byte
	fmt.Fprintf(os.Stderr, prompt)
	b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	return string(b), err
}

func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func (s SSlice) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}
