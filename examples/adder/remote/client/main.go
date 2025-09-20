package main

import (
	"context"
	"os"

	"github.com/nnlgsakib/go-wwfs-cmds/examples/adder"

	//cmds "github.com/nnlgsakib/go-wwfs-cmds"
	cmds "github.com/nnlgsakib/go-wwfs-cmds"
	cli "github.com/nnlgsakib/go-wwfs-cmds/cli"
	http "github.com/nnlgsakib/go-wwfs-cmds/http"
)

func main() {
	// parse the command path, arguments and options from the command line
	req, err := cli.Parse(context.TODO(), os.Args[1:], os.Stdin, adder.RootCmd)
	if err != nil {
		panic(err)
	}

	req.Options["encoding"] = cmds.Text

	// create http rpc client
	client := http.NewClient(":6798")

	// create an emitter
	re, err := cli.NewResponseEmitter(os.Stdout, os.Stderr, req)
	if err != nil {
		panic(err)
	}

	// send request to server
	err = client.Execute(req, re, nil)
	if err != nil {
		panic(err)
	}
}
