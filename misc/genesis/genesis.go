package genesis

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
)

// DefaultGenesisBlock returns the Ethereum main net genesis block.
func SimulateGenesisBlock() *core.Genesis {
	return &core.Genesis{
		Config: &params.ChainConfig{
			ChainID:        big.NewInt(15),
			HomesteadBlock: big.NewInt(0),
			EIP155Block:    big.NewInt(0),
			EIP158Block:    big.NewInt(0),
			ByzantiumBlock: big.NewInt(0),
		},
		Nonce:      uint64(0xdeadbeefdeadbeef),
		ExtraData:  hexutil.MustDecode("0x"),
		GasLimit:   uint64(0x1e8480000),
		Difficulty: big.NewInt(0x40),
		Alloc:      decodePrealloc(simulateAllocData),
	}
}

// DefaultGenesisBlock returns the Ethereum main net genesis block.
func DefaultGenesisBlock() *core.Genesis {
	config := params.MainnetChainConfig
	config.HomesteadBlock = big.NewInt(0)
	config.EIP150Block = big.NewInt(0)
	config.EIP155Block = big.NewInt(0)
	config.EIP158Block = big.NewInt(0)
	config.DAOForkBlock = big.NewInt(0)
	config.DAOForkSupport = false
	config.ByzantiumBlock = big.NewInt(0)
	config.ConstantinopleBlock = big.NewInt(0)
	// this will be overridden
	config.ChainID = big.NewInt(0)
	genesis := &core.Genesis{
		Config: config,
		Nonce:  66,
		//CyberMiles for E-commerce
		ExtraData:  hexutil.MustDecode("0x43796265724d696c657320666f7220452d636f6d6d65726365"),
		GasLimit:   uint64(0x1e8480000),
		Difficulty: big.NewInt(17179869184),
		Alloc:      decodePrealloc(mainnetAllocData),
	}

	return genesis
}

// DevGenesisBlock returns the 'geth --dev' genesis block.
func DevGenesisBlock() *core.Genesis {
	return &core.Genesis{
		Config: &params.ChainConfig{
			ChainID:        big.NewInt(15),
			HomesteadBlock: big.NewInt(0),
			EIP155Block:    big.NewInt(0),
			EIP158Block:    big.NewInt(0),
			ByzantiumBlock: big.NewInt(0),
		},
		Nonce:      uint64(0xdeadbeefdeadbeef),
		ExtraData:  hexutil.MustDecode("0x"),
		GasLimit:   uint64(0x1e8480000),
		Difficulty: big.NewInt(0x40),
		Alloc:      decodePrealloc(devAllocData),
	}
}

func decodePrealloc(data string) core.GenesisAlloc {
	var p []struct{ Addr, Balance *big.Int }
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	ga := make(core.GenesisAlloc, len(p))
	for _, account := range p {
		ga[common.BigToAddress(account.Addr)] = core.GenesisAccount{Balance: account.Balance}
	}
	return ga
}
