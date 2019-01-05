package util

import (
	"testing"
	"path/filepath"
)

func TestFilePath(t *testing.T) {
	t.Skip("Skipped... 정리 필요")
	inputFilePath := `C:\Users\Yongho1037\Desktop\gcp-summit-todolist.asdfasdfasdfasdf.asdfasdfasdf.txt`
	fileName := filepath.Base(inputFilePath)
	outputPath := filepath.Dir(inputFilePath)
	t.Log(fileName)
	t.Log(outputPath)
}
