// Package phone provides Console Command Operation
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// // ///////////////////////////////////////////////////////////////////
//                      Console Display Function                      //
// /////////////////////////////////////////////////////////////////////

// // ///////////////////////////////////////////////////////////////////
//                  Launch The Console User Interface                 //
// /////////////////////////////////////////////////////////////////////
var isInit = false

func ConsoleInit(Console_Status bool) {
	//启动命令行
	if Console_Status {

		running := true
		reader := bufio.NewReader(os.Stdin)
		for running {
			defer func() {
				if err := recover(); err != nil {
					Error("Console发生错误：", err)
				}
			}()

			color.Set(color.FgCyan)
			fmt.Printf("Console->")
			if !isInit {
				fmt.Printf(" <you can type 'h' for help>")
				isInit = true
			}
			color.Unset()

			data, _, _ := reader.ReadLine()
			command := string(data)

			switch command {
			case "help", "h":
				Info(`
欢迎您使用：面向OpenGuass服务器的 Prometheus 监控采集器
Prometheus exporter for OpenGauss server metrics.

支持版本 Supported versions:
* OpenGauss >= 2.0.1.

注意：并非所有的方法支持OpenGauss 2.0.1以下的版本
NOTE: Not all collection methods are supported on OpenGauss < 2.0.1
				`)
			case "exit", "q":
				//退出服务端
				running = false
				Info("Stopped Exporter")
			default:
				//logger.Println("command", command)
				Warning("Can't Found this Command: " + command)
			}

		}

	}
}
