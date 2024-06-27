package utils

import (
	"FastSocketsTool/prompt"
	"fmt"
	"github.com/zimolab/charsetconv"
	"strings"
)

func AssignedCodedConversionsF(data []byte, srcCharset charsetconv.Charset) string {
	encodeString := AssignedCodedConversions(data, srcCharset)
	builder := strings.Builder{}
	builder.WriteString(encodeString)
	if data[len(data)-1] != 10 {
		builder.WriteString("\n")
	}
	return builder.String()
}

func AssignedCodedConversions(data []byte, srcCharset charsetconv.Charset) string {
	encodeString, err := charsetconv.DecodeToString(data, srcCharset)
	if err != nil {
		prompt.Prompt("the_encoding_conversion_failed")
		fmt.Println("AssignedCodedConversionsF(Err):", err)
		return string(data)
	}
	return encodeString
}
