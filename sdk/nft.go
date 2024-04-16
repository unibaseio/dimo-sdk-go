package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/mitchellh/go-homedir"
)

func UploadNFT(baseUrl string, auth types.Auth, prompt uint64, spaceID, modelName, filePath string, pf []byte) error {
	p, err := homedir.Expand(filePath)
	if err != nil {
		return err
	}

	pbyte, err := os.ReadFile(p)
	if err != nil {
		return err
	}

	mm := types.NFTMeta{
		Name:   utils.DataToName(pbyte),
		Size:   uint64(len(pbyte)),
		Type:   mime.TypeByExtension(filepath.Ext(p)),
		Price:  big.NewInt(150),
		Prompt: prompt,
		Model:  modelName,
		Space:  spaceID,
		Owner:  auth.Addr,
	}

	mmb, err := json.Marshal(mm)
	if err != nil {
		return err
	}

	ipr, ipw := io.Pipe()
	mwriter := multipart.NewWriter(ipw)
	go func() {
		//mwriter.WriteField("address", addr)
		mwriter.WriteField("txMsg", hex.EncodeToString(mmb))
		mwriter.WriteField("proof", hex.EncodeToString(pf))

		defer ipw.Close()
		defer mwriter.Close()

		part, err := mwriter.CreateFormFile("file", p)
		if err != nil {
			return
		}
		pf, err := os.Open(p)
		if err != nil {
			return
		}
		defer pf.Close()

		io.Copy(part, pf)
	}()

	haddr := baseUrl + "/api/uploadNFT"
	hreq, err := http.NewRequest("POST", haddr, ipr)
	if err != nil {
		return err
	}

	aub, err := json.Marshal(auth)
	if err != nil {
		return err
	}

	hreq.Header.Add("Authorization", hex.EncodeToString(aub))
	hreq.Header.Add("Content-Type", mwriter.FormDataContentType())
	defaultHTTPClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			WriteBufferSize:       16 << 10, // 16KiB moving up from 4KiB default
			ReadBufferSize:        16 << 10, // 16KiB moving up from 4KiB default
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    true,
		},
	}

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return nil
}

func BuyNFT(baseUrl string, auth types.Auth, nftName string) error {
	form := url.Values{}
	form.Set("name", nftName)

	_, err := doRequest(context.TODO(), baseUrl, "/api/buyNFT", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	return nil
}

func ListNFT(baseUrl string, auth types.Auth, filter string) (types.ListNFTResult, error) {
	res := types.ListNFTResult{}
	opt := types.Options{
		UserDefined: make(map[string]string),
	}
	opt.UserDefined["filter"] = filter

	optyByte, err := json.Marshal(opt)
	if err != nil {
		return res, err
	}

	form := url.Values{}
	form.Set("option", hex.EncodeToString(optyByte))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listNFT", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug(auth.Addr, " has: ", res)
	return res, nil
}

func GetNFT(baseUrl string, auth types.Auth, nftName string) ([]byte, error) {
	form := url.Values{}
	form.Set("name", nftName)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getNFT", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	logger.Debug("nft hash:", utils.DataToName(resByte))

	return resByte, nil
}
