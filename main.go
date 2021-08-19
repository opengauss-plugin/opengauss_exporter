package main

import (
	"flag"
)

var DATA_SOURCE_NAME string // 采集器目标数据库："postgresql://exporter:XXXXXXXX@localhost:5432/postgres?sslmode=disable"
var host string             // 采集器监听端口号：默认9432
var port string             // 采集器监听端口号：默认9432

func main() {

	flag.StringVar(&DATA_SOURCE_NAME, "dburi", `"postgresql://postgres@localhost:5432/postgres?sslmode=disable"`, `目标数据库URI地址`)
	flag.StringVar(&host, "h", "0.0.0.0", "采集器主机")
	flag.StringVar(&port, "p", "9432", "采集器端口")
	flag.Parse()

	defer func() {
		//捕获func抛出的panic，防止主程序崩溃
		if err := recover(); err != nil {
			Error("OpenGuass Exporter发生错误：", err)
		}
	}()

	// 初始化OpenGauss链接
	og_slow_select_count()

	// 启动RESTfulAPIs监听 /metrics
	go RestApi(":" + port)

	// 启动命令行工具条
	ConsoleInit(true)

}
