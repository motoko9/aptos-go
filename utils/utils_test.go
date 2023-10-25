package utils

import (
	"fmt"
	"testing"
)

func TestUtils_ExtractFromResource(t *testing.T) {
	{
		r1 := "0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::liquidity_pool::LiquidityPool<0xe3c0ae56b7de6b4f7071132f63b1937dc124028637d185c6143c4e7f4d4ad24c::metatoken::META, 0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDC, 0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable>"
		t1, t2, err := ExtractFromResource(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", t1, t2)
	}

	{
		r1 := "0x1::coin::coininfo<0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDT>"
		t1, t2, err := ExtractFromResource(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", t1, t2)
	}

	{
		r1 := "0x1::coin::CoinInfo<0x5a97986a9d031c4567e15b797be516910cfcb4156312482efc6a19c0a30c948::lp_coin::LP<0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDC, 0xd0b4efb4be7c3508d9a26a9b5405cf9f860d0b9e5fe2f498b90e68b8d2cedd3e::aptos_launch_token::AptosLaunchToken, 0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable>>"
		t1, t2, err := ExtractFromResource(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", t1, t2)
	}

	{
		r1 := "0x5a97986a9d031c4567e15b797be516910cfcb4156312482efc6a19c0a30c948::lp_coin::LP<0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDC, 0xd0b4efb4be7c3508d9a26a9b5405cf9f860d0b9e5fe2f498b90e68b8d2cedd3e::aptos_launch_token::AptosLaunchToken, 0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable>"
		t1, t2, err := ExtractFromResource(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", t1, t2)
	}

	{
		r1 := "0x1::account::Account"
		t1, t2, err := ExtractFromResource(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v\n", t1, t2)
	}
}

func TestUtils_ExtractAddressFromType(t *testing.T) {
	{
		r1 := "0x5a97986a9d031c4567e15b797be516910cfcb4156312482efc6a19c0a30c948::lp_coin::LP<0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDC, 0xd0b4efb4be7c3508d9a26a9b5405cf9f860d0b9e5fe2f498b90e68b8d2cedd3e::aptos_launch_token::AptosLaunchToken, 0x190d44266241744264b964a37b8f09863167a12d3e70cda39376cfb4e3561e12::curves::Stable>"
		t1, err := ExtractAddressFromType(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", t1)
	}
}

func TestUtils_ExtractFromFunction(t *testing.T) {
	{
		r1 := "0xf22bede237a07e121b56d91a491eb7bcdfd1f5907926a9e58338f964a01b17fa::asset::USDT"
		t1, t2, t3, err := ExtractFromFunction(r1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s, %s, %s\n", t1, t2, t3)
	}
}
