package util

import (
	"testing"
	"path/filepath"
)

func TestFileType(t *testing.T) {
	t.Skip("Skipped... 정리 필요")
	getFileType(t, `C:\Users\Yongho1037\Documents\Script.zip`)
	getFileType(t, `E:\works\th-maple-stage\BinSvr\Canvas.dll`)
	getFileType(t, `E:\works\th-maple-stage\BinSvr\Canvas.map`)
	getFileType(t, `C:\Users\Yongho1037\Desktop\laszlo_node_1.png`)
	getFileType(t, `C:\Users\Yongho1037\Desktop\ping.bat`)
	getFileType(t, `C:\Users\Yongho1037\Desktop\gcp-summit-todolist.txt`)
}

func getFileType(t *testing.T, path string) {
	fileType := GetFileType(path)
	t.Log("path : " + path + ", " + fileType)
}

func TestFilePath(t *testing.T) {
	t.Skip("Skipped... 정리 필요")
	inputFilePath := `C:\Users\Yongho1037\Desktop\gcp-summit-todolist.asdfasdfasdfasdf.asdfasdfasdf.txt`
	fileName := filepath.Base(inputFilePath)
	outputPath := filepath.Dir(inputFilePath)
	t.Log(fileName)
	t.Log(outputPath)
}
