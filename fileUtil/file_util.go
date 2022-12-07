package fileUtil

import (
	"io"
	"mime/multipart"
	"os"
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
