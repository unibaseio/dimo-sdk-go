package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/bls"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestTransfer(t *testing.T) {
	sk, addr := makeAccount()
	val := big.NewInt(1e18)
	val.Mul(val, big.NewInt(100000))
	addr = common.HexToAddress("0x3424aD62dD417b2edFfacBa665316dD8cFA27fF7")
	err := transfer(addr, val)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(sk)
}

func TestReward(t *testing.T) {
	sk, addr := makeAccount()
	val := big.NewInt(1e18)
	err := transfer(addr, val)
	if err != nil {
		t.Fatal(err)
	}
	ri, err := NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	au, err := makeAuth(big.NewInt(int64(DevChainID)), sk)
	if err != nil {
		t.Fatal(err)
	}
	for i := uint64(64); i < 3000; i += 64 {
		tx, err := ri.Mint(au, i)
		if err != nil {
			t.Fatal(i, err)
		}

		err = CheckTx(DevChain, tx.Hash())
		if err != nil {
			t.Fatal(i, err)
		}
	}
}

func TestTotalReward(t *testing.T) {
	ri, err := NewReward(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	ts, err := ri.TotalShare(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ts)
	cur, err := ri.Current(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cur)
	val, err := ri.Rest(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(val)
	rsi, err := ri.GetEReward(&bind.CallOpts{From: Base}, cur-1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsi)
	t.Fatal()
}

func TestNodeCheck(t *testing.T) {
	addr := common.HexToAddress("0x5f375a47491a77f8e5bd2d9f179894188142c315")
	err := CheckNode(addr, 1)
	if err != nil {
		t.Fatal(err)
	}
	psk, paddr := makeAccount()
	err = transfer(paddr, big.NewInt(2e18))
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(psk, 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckNode(paddr, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal("")
}

func TestCheck(t *testing.T) {
	ctx := context.TODO()
	bi, err := NewBank(ctx)
	if err != nil {
		t.Fatal(err)
	}
	taddr, err := bi.Get(&bind.CallOpts{From: Base}, "token")
	if err != nil {
		t.Fatal(err)
	}
	if taddr != TokenAddr {
		t.Fatal("wrong token")
	}

	si, err := NewSpace(ctx)
	if err != nil {
		t.Fatal(err)
	}
	baddr, err := si.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if baddr != BankAddr {
		t.Fatal("wrong bank in space")
	}

	gi, err := NewGPU(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bgaddr, err := gi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bgaddr != BankAddr {
		t.Fatal("wrong bank in gpu")
	}

	mi, err := NewModel(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bmaddr, err := mi.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bmaddr != BankAddr {
		t.Fatal("wrong bank in model")
	}

	ci, err := NewControl(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bcaddr, err := ci.Bank(&bind.CallOpts{From: Base})
	if err != nil {
		t.Fatal(err)
	}
	if bcaddr != BankAddr {
		t.Fatal("wrong bank in control")
	}
}

func transfer(to common.Address, val *big.Int) error {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		return err
	}
	sk := string(skbyte[:64])
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return err
	}
	err = Transfer(DevChain, privateKey, to, big.NewInt(1e15))
	if err != nil {
		return err
	}
	return TransferToken(DevChain, privateKey, TokenAddr, to, val)
}

func makeAccount() (*ecdsa.PrivateKey, common.Address) {
	privk, err := crypto.GenerateKey()
	if err != nil {
		panic("generate")
	}
	publicKey := privk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("public")
	}

	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("generate: ", addr)
	//skbyte := crypto.FromECDSA(privk)
	//common.Bytes2Hex(skbyte)

	return privk, addr
}

func TestBlock(t *testing.T) {
	client, err := ethclient.Dial(DevChain)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(latest)

	for i := latest; i > latest-5; i-- {
		bk, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(bk)
	}
}

func TestChoose(t *testing.T) {
	_, addr := makeAccount()

	seed := [32]byte{100}
	count := 10
	pcnt := 0
	for i := 0; i < 20; i++ {
		ci, err := choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		nci := Choose(addr, seed, uint64(count), uint64(pcnt), uint64(i))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}

func TestOrder(t *testing.T) {
	for i := 0; i < 100; i++ {
		count := rand.Uint64() / 2
		pcnt := rand.Int63n(int64(count))
		ci, err := getOrder(uint64(count), uint64(pcnt))
		if err != nil {
			t.Fatal(err)
		}
		nci, _ := GetOrder(uint64(count), uint64(pcnt))
		if ci != nci {
			t.Fatal("unequal at:", i, ci, nci)
		}
	}
}

func TestMarshal(t *testing.T) {
	var g bls.G1
	g.ScalarMultiplicationBase(big.NewInt(10))

	gsbyte := G1InSolidity(g)
	gr, err := SolidityToG1(gsbyte)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g.String())
	fmt.Println(gr.String())
	if !g.Equal(&gr) {
		t.Fatal("unequal")
	}
}

func TestReceipt(t *testing.T) {
	tx := common.HexToHash("0xa781ef9b7d27821e727e5ae6fb2a2a1bfac9668df78dbec689c3621ca794643c")
	receipt, err := GetTransactionReceipt(DevChain, tx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("tx receipt: ", receipt)

	fmt.Println("tx status: ", receipt.Status)
	if receipt.Status == 1 {
		fmt.Println("tx success")
		return
	}

	for _, log := range receipt.Logs {
		fmt.Println(log)
	}

	err = AnalyzeTransactionFailure(tx)
	fmt.Println("tx fail: ", err)
}
