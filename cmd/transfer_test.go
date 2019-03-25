package main

import (
	"encoding/hex"
	"encoding/json"

	"fmt"
	"os"
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/token"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAllProcess(t *testing.T){
	 api := eos.New("https://mainnet.eoscanada.com")

	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(readPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)	

	from := eos.AccountName(readEosAccount_2())
	to := eos.AccountName(readEosAccount_1())
	quantity, err := eos.NewEOSAssetFromString("0.0001 EOS")
	memo := "12345"

	assert.NoError(t, err)
	if err != nil {
		panic(fmt.Errorf("invalid quantity: %s", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := eos.NewTransaction([]*eos.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	assert.NoError(t, err)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %s", err))
	}

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %s", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	fmt.Printf("packTx: %v \n", packedTx)

	response, err := api.PushTransaction(packedTx)
	assert.NoError(t, err)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))

}



func readPrivateKey() string {
	// Right now, the key is read from an environment variable, it's an example after all.
	// In a real-world scenario, would you probably integrate with a real wallet or something similar
	envName := "EOS_GO_PRIVATE_KEY"
	privateKey := os.Getenv(envName)
	if privateKey == "" {
		panic(fmt.Errorf("private key environment variable %q must be set", envName))
	}

	return privateKey
}

func readEosAccount_1() string {
	envName := "EOS_GO_ACCOUNT_1"
	account_1 := os.Getenv(envName)
	if account_1 == "" {
		panic(fmt.Errorf("EOS_GO_ACCOUNT_1 environment variable %q must be set", envName))
	}

	return account_1
}


func readEosAccount_2() string {
	envName := "EOS_GO_ACCOUNT_2"
	account_2 := os.Getenv(envName)
	if account_2 == "" {
		panic(fmt.Errorf("EOS_GO_ACCOUNT_2 environment variable %q must be set", envName))
	}

	return account_2
}