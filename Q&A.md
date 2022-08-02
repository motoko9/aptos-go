# Q&A

1. VM STATUS is "Transaction Executed and Committed with Error LINKER_ERROR"

The parameter of type_arguments in transaction payload is wrong. Please check address, module and struct is right and on chain.

2. managed_coin::initialize error "Move abort: code 65536 at 0000000000000000000000000000000000000000000000000000000000000001::coin"

Coin only can be initialized by the account published this coin module. This account is the owner of this coin and can mint coin.

