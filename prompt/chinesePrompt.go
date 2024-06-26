package prompt

import "fmt"

// ChinesePrompt 中文提示
type ChinesePrompt struct {
	promptMap map[string]string
}

func (c *ChinesePrompt) setPrompt() {
	c.AddPrompt("connection_address_IsEmpty", "连接地址不能为空,使用-a指定")
	c.AddPrompt("port_IsEmpty", "端口号不能为空,使用-p指定")
	c.AddPrompt("port_is_OutOfRange", "端口号超出范围,只能在1-65535之间")
	c.AddPrompt("failed_to_connect_to_the_server", "无法连接到服务器")
	c.AddPrompt("the_encoding_conversion_failed", "编码转换失败,该消息不会被发送")
	c.AddPrompt("send_failure", "发送失败,连接可能断开")
	c.AddPrompt("input_failure", "获取用户的输入失败")
}

func (c *ChinesePrompt) Prompt(presuppose string) {
	fmt.Println(c.promptMap[presuppose])
}

func (c *ChinesePrompt) InitPrompt() {
	c.promptMap = make(map[string]string)
	c.setPrompt()
}

func (c *ChinesePrompt) AddPrompt(key string, value string) {
	c.promptMap[key] = value
}