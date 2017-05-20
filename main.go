package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var charIndex []int
	fmt.Printf("Enter the characters needed from your password, separated by space (e.g. 3 7 2):")
	reader := bufio.NewReader(os.Stdin)
	charsString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("couldn't read the characters needed: %v", err)
	}
	for _, x := range strings.Split(strings.TrimSpace(charsString), " ") {
		i, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal("couldn't parse characters, you sure they're numbers?%v", err)
		}
		if i < 1 {
			log.Fatal("counting starts at 1, not 0! try again")
		}
		charIndex = append(charIndex, i)
	}
	fmt.Print("Enter the password (e.g. hunter2): ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("couldn't get password!")
	}
	password := string(bytePassword)
	password = strings.TrimSpace(password)
	for _, x := range charIndex {
		fmt.Printf("%v", string(password[x-1]))
	}
}
