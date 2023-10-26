package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/aptos-labs/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/motoko9/aptos-go/rpc"
	"github.com/motoko9/aptos-go/utils"
	"github.com/motoko9/aptos-go/wallet"
	"golang.org/x/crypto/sha3"
	"strings"
	"testing"
	"time"
)

var (
	SignPrefixBytes = []byte("APTOS::RawTransaction")
	EmptyAddress    = "0000000000000000000000000000000000000000000000000000000000000000"
)

func ExtractFromFunction(t string) (string, string, string, error) {
	indexStart := strings.Index(t, "::")
	if indexStart == -1 {
		return "", "", "", fmt.Errorf("type is invalid")
	}
	address := t[0:indexStart]
	//
	t = t[indexStart+2:]
	indexStart = strings.Index(t, "::")
	if indexStart == -1 {
		return "", "", "", fmt.Errorf("type is invalid")
	}
	module := t[0:indexStart]
	//
	indexStart += 2
	function := t[indexStart:]
	return address, module, function, nil
}

func buildArgumentType(t string) *utils.TypeTag__Struct {
	address, module, s, _ := ExtractFromFunction(t)
	moduleAddress := buildModuleAddress(address)
	return &utils.TypeTag__Struct{Value: utils.StructTag{
		Address:  moduleAddress,
		Module:   utils.Identifier(module),
		Name:     utils.Identifier(s),
		TypeArgs: []utils.TypeTag{},
	}}
}

func buildModuleAddress(address string) utils.AccountAddress {
	if strings.Contains(address, "0x") {
		address = address[2:]
	}
	fullAddress := make([]byte, 64)
	copy(fullAddress, EmptyAddress)
	offset := len(fullAddress) - len(address)
	for i := 0; i < len(address); i++ {
		fullAddress[i+offset] = address[i]
	}
	var moduleAddress utils.AccountAddress
	moduleAddressBz, _ := hex.DecodeString(string(fullAddress))
	copy(moduleAddress[:], moduleAddressBz)
	return moduleAddress
}

