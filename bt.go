package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

func main() {
	// Set the parameters for the main network.
	params := &chaincfg.MainNetParams

	// Prompt the user for a password.
	fmt.Print("Enter password: ")
	password, err := readPassword()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a new master key using the password and the desired parameters.
	masterKey, err := hdkeychain.GenerateMasterKey([]byte(password), params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a new address from the master key.
	address, err := masterKey.Address(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the address.
	fmt.Println(address)
}

// readPassword reads a password from the user's input.
func readPassword() (string, error) {
	password, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return password[:len(password)-1], nil
}
