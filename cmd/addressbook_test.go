package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	addressbook "github.com/cdyfng/go_eos/addressbook"
	eos "github.com/eoscanada/eos-go"
	"github.com/cdyfng/go_eos/config"
	"github.com/stretchr/testify/assert"
)

func TestAddressUpsert(t *testing.T) {
	api := eos.New("https://mainnet.eoscanada.com")

	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(config.ReadPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)	

	from := eos.AccountName(config.ReadEosAccount_2())
	fmt.Printf("from: %v \n", from)

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := eos.NewTransaction([]*eos.Action{addressbook.Upsert(from, "alice4", "liddell4", 29, "street none4", "123 drink me way4", "wonderland4")}, txOpts)
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


func TestAddressErase(t *testing.T) {
	api := eos.New("https://mainnet.eoscanada.com")

	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(config.ReadPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)	

	from := eos.AccountName(config.ReadEosAccount_2())
	fmt.Printf("from: %v \n", from)

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := eos.NewTransaction([]*eos.Action{addressbook.Erase(from)}, txOpts)
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
