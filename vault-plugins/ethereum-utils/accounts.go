package ethereum_utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	wallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Account struct {
	Mnemonic        []string                 `json:"mnemonic"`
	Address         string                   `json:"address"`
	PrivateKey      string                   `json:"key"`
	Whitelist       map[*common.Address]bool `json:"whitelist"`
	Blacklist       map[*common.Address]bool `json:"blacklist"`
	IsInitialized   bool                     `json:"is_initialized"`
	EnableWhitelist bool                     `json:"enable_whitelist"`
}

type Transaction struct {
	Nonce    uint64          `json:"nonce"`
	Address  *common.Address `json:"address"`
	Amount   *big.Int        `json:"amount"`
	GasPrice *big.Int        `json:"gas_price"`
	GasLimit uint64          `json:"gas_limit"`
}


func (account *Account) Validate(address *common.Address) error {

	if account.Blacklist[address] {
		return fmt.Errorf("Blacklisted address: %s", address.Hex())
	}

	if account.EnableWhitelist {
		if !account.Whitelist[address] {
			return fmt.Errorf("Not Whitelisted: %s:", address.Hex())
		}
	}

	return nil
}

func (account *Account) SignTxn(txn *Transaction) error {return nil}









