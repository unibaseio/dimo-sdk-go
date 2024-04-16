package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/MOSSV2/dimo-sdk-go/lib/key"
	"github.com/MOSSV2/dimo-sdk-go/lib/types"
	"github.com/MOSSV2/dimo-sdk-go/sdk"
)

func main() {
	modelstr := flag.String("name", "CompVis/stable-diffusion-v1-1", "model repo name in huggingface")
	revstr := flag.String("revision", "", "model revision")
	flag.Parse()

	au := key.BuildAuthLocal([]byte("hub"))

	er, err := sdk.ListEdge(sdk.ServerURL, au, types.HubType)
	if err != nil {
		return
	}
	if len(er.Edges) == 0 {
		fmt.Println("no hub node")
		return
	}

	url := er.Edges[rand.Intn(len(er.Edges))].ExposeURL
	fmt.Println("remote hub: ", url)

	rc := types.RepoCore{
		Name:     *modelstr,
		Revision: *revstr,
	}

	err = sdk.SubmitModelRepo(url, au, rc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("submit model repo: ", rc)
}
