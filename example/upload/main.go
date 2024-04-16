package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/contract"
	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mitchellh/go-homedir"
)

func main() {
	skstr := flag.String("sk", "", "private key for sending transaction")
	pathstr := flag.String("path", "", "dir or file path to upload")
	mf := flag.Bool("model", false, "download model or regular file/dir")
	flag.Parse()

	sk, err := crypto.HexToECDSA(*skstr)
	if err != nil {
		sk, err = crypto.GenerateKey()
		if err != nil {
			return
		}
		skbyte := crypto.FromECDSA(sk)
		fmt.Printf("=== generate privatekey: %s ===\n", hex.EncodeToString(skbyte))
	}

	fp, err := homedir.Expand(*pathstr)
	if err != nil {
		return
	}

	if *mf {
		err = UploadModel(sk, fp)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err = UploadFile(sk, fp)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func UploadFile(sk *ecdsa.PrivateKey, fp string) error {
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		return err
	}

	// charge from server
	sdk.Login(sdk.ServerURL, au)

	ar, err := sdk.BalanceOf(sdk.ServerURL, au)
	if err != nil {
		return err
	}
	// wait for has balance
	for ar.Value.Cmp(big.NewInt(0)) == 0 {
		time.Sleep(3 * time.Second)
		ar, err = sdk.BalanceOf(sdk.ServerURL, au)
		if err != nil {
			return err
		}
	}

	fi, err := os.Stat(fp)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		// upload to stream and submit to gateway
		res, submitter, err := sdk.UploadToStream(sdk.ServerURL, au, fp, "")
		if err != nil {
			return err
		}
		fmt.Printf("upload %s, sha256: %s\n", fp, res.Name)
		fmt.Printf("submit %s to chain\n", res.Name)

		// submit meta to chain
		err = contract.AddFileAndPiece(sk, res.Name, res.FileCore, submitter)
		if err != nil {
			return err
		}
		return nil
	}

	// recursive upload files in directory
	return filepath.Walk(fp, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		res, submitter, err := sdk.UploadToStream(sdk.ServerURL, au, fileName, "")
		if err != nil {
			return nil
		}
		fmt.Printf("upload %s, sha256: %s\n", fp, res.Name)
		fmt.Printf("submit %s to chain\n", res.Name)
		err = contract.AddFileAndPiece(sk, res.Name, res.FileCore, submitter)
		if err != nil {
			return err
		}
		return nil
	})
}

func UploadModel(sk *ecdsa.PrivateKey, fp string) error {
	au, err := key.BuildAuth(sk, []byte("upload"))
	if err != nil {
		return err
	}

	sdk.Login(sdk.ServerURL, au)

	ar, err := sdk.BalanceOf(sdk.ServerURL, au)
	if err != nil {
		return err
	}
	for ar.Value.Cmp(big.NewInt(0)) == 0 {
		time.Sleep(3 * time.Second)
		ar, err = sdk.BalanceOf(sdk.ServerURL, au)
		if err != nil {
			return err
		}
	}

	fi, err := os.Stat(fp)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("model path should be dir")
	}

	_, err = sdk.GetModel(sdk.ServerURL, au, path.Base(fp))
	if err == nil {
		return fmt.Errorf("already has model %s", path.Base(fp))
	}

	mrm, err := sdk.UploadModelFiles(sdk.ServerURL, sk, au, fp)
	if err != nil {
		return err
	}

	err = sdk.SubmitModel(sdk.ServerURL, au, mrm)
	if err != nil {
		return err
	}

	fmt.Printf("upload and submit model %s\n", mrm.Name)

	return nil
}
