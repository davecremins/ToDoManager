package content

import (
	"strings"
)

type ContentOperator interface {
	Seek(offset int64, whence int) (int64, error)
	Read(b []byte) (int, error)
}

const (
	bufferSize = 1024
	SAFE_LOOP  = 10
)

func FindSearchStr(contents ContentOperator, size int64, searchStr string) string {
	return bottomUpSearch(contents, size, searchStr)
}

func bottomUpSearch(contents ContentOperator, size int64, searchStr string) string {
	i := 0
	bufSize := int64(bufferSize)
	if size < bufSize {
		bufSize = size
	}
	buffer := make([]byte, bufSize)
	readPosition := size - bufSize
	var builder string

	for {
		i++
		contents.Seek(readPosition, 0)
		contents.Read(buffer)
		builder = string(buffer) + builder

		if strings.LastIndex(string(buffer), searchStr) >= 0 {
			break
		}

		readPosition -= bufSize
		if i == SAFE_LOOP {
			panic("SAFE LOOP count hit without fulfillment")
		}
	}

	pos := strings.LastIndex(builder, searchStr)
	length := len(builder)
	return builder[pos:length]
}

func GetInitContentWithPlaceholders() string {
	initContent := "TODOs %s\n================\n\nCompleted %s\n====================\n"
	return initContent
}
