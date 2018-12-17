package loader

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//!+LoadInitialPlayersFromFile loads a player set from a file list
func LoadInitialPlayersFromFile(file string) ([]string, error) {
	fmt.Println("Loading players from ", file)
	inputs, err := os.Open(file)
	if err != nil {
		return make([]string, 0), err
	}
	defer inputs.Close()

	var lines []string
	scanner := bufio.NewScanner(inputs)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	return lines, scanner.Err()
}

//!-LoadInitialPlayersFromFile
