# Q&A

1. VM STATUS is "Transaction Executed and Committed with Error LINKER_ERROR"

The parameter of type_arguments in transaction payload is wrong. Please check address, module and struct is right and on chain.

2. managed_coin::initialize error "Move abort: code 65536 at 0000000000000000000000000000000000000000000000000000000000000001::coin"

Coin only can be initialized by the account published this coin module. This account is the owner of this coin and can mint coin.

3. publish move module to chain, but transaction can not be included into ledger.


4. VM STATUS is "Transaction Executed and Committed with Error LOOKUP_FAILED", when publish move module

5. encode_submission response, {"message":"The given transaction is invalid","error_code":"invalid_input","vm_error_code":null}

Transaction payload is not right, should check:
Function is right, address, module, struct.
TypeArguments is right.
Arguments should be right, include the type of argument.

6. publish move module with error
   "Error": "Unexpected error: Unable to resolve packages for package 'usdt'"
   
Move module should add dependencies AptosFramework

```
[dependencies]
AptosFramework = { git = "https://github.com/aptos-labs/aptos-core.git", subdir = "aptos-move/framework/aptos-framework", rev = "devnet" }
```

7. publish move module with error
```
   tangaoyuan@ip-172-26-1-15 contracts % aptos move publish --profile arbitrage
   Compiling, may take a little while to download git dependencies...
   UPDATING GIT DEPENDENCY https://github.com/blockchain-develop/v1-core.git
   UPDATING GIT DEPENDENCY https://github.com/aptos-labs/aptos-core.git
   UPDATING GIT DEPENDENCY https://github.com/pontem-network/liquidswap.git
   UPDATING GIT DEPENDENCY https://github.com/aptos-labs/aptos-core.git
   UPDATING GIT DEPENDENCY https://github.com/aptos-labs/aptos-core.git
   {
   "Error": "Unexpected error: Unable to resolve packages for package 'arbitragev1'"
   }
```
module's dependency are conflict


