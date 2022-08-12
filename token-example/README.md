# Coin example on Aptos

* Aptos Env
* Coin Example

## Aptos Env

First, you need to [install Aptos-core](https://aptos.dev/guides/getting-started) & [install Aptos command line tool](https://aptos.dev/cli-tools/aptos-cli-tool/install-aptos-cli)

Second, you need to config Aptos command line tool, reference [Use Aptos CLI](https://aptos.dev/cli-tools/aptos-cli-tool/use-aptos-cli)

## Coin Example
[source code](https://github.com/aptos-labs/aptos-core/tree/main/aptos-move/move-examples/moon_coin)

### Bytecode
```bash
tangaoyuan@tangaoyuandeMacBook-Pro aptos-core % aptos move compile --package-dir ./aptos-move/move-examples/moon_coin --named-addresses MoonCoinType=0xa0db31e3cc6f597ec084d7fbcf4cb562522ab83d2fe7f3567af79f85627fcd9c
{
  "Result": [
    "A0DB31E3CC6F597EC084D7FBCF4CB562522AB83D2FE7F3567AF79F85627FCD9C::moon_coin"
  ]
}
```
After compile, there is a build output in package directory. You can get codes in package_dir/build/Examples/bytecode_modules.

### Move RPC

* [deploy coin](coin_publish_test.go)
* [initialize coin](coin_initialize_test.go)
* [register recipient](register_recipient_test.go)
* [mint coin to recipient](mint_test.go)

In other networks, since tokens/coins are just balance numbers in a contract, anyone can "send" anyone else a random coin, even if the recipient doesn't want it. In Aptos, a user needs to explicitly register to receive a Coin<RandomCoin> before it can be sent to them.

