package gtsfile

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 将结构序列化成文本，覆盖写入
func WriteStructToFile(filePath string, Struct interface{}) error {
	if filePath == "" {
		return fmt.Errorf("File Path is empty. ")
	}
	err := CreateFile(filePath)
	if err != nil {
		return err
	}
	jsonData, _ := json.Marshal(Struct)
	err2 := ioutil.WriteFile(filePath, jsonData, 0755)
	if err2 != nil {
		return err2
	}
	return nil
}

// 将Json反序列化回结构，返回接口，需要自己作一下类型转换
func ReadFileToStruct(filePath string, Struct interface{}) (interface{}, error) {
	if filePath == "" {
		return nil, fmt.Errorf("File Path is empty. ")
	}
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(dat, &Struct)
	if err2 != nil {
		return nil, err2
	}
	return Struct, nil
}

// Turn all lines of text into an []string
func TextLinesToSlice(fileName string) ([]string, error) {
	var (
		allLines []string
	)
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			allLines = append(allLines, line)
		}
		if err != nil {
			if err == io.EOF {
				return allLines, nil
			}
			return nil, err
		}
	}
	return nil, err
}

// 创建文件，可以自动判断文件夹是否存在
func CreateFile(fileName string) error {
	file, error := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if error != nil {
		dirPath := CreateDir(fileName)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
		file2, error := os.Create(fileName)
		if error != nil {
			return error
		}
		defer file2.Close()
	}
	defer file.Close()
	return nil
}

// 创建文件夹，传入的格式为"./test/1/2/3.txt" or "./test/1/2/"
func CreateDir(fileName string) string {
	nameSlice := strings.Split(fileName, "/")
	var nfileName string
	nfileName = strings.Join([]string{nfileName, nameSlice[0]}, "")
	for i, v := range nameSlice {
		fmt.Println(v)
		if i < len(nameSlice)-1 && i != 0 {
			nfileName = strings.Join([]string{nfileName, v}, "/")
		}
	}
	return nfileName
}

// 判断文件夹/文件是否存在  存在返回:true
func IsDirExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// 在文本的最后行追加写入
func WriteAtEnd(fileName string, content ...string) error {
	if !IsDirExist(fileName) {
		CreateFile(fileName)
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}
	for _, v := range content {
		if _, err = file.WriteString(v + "\n"); err != nil {
			return fmt.Errorf("An error occurred while writing text. ")
		}
	}
	defer file.Close()
	return nil
}

// 删除文件
func RemoveFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

// 删除指定path下的所有文件
func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
