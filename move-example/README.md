# Move example on Aptos

* Aptos Env
* Move Module by Aptos CLI
* Move Module by RPC
* Publishing
* Reading resource
* Modifying resource

## Aptos Env

First, you need to [install Aptos-core](https://aptos.dev/guides/getting-started) & [install Aptos command line tool](https://aptos.dev/cli-tools/aptos-cli-tool/install-aptos-cli)

Second, you need to config Aptos command line tool, reference [Use Aptos CLI](https://aptos.dev/cli-tools/aptos-cli-tool/use-aptos-cli)

## Move Module by Aptos CLI

You can create account, query account, build & publish & running Move function with Aptos CLI, reference[Use Aptos CLI](https://aptos.dev/cli-tools/aptos-cli-tool/use-aptos-cli)

## Move Module by RPC

### Bytecode

After compile, there is a build output in package directory. You can get codes in package_dir/build/Examples/bytecode_modules.

### Move RPC

* [publish Move Module](./move_publish_test.go)
* [Read Move Module](./move_read_test.go)
* [Write Move Module](./move_write_test.go)

