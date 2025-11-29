package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SimpleConsoleWriter struct{}

func (w *SimpleConsoleWriter) Alert(v interface{}) {
	w.output("🚨", v)
}

func (w *SimpleConsoleWriter) Close() error {
	return nil
}

func (w *SimpleConsoleWriter) Debug(v interface{}, fields ...logx.LogField) {
	// 开发环境不输出 debug
}

func (w *SimpleConsoleWriter) Error(v interface{}, fields ...logx.LogField) {
	w.output("❌", v)
}

func (w *SimpleConsoleWriter) Info(v interface{}, fields ...logx.LogField) {
	content := fmt.Sprint(v)

	// 过滤掉 SQL 查询日志
	if strings.Contains(content, "sql query:") {
		return
	}

	// 简化 HTTP 日志
	if strings.Contains(content, "[HTTP]") {
		w.outputHTTP(content)
		return
	}

	w.output("ℹ️", v)
}

func (w *SimpleConsoleWriter) Severe(v interface{}) {
	w.output("💥", v)
}

func (w *SimpleConsoleWriter) Slow(v interface{}, fields ...logx.LogField) {
	w.output("🐌", v)
}

func (w *SimpleConsoleWriter) Stack(v interface{}) {
	fmt.Println(v)
}

func (w *SimpleConsoleWriter) Stat(v interface{}, fields ...logx.LogField) {
	// 不输出 stat 统计日志
}

func (w *SimpleConsoleWriter) output(prefix string, v interface{}) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("%s %s %v\n", timestamp, prefix, v)
}

func (w *SimpleConsoleWriter) outputHTTP(content string) {
	// 解析 HTTP 日志：[HTTP]  200  -  PUT  /api/bank-card/balance - 127.0.0.1:6252 - Apifox/1.0.0
	timestamp := time.Now().Format("15:04:05")

	// 提取关键信息
	parts := strings.Fields(content)
	if len(parts) >= 6 {
		method := parts[4] // PUT
		path := parts[5]   // /api/bank-card/balance
		status := parts[1] // 200

		// 提取 duration
		duration := ""
		for _, part := range parts {
			if strings.Contains(part, "duration=") {
				duration = strings.TrimPrefix(part, "duration=")
				break
			}
		}

		// 根据状态码选择颜色
		statusColor := "32" // 绿色
		if status[0] == '4' {
			statusColor = "33" // 黄色
		} else if status[0] == '5' {
			statusColor = "31" // 红色
		}

		fmt.Printf("%s 🌐 \033[%sm%s\033[0m %s %s",
			timestamp, statusColor, status, method, path)

		if duration != "" {
			fmt.Printf(" (%s)", duration)
		}
		fmt.Println()
	} else {
		fmt.Printf("%s 🌐 %s\n", timestamp, content)
	}
}

// 初始化简洁日志
func InitSimpleLogger() {
	logx.SetWriter(&SimpleConsoleWriter{})
	logx.DisableStat()
}
