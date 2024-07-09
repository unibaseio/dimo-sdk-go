package sdk

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/MOSSV2/dimo-sdk-go/lib/types"
)

func UploadFileMeta(baseUrl string, auth types.Auth, fcws types.FileReceipt) error {

	form := url.Values{}
	fcwsb, err := json.Marshal(fcws)
	if err != nil {
		return err
	}
	form.Set("meta", hex.EncodeToString(fcwsb))

	_, err = doRequest(context.TODO(), baseUrl, "/api/uploadFileMeta", auth, strings.NewReader(form.Encode()))
	return err
}

func GetReplicaReceipt(baseUrl string, auth types.Auth, name string) (types.ReplicaCore, error) {
	var res types.ReplicaCore
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getReplicaReceipt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListReplica(baseUrl string, auth types.Auth, filter string) (types.ListReplicaResult, error) {
	res := types.ListReplicaResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listPiece", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("replica list: ", res)
	return res, nil
}

func GetPieceReceipt(baseUrl string, auth types.Auth, name string) (types.PieceReceipt, error) {
	var res types.PieceReceipt
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getPieceReceipt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListPiece(baseUrl string, auth types.Auth, filter string) (types.ListPieceResult, error) {
	res := types.ListPieceResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listPiece", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("piece list: ", res)
	return res, nil
}

func GetFileReceipt(baseUrl string, auth types.Auth, name string) (types.FileReceipt, error) {
	var res types.FileReceipt
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/getFileReceipt", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ListFile(baseUrl string, auth types.Auth, filter string) (types.ListFileResult, error) {
	res := types.ListFileResult{}
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

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/listFile", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	logger.Debug("file list: ", res)
	return res, nil
}

func RequestPiece(baseUrl string, auth types.Auth, name string) (types.PieceWitness, error) {
	form := url.Values{}
	form.Set("name", name)

	var res types.PieceWitness
	resByte, err := doRequest(context.TODO(), baseUrl, "/api/requestPiece", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func ConfirmPiece(baseUrl string, auth types.Auth, name, proof string) ([]byte, error) {
	form := url.Values{}
	form.Set("name", name)

	resByte, err := doRequest(context.TODO(), baseUrl, "/api/confirmPiece", auth, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	return resByte, nil
}
