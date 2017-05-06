////////////////////////////////////////////////////////////////////////////
// Program: ccjfC
// Purpose: Chinese-Character Jian<=>Fan converter
// Authors: Tong Sun (c) 2016-2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// ccjfC

type rootT struct {
	cli.Helper
	Filei   *clix.Reader `cli:"i,in" usage:"The Chinese text file to read from (or stdin)"`
	Verbose int          `cli:"v,verbose" usage:"Be verbose (with level 1 & 2)"`
}

var root = &cli.Command{
	Name:   "ccjfC",
	Desc:   "Chinese-Character Jian<=>Fan converter\nbuilt on " + buildTime,
	Text:   "Command line tool to convert between Chinese jian/fan characters",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     ccjfC,

	NumArg: cli.AtLeast(1),
}

// func main() {
// 	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
// 	//NOTE: You can set any writer implements io.Writer
// 	// default writer is os.Stdout
// 	if err := cli.Root(root,
// 		cli.Tree(jfDef),
// 		cli.Tree(fjDef)).Run(os.Args[1:]); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	fmt.Println("")
// }

// func ccjfC(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }

////////////////////////////////////////////////////////////////////////////
// jf

// func jfCLI(ctx *cli.Context) error {
// 	rootArgv = ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*jfT)
// 	fmt.Printf("[jf]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }

type jfT struct {
}

var jfDef = &cli.Command{
	Name: "jf",
	Desc: "jian => fan",
	Text: "convert from jian to fan",
	Argv: func() interface{} { return new(jfT) },
	Fn:   jfCLI,

	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// fj

// func fjCLI(ctx *cli.Context) error {
// 	rootArgv = ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*fjT)
// 	fmt.Printf("[fj]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }

type fjT struct {
}

var fjDef = &cli.Command{
	Name: "fj",
	Desc: "fan => jian",
	Text: "convert from fan to jian",
	Argv: func() interface{} { return new(fjT) },
	Fn:   fjCLI,

	CanSubRoute: true,
}
