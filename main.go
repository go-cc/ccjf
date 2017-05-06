// 中文简繁互转工具
package main

////////////////////////////////////////////////////////////////////////////
// Porgram: ccjf
// Purpose: Chinese-Character Jian<=>Fan converter
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits:
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v ccjfCLIGen.sh

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ccst "github.com/go-cc/cc-jianfan"
	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// convert types take a string and return a string value.
type convert func(string) string

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "ccjf"
	VERSION   = "0.1.0"
	buildTime = "2017-05-06"
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(jfDef),
		cli.Tree(fjDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

func ccjfC(ctx *cli.Context) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////
// jf

func jfCLI(ctx *cli.Context) error {
	return jfCvt(ctx, ccst.S2T)
}

////////////////////////////////////////////////////////////////////////////
// fj

func fjCLI(ctx *cli.Context) error {
	return jfCvt(ctx, ccst.T2S)
}

////////////////////////////////////////////////////////////////////////////
// jfCvt, Jian<=>Fan converter

func jfCvt(ctx *cli.Context, f convert) error {
	argv := ctx.RootArgv().(*rootT)
	//fmt.Printf("[%s]:\n  %+v\n  %v\n", ctx.Command().Name, argv, ctx.Args())

	// input data
	var dd string
	if ctx.IsSet("--in") { // -i,--in option is specified
		data, err := ioutil.ReadAll(argv.Filei)
		abortOn("Input", err)
		argv.Filei.Close()
		dd = string(data)
	} else {
		dd = strings.Join(ctx.Args(), " ")
		if dd == "" {
			ctx.WriteUsage()
			os.Exit(0)
		}
	}

	fmt.Printf(f(dd))
	return nil
}

//==========================================================================
// Support functions

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
