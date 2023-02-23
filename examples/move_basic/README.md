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

[source code](https://github.com/motoko9/aptos-program-library/tree/master/helloworld/contract)

```
aptos move publish --package-dir . --named-addresses NamedAddr=0x5e7f8779c8c26ec3cbba37337142b2aaa2291b4779f4b386a0de83da177df510
```

```
package size 1660 bytes
{
  "Result": {
    "transaction_hash": "0xec24b2c091d0c02a55a2caa223e9883fa10bfb62a920967b0b694dcebe14e61b",
    "gas_used": 189,
    "gas_unit_price": 1,
    "sender": "5e7f8779c8c26ec3cbba37337142b2aaa2291b4779f4b386a0de83da177df510",
    "sequence_number": 0,
    "success": true,
    "timestamp_us": 1661650597217213,
    "version": 10901294,
    "vm_status": "Executed successfully"
  }
}
```

## Call module

* [publish Move Module](./move_publish_test.go)
* [Read Move Module](./move_read_test.go)
* [Write Move Module](./move_write_test.go)

