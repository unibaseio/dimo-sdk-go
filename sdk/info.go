package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func Info(baseUrl string) (types.EdgeReceipt, error) {
	res := types.EdgeReceipt{}
	resp, err := http.Get(baseUrl + "/api/info")
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s", resp.Status)
	}

	resb, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resb, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func Login(baseUrl string, auth types.Auth) error {
	_, err := doRequest(context.TODO(), baseUrl, "/api/login", auth, nil)
	if err != nil {
		return err
	}

	return nil
}

func BalanceOf(baseUrl string, auth types.Auth) (types.AccountResult, error) {
	re := types.AccountResult{}

	res, err := doRequest(context.TODO(), baseUrl, "/api/balanceOf", auth, nil)
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(res, &re)
	if err != nil {
		return re, err
	}

	logger.Debugf("%s has balance: %d", auth.Addr, re.Value)
	return re, nil
}

func GetChal(baseUrl string) (types.ChalResult, error) {
	var res types.ChalResult
	resp, err := http.Get(baseUrl + "/api/chal")
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("response: %s, msg: %s", resp.Status, resByte)
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
