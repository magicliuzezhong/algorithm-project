//
// Package test_project
// @Description：
// @Author：liuzezhong 2021/8/12 16:27
// @Company cloud-ark.com
//
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	/*fmt.Println(1e6)

	var server = &http.Server{
		Addr:    ":8080",
	}
	go func() {
		time.Sleep(time.Second * 3)
		//server.Shutdown(nil)
		server.Close()
	}()
	var err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}*/

	var model = OtaModel{
		Version: "1.0.0",
		Url:     "https://localhost:8081/test/file",
		//FileLength: 45,
		//FileLength: 2157590355,
		FileLength: 2157590355,
	}
	fmt.Println(model)
	var ch = download(model)
	for val := range ch {
		fmt.Println("progress: ", val.Progress, ", msg :  ", val.Msg, ", code:  ", val.Code)
	}

	//var file , _ = os.Open("/Users/liuzezhong/home/test.mp4")
	//var stat , _ = file.Stat()
	//fmt.Println(stat.Size())
}

//
// OtaModel
// @Description: Ota Model
//
type OtaModel struct {
	// Version 版本号
	Version string `json:"version"`
	// Url 下载地址
	Url string `json:"url"`
	// CheckSum 校验码
	CheckSum string `json:"check_sum"`
	// FileLength 文件长度
	FileLength int64 `json:"file_length"`
}

type ProgressCode int

const (
	Fail        ProgressCode = 0 //失败
	Success     ProgressCode = 1 //成功
	Downloading ProgressCode = 2 //下载中
)

//
// OtaUpgradeProgress
// @Description: OTA升级进度
//
type OtaUpgradeProgress struct {
	// Code 状态码 0：失败，1：成功，2：下载中，progress代表了进度
	Code ProgressCode `json:"code"`
	// Msg 状态信息
	Msg string `json:"msg"`
	// Progress 下载进度
	Progress int `json:"progress"`
}

func download(model OtaModel) <-chan OtaUpgradeProgress {
	var resultCh = make(chan OtaUpgradeProgress, 1)
	if "1.0.0" != model.Version || model.Url == "" {
		resultCh <- createErrProgress(errors.New("current version is latest version"))
		return resultCh
	}
	var downloadErr = make(chan error)
	var progress = 0
	go func() {
		var file, fileErr = os.Create("/Users/liuzezhong/home/test.mp4")
		if fileErr != nil {
			downloadErr <- fileErr
			return
		}
		defer file.Close()
		var response, responseErr = http.Get(model.Url)
		if responseErr != nil {
			downloadErr <- responseErr
			return
		}
		defer response.Body.Close()
		var responseLength = response.ContentLength
		if model.FileLength < responseLength {
			downloadErr <- errors.New("fileLength has error")
			return
		}
		reader := bufio.NewReaderSize(response.Body, 32*1024) // 获得get请求响应的reader对象
		writer := bufio.NewWriter(file)                       // 获得文件的writer对象
		var buffer = make([]byte, 32*1024)
		var downloadFileLen int64 = 0
		for {
			var readIndex, readErr = reader.Read(buffer)
			if readErr != nil {
				if readErr != io.EOF {
					downloadErr <- readErr
				} else {
					if downloadFileLen > model.FileLength { //文件总长度
						downloadErr <- errors.New("file length has error")
					} else if downloadFileLen < model.FileLength {
						downloadErr <- errors.New("file length is sort")
					}
				}
				return
			}
			var writeIndex, writeErr = writer.Write(buffer[0:readIndex])
			if writeErr != nil {
				downloadErr <- writeErr
				return
			}
			downloadFileLen += int64(writeIndex) //这里能统计数据
			progress = int(float64(downloadFileLen) / float64(model.FileLength) * 100)
		}
	}()
	go func() {
		defer func() {
			close(downloadErr)
			close(resultCh)
		}()
		var downloadMaxTimer = time.After(time.Second * 30)   //30秒都没下载完成那么直接超时
		var ticker = time.NewTicker(time.Millisecond * 100).C //每100毫秒上报一次
		for {
			select {
			case dlErr := <-downloadErr:
				resultCh <- createErrProgress(dlErr)
				return
			case <-ticker:
				resultCh <- OtaUpgradeProgress{
					Code:     Downloading,
					Msg:      "downloading",
					Progress: progress,
				}
				if progress >= 100 {
					resultCh <- OtaUpgradeProgress{
						Code: Success,
						Msg:  "ready reboot iot device",
					}
					return
				}
			case <-downloadMaxTimer:
				resultCh <- createErrProgress(errors.New("download file timeout"))
				return
			}
		}
	}()
	return resultCh
}

func createErrProgress(err error) OtaUpgradeProgress {
	return OtaUpgradeProgress{
		Code: Fail,
		Msg:  err.Error(),
	}
}
