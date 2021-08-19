// Package provides Console Command Operation
package main

import (
	"fmt"

	"github.com/fatih/color"
)

// // ///////////////////////////////////////////////////////////////////
//                Define The Log/Console fmt.Println Level                //
// /////////////////////////////////////////////////////////////////////
func Default(s ...interface{}) {
	fmt.Println(s)
}
func Info(s ...interface{}) {
	color.Set(color.FgCyan)
	fmt.Println(s)
	color.Unset()
}
func Success(s ...interface{}) {
	color.Set(color.FgGreen)
	fmt.Println(s)
	color.Unset()
}
func Warning(s ...interface{}) {
	color.Set(color.FgYellow)
	fmt.Println(s)
	color.Unset()
}
func Error(s ...interface{}) {
	color.Set(color.FgRed)
	fmt.Println(s)
	color.Unset()
}
func Debug(s ...interface{}) {
	color.Set(color.FgWhite)
	color.Set(color.BgCyan)
	fmt.Println(s)
	color.Unset()
}
