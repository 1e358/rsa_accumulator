package main

import (
	"fmt"
	"runtime"

	"github.com/1e358/rsa_accumulator/experiments"
	"github.com/1e358/rsa_accumulator/merkleswap"
	"github.com/1e358/rsa_accumulator/zkmultiswap"
)

func main() {

	//zkmultiswap.TestMultiSwap(10)
	// updateRates denotes the percentage of updates in the total users, i.e. number of updates = users/updateRates
	updateRates := 64

	twoTo14 := 16384
	twoTo15 := 32768
	twoTo16 := 65536
	twoTo17 := 131072
	twoTo18 := 262144
	twoTo19 := 524288

	// outputs a Solidity smart
	zkmultiswap.TestMultiSwapAndOutputSmartContract(10)
	err := zkmultiswap.TestMultiSwapAndOutputSmartContract2(10)
	if err != nil {
		panic(err)
	}

	// // test Membership proof Verification and proof size
	experiments.TestMembershipVerify()

	// the following code is to test the time percentage for each part of the system in a single core.
	fmt.Println("TestNotusMultiSwap for ", twoTo14, " users and ", uint32(twoTo14/updateRates), " updates.")
	experiments.TestNotusMultiSwap(uint32(twoTo14), uint32(twoTo14/updateRates))
	fmt.Println("TestNotusMultiSwap for ", twoTo15, " users and ", uint32(twoTo15/updateRates), " updates.")
	experiments.TestNotusMultiSwap(uint32(twoTo15), uint32(twoTo15/updateRates))
	fmt.Println("TestNotusMultiSwap for ", twoTo16, " users and ", uint32(twoTo16/updateRates), " updates.")
	experiments.TestNotusMultiSwap(uint32(twoTo16), uint32(twoTo16/updateRates))

	fmt.Println("Test Membership precomputation under different group size")
	experiments.TestDifferentMembershipForDISingleThread()

	fmt.Println("Test Notus With Different Size")
	experiments.TestMultiSwapWithDifferentSize()
	runtime.GC()
	fmt.Println("TestNotusMultiSwap for ", twoTo15, " users and ", uint32(twoTo15/updateRates), " updates in parallel")
	fmt.Println("Assuming 32 group partations for 2^20 user, each group of 2^15 users uses 32 cores, total 1024 cores")
	experiments.TestNotusParallel(uint32(twoTo15), uint32(twoTo15/updateRates))
	runtime.GC()
	fmt.Println("TestNotusMultiSwap for ", twoTo16, " users and ", uint32(twoTo16/updateRates), " updates in parallel")
	fmt.Println("Assuming 16 group partations for 2^20 user, each group of 2^16 users uses 32 cores, total 512 cores")
	experiments.TestNotusParallel(uint32(twoTo16), uint32(twoTo16/updateRates))
	runtime.GC()
	fmt.Println("TestNotusMultiSwap for ", twoTo17, " users and ", uint32(twoTo17/updateRates), " updates in parallel")
	fmt.Println("Assuming 8 group partations for 2^20 user, each group of 2^17 users uses 32 cores, total 256 cores")
	experiments.TestNotusParallel(uint32(twoTo17), uint32(twoTo17/updateRates))
	runtime.GC()
	fmt.Println("TestNotusMultiSwap for ", twoTo18, " users and ", uint32(twoTo18/updateRates), " updates in parallel")
	fmt.Println("Assuming 4 group partations for 2^20 user, each group of 2^18 users uses 32 cores, total 128 cores")
	experiments.TestNotusParallel(uint32(twoTo18), uint32(twoTo18/updateRates))
	runtime.GC()
	fmt.Println("TestNotusMultiSwap for ", twoTo19, " users and ", uint32(twoTo19/updateRates), " updates in parallel")
	fmt.Println("Assuming 2 group partations for 2^20 user, each group of 2^19 users uses 32 cores, total 64 cores")
	experiments.TestNotusParallel(uint32(twoTo19), uint32(twoTo19/updateRates))
	runtime.GC()

	// simulate the cost of a Merkle Swap with depth 28 for the following number of users
	merkleswap.TestMerkleMultiSwap(1024)
}
