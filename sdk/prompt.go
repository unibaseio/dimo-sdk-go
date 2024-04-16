package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func SubmitPrompt(baseUrl string, auth types.Auth, prompt string) error {
	mm := types.PromptMeta{
		//Prompt: prompt,
	}

	mmb, err := json.Marshal(mm)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("txMsg", hex.EncodeToString(mmb))

	res, err := doRequest(context.TODO(), baseUrl, "/api/submitPrompt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	re := new(types.PromptMeta)
	err = json.Unmarshal(res, re)
	if err != nil {
		return err
	}

	return nil
}

func GetPrompt(baseUrl string, auth types.Auth, prompt uint64) (types.PromptMeta, error) {
	res := types.PromptMeta{}
	form := url.Values{}
	form.Set("prompt", strconv.Itoa(int(prompt)))

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getPrompt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// own prompt
func ListPrompt(baseUrl string, auth types.Auth, filter string) (types.ListPromptResult, error) {
	res := types.ListPromptResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listPrompt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("prompt list my: ", res)
	return res, nil
}
