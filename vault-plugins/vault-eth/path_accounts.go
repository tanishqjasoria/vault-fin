package vault_eth

import (
	"context"
	wallet "../ethereum-utils"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func AccountPaths(backend *backend) []*framework.Path {

	return []*framework.Path {
		// Path to list all the created accounts
		{
			Pattern: "accounts/",
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.ListOperation: backend.ListAccounts,
			},
			HelpSynopsis: "List all the ethereum accounts",
			HelpDescription: "Everything would be listed",
		},

		// Import a new account at the given name path
		{
			Pattern: "accounts/import/" + framework.GenericNameRegex("name"),
			HelpSynopsis: "Import an Ethereum account using a provided passphrase.",
			HelpDescription: `
					Creates (or updates) an Ethereum account: an account controlled by a private key. Also
					The generator produces a high-entropy passphrase with the provided length and requirements.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name": {Type: framework.TypeString},
				"mnemonic": {
					Type:        framework.TypeString,
					Default:     Empty,
					Description: "The mnemonic to use to create the account.",
				},
				"index": {
					Type:        framework.TypeInt,
					Description: "The index used in BIP-44.",
					Default:     0,
				},
				"blacklist": {
					Type:        framework.TypeCommaStringSlice,
					Description: "The list of accounts that this account can't send transactions to.",
				},
				"whitelist": {
					Type:        framework.TypeCommaStringSlice,
					Description: "The list of accounts that this account can send transactions to.",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				//logical.ReadOperation:   b.AccountRead,
				logical.CreateOperation: backend.AccountCreate,
				logical.UpdateOperation: backend.AccountUpdate,
				logical.DeleteOperation: backend.AccountDelete,
			},
		},
		// Create a new account at the given name path
		{
			Pattern: "accounts/" + framework.GenericNameRegex("name"),
			HelpSynopsis: "Create an Ethereum account using a generated passphrase.",
			HelpDescription: `
					Creates (or updates) an Ethereum account: an account controlled by a private key. Also
					The generator produces a high-entropy passphrase with the provided length and requirements.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name": {Type: framework.TypeString},
				"mnemonic": {
					Type:        framework.TypeString,
					Default:     Empty,
					Description: "The mnemonic to use to create the account. If not provided, one is generated.",
				},
				"index": {
					Type:        framework.TypeInt,
					Description: "The index used in BIP-44.",
					Default:     0,
				},
				"blacklist": {
					Type:        framework.TypeCommaStringSlice,
					Description: "The list of accounts that this account can't send transactions to.",
				},
				"whitelist": {
					Type:        framework.TypeCommaStringSlice,
					Description: "The list of accounts that this account can send transactions to.",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.ReadOperation:   backend.AccountRead,
				logical.CreateOperation: backend.AccountCreate,
				logical.UpdateOperation: backend.AccountUpdate,
				logical.DeleteOperation: backend.AccountDelete,
			},
		},
		//
		{
			Pattern:	"accounts/" + framework.GenericNameRegex("name") + "/transfer",
			HelpSynopsis: "Send ETH from an account.",
			HelpDescription: `
					Send ETH from an account.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name": {Type: framework.TypeString},
				"to": {
					Type:	framework.TypeString,
					Description: "The address of the wallet to send ETH to.",
				},
				"amount": {
					Type:	framework.TypeString,
					Description: "Amount of ETH (in wei).",
				},
				"gas_limit": {
					Type:	framework.TypeString,
					Description: "The gas limit for the transaction - defaults to 21000.",
				},
				"gas_price": {
					Type:	framework.TypeString,
					Description: "The gas price for the transaction in wei.",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.UpdateOperation: backend.Transfer,
				logical.CreateOperation: backend.Transfer,
			},
		},
		//
		{
			Pattern:	"accounts/" + framework.GenericNameRegex("name") + "/balance",
			HelpSynopsis: "Return the balance for an account.",
			HelpDescription: `
					Return the balance in wei for an address.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name":    {Type: framework.TypeString},
				"address": {Type: framework.TypeString},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.ReadOperation: backend.ReadBalance,
			},
		},
		//
		{
			Pattern:	"accounts/" + framework.GenericNameRegex("name") + "/sign-tx",
			HelpSynopsis: "Sign a transaction.",
			HelpDescription: `
					Sign a transaction.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name":    {Type: framework.TypeString},
				"address": {Type: framework.TypeString},
				"to": {
					Type:        framework.TypeString,
					Description: "The address of the wallet to send ETH to.",
				},
				"data": {
					Type:        framework.TypeString,
					Description: "The data to sign.",
				},
				"encoding": {
					Type:        framework.TypeString,
					Default:     "utf8",
					Description: "The encoding of the data to sign.",
				},
				"amount": {
					Type:        framework.TypeString,
					Description: "Amount of ETH (in wei).",
				},
				"nonce": {
					Type:        framework.TypeString,
					Description: "The transaction nonce.",
				},
				"gas_limit": {
					Type:        framework.TypeString,
					Description: "The gas limit for the transaction - defaults to 21000.",
					Default:     "21000",
				},
				"gas_price": {
					Type:        framework.TypeString,
					Description: "The gas price for the transaction in wei.",
					Default:     "0",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.CreateOperation: backend.SignTxn,
				logical.UpdateOperation: backend.SignTxn,
				logical.ReadOperation: backend.SignTxn,
			},
		},
		//
		{
			Pattern:	"accounts/" + framework.GenericNameRegex("name") + "/deploy",
			HelpSynopsis: "Deploy a smart contract from an account.",
			HelpDescription: `
					Deploy a smart contract to the network.
			`,
			Fields: map[string]*framework.FieldSchema{
				"name":    {Type: framework.TypeString},
				"address": {Type: framework.TypeString},
				"version": {
					Type:        framework.TypeString,
					Description: "The smart contract version.",
				},
				"abi": {
					Type:        framework.TypeString,
					Description: "The contract ABI.",
				},
				"bin": {
					Type:        framework.TypeString,
					Description: "The compiled smart contract.",
				},
				"gas_limit": {
					Type:        framework.TypeString,
					Description: "The gas limit for the transaction - defaults to 0 meaning estimate.",
					Default:     "0",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.UpdateOperation: backend.DeployContract,
				logical.CreateOperation: backend.DeployContract,
			},
		},
		////////////////////////////////////////////
		{
			Pattern:	"accounts/" + framework.GenericNameRegex("name") + "/sign",
			HelpSynopsis: "Sign a message",
			HelpDescription: `
					Sign calculates an ECDSA signature for:
					keccack256("\x19Ethereum Signed Message:\n" + len(message) + message).
					
					https://eth.wiki/json-rpc/API#eth_sign
		`,
			Fields: map[string]*framework.FieldSchema{
				"name": {Type: framework.TypeString},
				"message": {
					Type:        framework.TypeString,
					Description: "Message to sign.",
				},
			},
			ExistenceCheck: CheckIfExist,
			Callbacks: map[logical.Operation]framework.OperationFunc{
				logical.CreateOperation: backend.SignMessage,
				logical.UpdateOperation: backend.SignMessage,
			},
		},
	}
}


func (backend *backend) ListAccounts(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error) {

	accounts, err := request.Storage.List(ctx, "accounts/")
	if err != nil {
		return nil, err
	}
	return logical.ListResponse(accounts), nil
}

func (backend *backend) AccountCreate(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) AccountUpdate(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) AccountRead(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) AccountTransfer(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) AccountDelete(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) ReadBalance(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) SignTxn(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) DeployContract(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) SignMessage(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}

func (backend *backend) Transfer(ctx context.Context, request *logical.Request, data *framework.FieldData) (*logical.Response, error){
	// Your Code Goes Here!!
	return logical.ListResponse(" "), nil
}



