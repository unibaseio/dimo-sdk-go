package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MOSSV2/dimo-sdk-go/lib/log"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/sha3"
)

var logger = log.Logger("sdk")

func init() {
	log.SetLogLevel("DEBUG")
}

const (
	ServerURL = "http://52.220.254.5:8080"
)

func DecodeAuth(authstr string) (types.Auth, error) {
	au := types.Auth{}
	if authstr == "" {
		return au, fmt.Errorf("nil authorization")
	}

	ab, err := hex.DecodeString(authstr)
	if err == nil {
		err = json.Unmarshal(ab, &au)
		if err != nil {
			return au, err
		}
	} else {
		type JSAuth struct {
			Type string
			Addr common.Address
			Time int64
			Hash string
			Sign string
		}
		jau := JSAuth{}
		err := json.Unmarshal([]byte(authstr), &jau)
		if err != nil {
			return au, err
		}

		au.Type = jau.Type
		au.Addr = jau.Addr
		au.Time = jau.Time
		if strings.HasPrefix(jau.Hash, "0x") {
			au.Hash, err = hex.DecodeString(jau.Hash[2:])
			if err != nil {
				return au, err
			}
		}

		if strings.HasPrefix(jau.Sign, "0x") {
			au.Sign, err = hex.DecodeString(jau.Sign[2:])
			if err != nil {
				return au, err
			}
		}
	}
	return au, nil
}

func VerifyAuth(au types.Auth) error {
	if len(au.Sign) == 65 {
		recoveryID := int(au.Sign[64])
		if recoveryID >= 27 && recoveryID <= 34 {
			recoveryID -= 27
		} else if recoveryID >= 35 && recoveryID <= 38 {
			recoveryID -= 35
		}
		au.Sign[64] = byte(recoveryID)
	}

	b := make([]byte, len(au.Hash)+8)
	copy(b, au.Hash)
	binary.BigEndian.PutUint64(b[len(au.Hash):], uint64(au.Time))
	var sum []byte
	switch au.Type {
	case "personal_sign":
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte{0x19})
		hash.Write([]byte("Ethereum Signed Message:"))
		hash.Write([]byte{0x0A})
		hash.Write([]byte(strconv.Itoa(len(b))))
		hash.Write(b)
		sum = hash.Sum(nil)
	default:
		sums := sha256.Sum256(b)
		sum = sums[:]
	}

	rePub, err := crypto.Ecrecover(sum, au.Sign)
	if err != nil {
		return err
	}

	if !bytes.Equal(au.Addr.Bytes(), utils.ToEthAddress(rePub)) {
		return fmt.Errorf("invalid auth %s", au.Addr)
	}

	return nil
}

// hash is random byte now
func BuildAuth(addr, privk string, hash []byte) types.Auth {
	h := sha256.New()
	ts := time.Now().Unix()
	h.Write(hash)
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(ts))
	h.Write(b)

	sum := h.Sum(nil)

	pb, _ := crypto.HexToECDSA(privk)

	sign, _ := crypto.Sign(sum[:], pb)

	au := types.Auth{
		Addr: utils.HexToAddress(addr),
		Time: ts,
		Hash: hash,
		Sign: sign,
	}
	return au
}

func doRequest(ctx context.Context, baseUrl, method string, au types.Auth, r io.Reader) ([]byte, error) {
	haddr := baseUrl + method
	hreq, err := http.NewRequestWithContext(ctx, "POST", haddr, r)
	if err != nil {
		return nil, err
	}

	aub, err := json.Marshal(au)
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Authorization", hex.EncodeToString(aub))
	hreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	bar := progressbar.DefaultBytes(-1, baseUrl+" download:")

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := progressbar.NewReader(resp.Body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return res, nil
}

func Get(ctx context.Context, baseUrl string, r io.Reader) ([]byte, error) {
	haddr := baseUrl
	hreq, err := http.NewRequestWithContext(ctx, "GET", haddr, r)
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

	bar := progressbar.DefaultBytes(-1, baseUrl+" download:")

	resp, err := defaultHTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pr := progressbar.NewReader(resp.Body, bar)
	res, err := io.ReadAll(&pr)
	if err != nil {
		return nil, err
	}
	bar.Finish()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response: %s, msg: %s", resp.Status, res)
	}

	return res, nil
}