func TestSerde_Coin_Transfer(t *testing.T) {
	userWallet, err := wallet.NewFromKey("PrivateKey")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	userPubkey := userWallet.PublicKey()
	fmt.Printf("user address: %s\n", userWallet.Address())
	//
	client := New(rpc.MainNet_RPC, true)
	account, aptosErr := client.Account(context.Background(), userAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	typeArg := "0x1::aptos_coin::AptosCoin"
	moduleAddr := "0x1"
	//moduleName := "coin"
	//moduleFunc := "transfer"
	toAddr := ""
	amount := uint64(1)
	// message
	serializer := bcs.NewSerializer()
	// sender
	userAccountAddress := buildModuleAddress(userAddress)
	userAccountAddress.Serialize(serializer)
	// sequence
	serializer.SerializeU64(account.SequenceNumber)
	//
	moduleAccountAddress := buildModuleAddress(moduleAddr)
	//
	tyArgs := make([]utils.TypeTag, 0)
	tyArgs = append(tyArgs, buildArgumentType(typeArg))
	//
	args := make([][]byte, 0)
	{
		temp := bcs.NewSerializer()
		toAccountAddress := buildModuleAddress(toAddr)
		toAccountAddress.Serialize(temp)
		args = append(args, temp.GetBytes())
	}
	{
		temp := bcs.NewSerializer()
		temp.SerializeU64(amount)
		args = append(args, temp.GetBytes())
	}
	payloadBCS := utils.TransactionPayload__EntryFunction{
		Value: utils.EntryFunction{
			Module: utils.ModuleId{
				Address: moduleAccountAddress,
				Name:    "coin",
			},
			Function: "transfer",
			TyArgs:   tyArgs,
			Args:     args,
		},
	}
	payloadBCS.Serialize(serializer)
	// MaxGasAmount
	maxGasAmount := uint64(80000)
	serializer.SerializeU64(maxGasAmount)
	// GasUnitPrice
	gasUnitPrice := uint64(1000000)
	serializer.SerializeU64(gasUnitPrice)
	// ExpirationTimestampSecs
	expirationTimestampSecs := uint64(time.Now().Unix() + 6)
	serializer.SerializeU64(expirationTimestampSecs)
	// ChainId
	chainId := uint8(1)
	serializer.SerializeU8(chainId)
	//
	transactionRaw := serializer.GetBytes()
	//
	hasher := sha3.New256()
	hasher.Write(SignPrefixBytes)
	prefixDigest := hasher.Sum(nil)
	raw := make([]byte, 0)
	raw = append(raw, prefixDigest...)
	raw = append(raw, transactionRaw...)
	// sign
	signature, err := userWallet.Sign(raw)
	if err != nil {
		panic(err)
	}
	// todo
	signedTransaction := utils.SignedTransaction{
		RawTxn: utils.RawTransaction{},
		Authenticator: &utils.TransactionAuthenticator__Ed25519{
			PublicKey: utils.Ed25519PublicKey(userPubkey),
			Signature: utils.Ed25519Signature(signature),
		},
	}
	fmt.Printf("tx hash: %v\n", signedTransaction)
}

func TestSerde_Coin_Transfer1(t *testing.T) {
	userWallet, err := wallet.NewFromKey("1745c2deea6b6756ca757ceac9a3cbc8e02737dad8e99ce6b4ffa52a521aacf5")
	if err != nil {
		panic(err)
	}
	userAddress := userWallet.Address()
	userPubkey := userWallet.PublicKey()
	fmt.Printf("user address: %s\n", userWallet.Address())
	//
	client := New(rpc.TestNet_RPC, true)
	account, aptosErr := client.Account(context.Background(), userAddress, 0)
	if aptosErr != nil {
		panic(aptosErr)
	}
	//
	typeArg := "0x1::aptos_coin::AptosCoin"
	moduleAddr := "0x1"
	//moduleName := "coin"
	//moduleFunc := "transfer"
	toAddr := "0xa44ddb197329f6ea82e95116a2e8f55beeccfdf1ef7c9e367ddc0f05639f4607"
	amount := uint64(10000000)
	//
	tyArgs := make([]utils.TypeTag, 0)
	tyArgs = append(tyArgs, buildArgumentType(typeArg))
	//
	args := make([][]byte, 0)
	{
		temp := bcs.NewSerializer()
		toAccountAddress := buildModuleAddress(toAddr)
		toAccountAddress.Serialize(temp)
		args = append(args, temp.GetBytes())
	}
	{
		temp := bcs.NewSerializer()
		temp.SerializeU64(amount)
		args = append(args, temp.GetBytes())
	}
	payloadBCS := utils.TransactionPayload__EntryFunction{
		Value: utils.EntryFunction{
			Module: utils.ModuleId{
				Address: buildModuleAddress(moduleAddr),
				Name:    "coin",
			},
			Function: "transfer",
			TyArgs:   tyArgs,
			Args:     args,
		},
	}
	rawTransaction := utils.RawTransaction{
		Sender:                  buildModuleAddress(userAddress),
		SequenceNumber:          account.SequenceNumber,
		Payload:                 &payloadBCS,
		MaxGasAmount:            uint64(80000),
		GasUnitPrice:            uint64(100),
		ExpirationTimestampSecs: uint64(time.Now().Unix() + 600),
		ChainId:                 utils.ChainId(2), // 1 - mainnet
	}
	//
	serializer := bcs.NewSerializer()
	rawTransaction.Serialize(serializer)
	rawTransactionData := serializer.GetBytes()
	//
	hasher := sha3.New256()
	hasher.Write(SignPrefixBytes)
	prefixDigest := hasher.Sum(nil)
	raw := make([]byte, 0)
	raw = append(raw, prefixDigest...)
	raw = append(raw, rawTransactionData...)
	// sign
	signature, err := userWallet.Sign(raw)
	if err != nil {
		panic(err)
	}
	//
	signedTransaction := utils.SignedTransaction{
		RawTxn: rawTransaction,
		Authenticator: &utils.TransactionAuthenticator__Ed25519{
			PublicKey: utils.Ed25519PublicKey(userPubkey),
			Signature: utils.Ed25519Signature(signature),
		},
	}
	serializer1 := bcs.NewSerializer()
	signedTransaction.Serialize(serializer1)
	signedTransactionData := serializer1.GetBytes()
	//
	txHash, aptosErr := client.SubmitTransactionBin(context.Background(), signedTransactionData)
	if aptosErr != nil {
		panic(aptosErr)
	}
	fmt.Printf("tx hash: %s\n", txHash)
}
