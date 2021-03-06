package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

const PrivateKeysEnvName = "ETH_PRIVATE_KEYS"
const NodeURLEnvName = "ETH_NODE_URL"

func DecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return b
}

func PrepareEthTransaction(
	ctx context.Context,
	client *ethclient.Client,
	account *Account,
) (*bind.TransactOpts, error) {

	nonce, err := client.PendingNonceAt(ctx, account.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, "getting pending nonce")
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting gas price")
	}

	ethBalance, err := client.BalanceAt(ctx, account.GetAddress(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting balance")
	}

	cost := new(big.Int)
	cost.Mul(gasPrice, big.NewInt(700000))
	if ethBalance.Cmp(cost) < 0 {
		return nil, errors.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	}

	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting network id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(account.GetPrivateKey(), netID)
	if err != nil {
		return nil, errors.Wrap(err, "creating transactor")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func Keccak256(input []byte) [32]byte {
	hash := crypto.Keccak256(input)
	var hashed [32]byte
	copy(hashed[:], hash)

	return hashed
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

func (a *Account) GetAddress() common.Address {
	return a.Address
}

func (a *Account) GetPrivateKey() *ecdsa.PrivateKey {
	return a.PrivateKey
}

func GetAccountFor(accountNo int) (*Account, error) {
	accounts, err := GetAccounts()
	if err != nil {
		return nil, errors.Wrap(err, "getting accounts")
	}
	if accountNo < 0 || accountNo >= len(accounts) {
		return nil, errors.New("account not found")
	}
	return accounts[accountNo], nil
}

// GetAccounts returns a slice of Account from private keys in
// PrivateKeysEnvName environment variable.
func GetAccounts() ([]*Account, error) {
	_privateKeys := os.Getenv(PrivateKeysEnvName)
	privateKeys := strings.Split(_privateKeys, ",")

	// Create an Account instance per private keys.
	accounts := make([]*Account, len(privateKeys))
	for i, pkey := range privateKeys {
		privateKey, err := crypto.HexToECDSA(strings.TrimSpace(pkey))
		if err != nil {
			return nil, errors.Wrap(err, "getting private key to ECDSA")
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		accounts[i] = &Account{Address: publicAddress, PrivateKey: privateKey}
	}
	return accounts, nil
}

func NewClient(ctx context.Context, logger log.Logger) (*ethclient.Client, error) {
	nodeURL := os.Getenv(NodeURLEnvName)

	client, err := ethclient.DialContext(ctx, nodeURL)
	if err != nil {
		return nil, errors.Wrap(err, "create rpc client instance")
	}

	if !strings.Contains(strings.ToLower(nodeURL), "arbitrum") { // Arbitrum nodes doesn't support sync checking.
		// Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.SyncProgress(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "determining if Ethereum client is syncing")
		}
		if s != nil {
			return nil, errors.New("ethereum node is still syncing with the network")
		}
	}

	id, err := client.NetworkID(ctx)
	if err != nil {
		return nil, level.Error(logger).Log("msg", "get nerwork ID", "err", err)
	}

	level.Info(logger).Log("msg", "client created", "netID", id.String())

	return client, nil
}
