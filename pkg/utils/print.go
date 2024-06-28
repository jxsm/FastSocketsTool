package utils

import (
	"FastSocketsTool/prompt"
	"fmt"
	"github.com/zimolab/charsetconv"
	"strings"
)

func AssignedCodedDecodeF(data []byte, srcCharset charsetconv.Charset) string {
	encodeString := AssignedCodedDecode(data, srcCharset)
	builder := strings.Builder{}
	builder.WriteString(encodeString)
	if data[len(data)-1] != 10 {
		builder.WriteString("\n")
	}
	return builder.String()
}

func AssignedCodedDecode(data []byte, srcCharset charsetconv.Charset) string {
	encodeString, err := charsetconv.DecodeToString(data, srcCharset)
	if err != nil {
		prompt.Prompt("the_encoding_conversion_failed")
		fmt.Println("AssignedCodedDecode(Err):", err)
		return string(data)
	}
	return encodeString
}

func AssignedCodedEncode(str string, srcCharset charsetconv.Charset) string {
	encodeString, err := charsetconv.EncodeString(str, srcCharset)
	if err != nil {
		prompt.Prompt("the_encoding_conversion_failed")
		fmt.Println("AssignedCodedDecodeF(Err):", err)
		return str
	}
	return string(encodeString)
}
