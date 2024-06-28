package contract

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestGPU(t *testing.T) {
	sk, addr := makeAccount()
	err := transfer(addr, big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	err = GPUCheck(sk)
	if err != nil {
		t.Fatal(err)
	}

	err = RegisterNode(sk, 3, nil)
	if err != nil {
		t.Fatal(err)
	}

	ginfo, err := utils.GetGPUInfo("0")
	if err != nil {
		t.Fatal(err)
	}

	_gi, err := GetGPUIndex(ginfo.Name)
	if err != nil {
		err = AddGPU(sk, ginfo.Name)
		if err != nil {
			t.Fatal(err)
		}
	}

	_gi, err = GetGPUIndex(ginfo.Name)
	if err != nil {
		t.Fatal(err)
	}

	target, seed, err := GetGPUSetting()
	if err != nil {
		t.Fatal(err)
	}

	salt := make([]byte, 32)
	nt := time.Now()
	for i := 0; ; i++ {
		binary.BigEndian.PutUint64(salt[24:], uint64(i))
		val := crypto.Keccak256(seed[:], []byte(ginfo.Name), salt)
		if new(big.Int).SetBytes(val).Cmp(target) <= 0 {
			break
		}
		if i%100 == 0 {
			t.Log(i)
		}
	}
	t.Log(time.Since(nt))
	err = GPUMint(sk, _gi, salt)
	if err != nil {
		t.Fatal(err)
	}

	_ginfo, err := GetGPUInfo(_gi)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(_ginfo)
}

func TestModel(t *testing.T) {
	sk, addr := makeAccount()
	err := transfer(addr, big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	mc := types.ModelMeta{
		ModelCore: types.ModelCore{
			Name: "model-test",
		},
		Count: 100,
		Hash:  hex.EncodeToString([]byte("test")),
	}

	_mi, err := GetModelIndex(mc.Name)
	if err == nil {
		t.Fatal("already has")
	}

	err = AddModel(sk, mc)
	if err != nil {
		t.Fatal(err)
	}
	_mi, err = GetModelIndex(mc.Name)
	if err != nil {
		t.Fatal(err)
	}

	minfo, err := GetModelInfo(_mi)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(minfo)
}

func TestSpace(t *testing.T) {
	sk, addr := makeAccount()
	err := transfer(addr, big.NewInt(9e18))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("register compute node")
	err = RegisterNode(sk, 3, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("add gpu")
	_gn := "gpu-" + time.Now().Format(time.RFC3339)
	err = AddGPU(sk, _gn)
	if err != nil {
		t.Fatal(err)
	}

	minfo, err := GetModelInfo(0)
	if err != nil {
		t.Fatal(err)
	}

	mc := types.SpaceMeta{
		Name:  "space-" + time.Now().Format(time.RFC3339),
		Model: minfo.Name,
		GPU:   _gn,
	}

	_ai, err := GetSpaceIndex(mc.Name)
	if err == nil {
		t.Fatal("already has")
	}

	fmt.Println("add space")
	err = AddSpace(sk, mc)
	if err != nil {
		t.Fatal(err)
	}
	_ai, err = GetSpaceIndex(mc.Name)
	if err != nil {
		t.Fatal(err)
	}

	ainfo, err := GetSpaceInfo(_ai)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(ainfo)
}
