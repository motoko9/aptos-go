package aptosmodule

import "github.com/motoko9/aptos-go/rpcmodule"

func init() {
	// register aptos framework event creator
	rpcmodule.RegisterEventObjectCreator("0x1::coin::WithdrawEvent", WithdrawEventCreator)
	rpcmodule.RegisterEventObjectCreator("0x1::coin::DepositEvent", DepositEventCreator)

	// register aptos framework resource creator
	rpcmodule.RegisterResourceObjectCreator("0x1::account::Account", AccountCreator)
	rpcmodule.RegisterResourceObjectCreator("0x1::coin::CoinInfo", CoinInfoCreator)
	rpcmodule.RegisterResourceObjectCreator("0x1::coin::CoinStore", CoinStoreCreator)
}
