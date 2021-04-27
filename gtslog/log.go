package gtslog

import (
	"github.com/mysticbinary/gotoscaffold/gtsfile"
	"github.com/mysticbinary/gotoscaffold/gtstime"
	"os"
	"strings"
)

// 写入格式：日期 INFO msg
func LogINFO(fileName string, msg string) error{
	var newMsg []string
	newMsg = append(newMsg, gtstime.GetNow())
	newMsg = append(newMsg, "INFO")
	newMsg = append(newMsg, msg)
	msgs := strings.Join(newMsg, " ")
	err := gtsfile.WriteAtEnd(fileName, msgs)
	return err
}

// 写入格式：日期 WARNING msg
func LogWARNING(fileName string, msg string) error{
	var newMsg []string
	newMsg = append(newMsg, gtstime.GetNow())
	newMsg = append(newMsg, "WARNING")
	newMsg = append(newMsg, msg)
	msgs := strings.Join(newMsg, " ")
	err := gtsfile.WriteAtEnd(fileName, msgs)
	return err
}

// 写入格式：日期 ERROR msg
// 这个会退出程序，请谨用
func LogERROR(fileName string, msg string) {
	var newMsg []string
	newMsg = append(newMsg, gtstime.GetNow())
	newMsg = append(newMsg, "ERROR")
	newMsg = append(newMsg, msg)
	newMsg = append(newMsg, "(An error was encountered and the program exited automatically for security reasons.)")
	msgs := strings.Join(newMsg, " ")
	gtsfile.WriteAtEnd(fileName, msgs)
	os.Exit(1)
}
