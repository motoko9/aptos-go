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

[source code](https://github.com/motoko9/aptos-program-library/tree/master/token/contract)

```
aptos move publish --package-dir . --named-addresses NamedAddr=0xf0881fc180ccb250f3e748730f03a17fb627d824f0a23cf934873304195b609a
```

```
package size 3652 bytes
{
  "Result": {
    "transaction_hash": "0xb7224ef9e2bbe451e8ab7965e1f727ead5643bc394f61ed05bdaff9ee9d17992",
    "gas_used": 406,
    "gas_unit_price": 1,
    "sender": "f0881fc180ccb250f3e748730f03a17fb627d824f0a23cf934873304195b609a",
    "sequence_number": 0,
    "success": true,
    "timestamp_us": 1661652056334521,
    "version": 10990987,
    "vm_status": "Executed successfully"
  }
}
```

## Call module

* [deploy coin](../../coin-example/coin_publish_test.go)
* [initialize coin](../../coin-example/coin_initialize_test.go)
* [register recipient](../../coin-example/register_recipient_test.go)
* [mint coin to recipient](../../coin-example/mint_test.go)



* [Publish Coin Module](./coin_module_publish_test.go)
* [Initialize Coin Module](./coin_initialize_test.go)
* [Query Coin info](./coin_info_test.go)
* [Register Coin Recipient](./register_recipient_test.go)
* [Mint Coin](./coin_mint_test.go)
* [Transfer Coin](./coin_transfer_test.go)

