package config

import (
	"os"
	"fmt"
	)
func ReadPrivateKey() string {
	// Right now, the key is read from an environment variable, it's an example after all.
	// In a real-world scenario, would you probably integrate with a real wallet or something similar
	envName := "EOS_GO_PRIVATE_KEY"
	privateKey := os.Getenv(envName)
	if privateKey == "" {
		panic(fmt.Errorf("private key environment variable %q must be set", envName))
	}

	return privateKey
}

func ReadEosAccount_1() string {
	envName := "EOS_GO_ACCOUNT_1"
	account_1 := os.Getenv(envName)
	if account_1 == "" {
		panic(fmt.Errorf("EOS_GO_ACCOUNT_1 environment variable %q must be set", envName))
	}

	return account_1
}


func ReadEosAccount_2() string {
	envName := "EOS_GO_ACCOUNT_2"
	account_2 := os.Getenv(envName)
	if account_2 == "" {
		panic(fmt.Errorf("EOS_GO_ACCOUNT_2 environment variable %q must be set", envName))
	}

	return account_2
}