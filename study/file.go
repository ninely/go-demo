package study

import (
	"path"
	"strings"
)

func FilePath() string {
	srcFileName := "../../root/run.sh"
	ext := path.Ext(path.Base(srcFileName))
	filePrefix := strings.TrimSuffix(path.Base(srcFileName), ext)
	return filePrefix
}
