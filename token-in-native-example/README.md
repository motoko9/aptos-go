# Move example on Aptos

* Aptos Env
* Move Module by Aptos CLI
* Move Module by RPC
* Publish
* Call module

## Aptos Env

First, you need to [install Aptos-core](https://aptos.dev/guides/getting-started) & [install Aptos command line tool](https://aptos.dev/cli-tools/aptos-cli-tool/install-aptos-cli)

Second, you need to config Aptos command line tool, reference [Use Aptos CLI](https://aptos.dev/cli-tools/aptos-cli-tool/use-aptos-cli)

## Move Module by Aptos CLI

You can create account, query account, build & publish & running Move function with Aptos CLI, reference[Use Aptos CLI](https://aptos.dev/cli-tools/aptos-cli-tool/use-aptos-cli)

## Move Module by RPC

## Publish

[source code](https://github.com/motoko9/aptos-program-library/tree/master/token-in-native/contract)

```
aptos move publish --package-dir . --named-addresses NamedAddr=0x1685cdc9a83c3da34c59208f34bddb3217f63bfbe0c393f04462d1ba06465d08
```

```
package size 494 bytes
{
  "Result": {
    "transaction_hash": "0x46bfa491acd2deaaf0a46554e04e57e979ca6c7d4cfda74c2efa8f2409dc686a",
    "gas_used": 62,
    "gas_unit_price": 1,
    "sender": "1685cdc9a83c3da34c59208f34bddb3217f63bfbe0c393f04462d1ba06465d08",
    "sequence_number": 0,
    "success": true,
    "timestamp_us": 1661656715311642,
    "version": 11307262,
    "vm_status": "Executed successfully"
  }
}
```

## Call module

* [Publish Coin Module](./coin_module_publish_test.go)
* [Initialize Coin Module](./coin_initialize_test.go)
* [Query Coin info](./coin_info_test.go)
* [Register Coin Recipient](./register_recipient_test.go)
* [Mint Coin](./coin_mint_test.go)
* [Transfer Coin](./coin_transfer_test.go)

