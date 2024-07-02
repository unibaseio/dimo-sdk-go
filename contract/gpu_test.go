package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
)

func TestGPU(t *testing.T) {
	sk, addr := makeAccount()
	err := transfer(addr, big.NewInt(1e14), big.NewInt(1e18))
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

	_ginfo, err := GetGPUInfo(_gi)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(_ginfo)
}

func TestModel(t *testing.T) {
	sk, addr := makeAccount()
	err := transfer(addr, big.NewInt(1e14), big.NewInt(1e18))
	if err != nil {
		t.Fatal(err)
	}

	mc := types.ModelMeta{
		Name: "model-test",
		Hash: hex.EncodeToString([]byte("test")),
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
	err := transfer(addr, big.NewInt(1e14), big.NewInt(9e18))
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

	mc := types.SpaceMeta{
		Name: "space-" + time.Now().Format(time.RFC3339),
		GPU:  _gn,
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
