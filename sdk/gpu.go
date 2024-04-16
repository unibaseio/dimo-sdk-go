package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func SubmitGPU(baseUrl string, auth types.Auth, mr types.GPUMeta) error {
	mmb, err := json.Marshal(mr)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitGPU", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	re := new(types.GPUMeta)
	err = json.Unmarshal(res, re)
	if err != nil {
		return err
	}

	return nil
}

func GetGPU(baseUrl string, auth types.Auth, GPUName string) (types.GPUMeta, error) {
	form := url.Values{}
	form.Set("name", GPUName)

	re := types.GPUMeta{}
	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getGPU", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return re, err
	}

	err = json.Unmarshal(resByte, &re)
	if err != nil {
		return re, err
	}

	return re, nil
}

func ListGPU(baseUrl string, auth types.Auth, filter string) (types.ListGPUResult, error) {
	res := types.ListGPUResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listGPU", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("GPU list: ", res)
	return res, nil
}
