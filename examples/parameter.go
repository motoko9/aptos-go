package examples

import (
    "github.com/motoko9/aptos-go/crypto"
)

var PrivateKey crypto.PrivateKey
var AlicePrivateKey crypto.PrivateKey

// address: 0xe50af979b3bbbbaa5e7ae9b0ddedad6cb345dcca45a3aa303312a9dab00147f7
var BobPrivateKey crypto.PrivateKey

func init() {
    var err error
    privHexStr := "b2947e614dbce167ca666ad29c2e84a743e1bd69f53297ee571ed4e6cac683fe393fd356f184e5d629929680c46dcc6b851f4d412fbcfc94138ad67cf08bfc55"
    PrivateKey, err = crypto.NewPrivateKeyFromHexString(privHexStr)
    if err != nil {
        panic("create private key failed")
    }

    AlicePrivHexStr := "a88ae4f8870e180ae8fca0472c1a2dca986b5d609390b71099cfdd727c1edb956da6045ebb54e313d7f37ecd9a6b067912fcc0082405e285222fb9ebf19b2364"
    AlicePrivateKey, err = crypto.NewPrivateKeyFromHexString(AlicePrivHexStr)
    if err != nil {
        panic("create alice private key failed")
    }

    BobPrivHexStr := "fc20bed4ec67f04b28f66faafc3e178c6c8936112c0e5f0a9c005fc056cf20fb729c5ad55087d8c9d2280c7d26e888a1ab4b463c56eb3901b5f9b150317cc3ae"
    BobPrivateKey, err = crypto.NewPrivateKeyFromHexString(BobPrivHexStr)
    if err != nil {
        panic("create alice private key failed")
    }
}