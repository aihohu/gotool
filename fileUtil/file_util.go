package fileUtil

import (
	"bufio"
	"fmt"
	"github.com/aihohu/gotool/httpUtil"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// SaveMkdirAllFile 保存文件，如果文件路径不存在，就创建
func SaveMkdirAllFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = os.Stat(dst)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		if err != nil {
			return err
		}
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// FileExist 验证文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// DownloadFile 下载文件
// @param url      文件地址
// @param filePath 存储路径
// @return 下载结果 true成功 false 失败
func DownloadFile(url, filePath string) bool {
	fileName := path.Base(url)
	return DownloadRenameFile(url, filePath, fileName)
}

// DownloadRenameFile 下载文件并重命名
// @param url      文件地址
// @param filePath 存储路径
// @param fileName 文件名
// @return 下载结果 true成功 false 失败
func DownloadRenameFile(url, filePath, fileName string) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("DownloadFileRename failed, err:%v\n\n", err)
		return false
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Printf("DownloadFileRename failed, err:%v\n\n", err)
		return false
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	_, _ = io.Copy(writer, reader)

	return true
}

// DownloadFileRequest 下载文件
// @param url      文件地址
// @param filePath 存储路径
// @param headers  headers
// @return 下载结果 true成功 false 失败
func DownloadFileRequest(url, filePath string, headers []httpUtil.Header) bool {
	fileName := path.Base(url)
	return DownloadRenameFileRequest(url, filePath, fileName, headers)
}

// DownloadRenameFileRequest 下载文件并重命名
// @param url      文件地址
// @param filePath 存储路径
// @param fileName 文件名
// @param headers headers
// @return 下载结果 true成功 false 失败
func DownloadRenameFileRequest(url, filePath, fileName string, headers []httpUtil.Header) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("DownloadRenameFileRequest failed, err:%v\n\n", err)
		return false
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(resp.Body, 32*1024)

	file, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Printf("DownloadRenameFileRequest failed, err:%v\n\n", err)
		return false
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	_, _ = io.Copy(writer, reader)
	return true
}
