package sdk

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/MOSSV2/dimo-sdk-go/contract"
	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
)

func TestModel(t *testing.T) {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		t.Fatal(err)
	}
	sks := string(skbyte[:64])
	sk, err := crypto.HexToECDSA(sks)
	if err != nil {
		t.Fatal(err)
	}

	base := ServerURL
	au, err := key.BuildAuth(sk, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	ml, err := ListModel(base, au, "latest")
	if err != nil {
		t.Fatal(err)
	}
	for _, mc := range ml.Models {
		_, err := contract.GetModelIndex(mc.Name)
		if err != nil {
			fmt.Println("add model: ", mc.Name)
			err = contract.AddModel(sk, mc.ModelMeta)
			if err != nil {
				t.Log(err)
			}
		}
	}
}

func TestGPU(t *testing.T) {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		t.Fatal(err)
	}
	sks := string(skbyte[:64])
	sk, err := crypto.HexToECDSA(sks)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; ; i++ {
		ginfo, err := contract.GetGPUInfo(uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(ginfo)
	}

	base := ServerURL
	au, err := key.BuildAuth(sk, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	ml, err := ListGPU(base, au, "latest")
	if err != nil {
		t.Fatal(err)
	}
	for _, mc := range ml.GPUs {
		_gi, err := contract.GetGPUIndex(mc.Name)
		if err != nil {
			continue
		}
		ginfo, err := contract.GetGPUInfo(_gi)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(mc)
		fmt.Println(ginfo)
	}
}

func TestSpace(t *testing.T) {
	skbyte, err := os.ReadFile("/tmp/sk")
	if err != nil {
		t.Fatal(err)
	}
	sks := string(skbyte[:64])
	sk, err := crypto.HexToECDSA(sks)
	if err != nil {
		t.Fatal(err)
	}

	base := ServerURL
	au, err := key.BuildAuth(sk, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	ml, err := ListSpace(base, au, "latest")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("=========")
	for _, mc := range ml.Spaces {

		_gi, err := contract.GetSpaceIndex(mc.Name)
		if err != nil {
			continue
		}
		ginfo, err := contract.GetSpaceInfo(_gi)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(mc)
		fmt.Println(ginfo)
	}
}

func TestList(t *testing.T) {
	dir, err := homedir.Expand("~/relay")
	if err != nil {
		t.Fatal(err)
	}

	mrm, err := ListLocalDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mrm.Name)
	for k, v := range mrm.Files {
		if len(v) == 64 {
			_, err := hex.DecodeString(v)
			if err == nil {
				t.Log(k, v)
			}
		}
	}

	mrm.Name = "relayd1"

	rootdir, err := homedir.Expand("~/test/test1")
	if err != nil {
		t.Fatal(err)
	}
	err = DownloadModel("", rootdir, types.Auth{}, mrm, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal()
}

func TestDecode(t *testing.T) {
	pa := "~/dimo-go/bin/stream-edge"
	pb := "~/as"

	pa, _ = homedir.Expand(pa)
	pb, _ = homedir.Expand(pb)

	fa, err := os.Open(pa)
	if err != nil {
		t.Fatal(err)
	}
	fb, err := os.Open(pb)
	if err != nil {
		t.Fatal(err)
	}

	bufa := make([]byte, 8)
	bufb := make([]byte, 8)
	fa.ReadAt(bufa, 9901120*0)
	fb.ReadAt(bufb, 9901120*0)

	if !bytes.Equal(bufa, bufb) {
		t.Fatal("not equal", hex.EncodeToString(bufa), hex.EncodeToString(bufb))
	}
}
