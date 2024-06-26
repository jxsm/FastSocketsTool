package prompt

import "fmt"

// EnglishPrompt English prompt
type EnglishPrompt struct {
	promptMap map[string]string
}

func (c *EnglishPrompt) setPrompt() {
	c.AddPrompt("connection_address_IsEmpty", "Connection address is empty use -h designation")
	c.AddPrompt("port_IsEmpty", "Port is empty use -p designation")
	c.AddPrompt("port_is_OutOfRange", "Port number is out of range, only between 1-65535")
	c.AddPrompt("failed_to_connect_to_the_server", "Failed to connect to the server")
	c.AddPrompt("the_encoding_conversion_failed", "Encoding conversion failed")
	c.AddPrompt("send_failure", "Failed to send, connection may be lost")
	c.AddPrompt("input_failure", "Failed to get user input")
	c.AddPrompt("connection_address_is_invalid", "The connection address is invalid and specified with -h")
	c.AddPrompt("unable_to_create_server", "Unable to create server")
	c.AddPrompt("invalid_ip_input", "invalid ip input")
	c.AddPrompt("no_previousConn", "No previous connection available")

}

func (c *EnglishPrompt) Prompt(presuppose string) {
	fmt.Print("[error]:")
	fmt.Println(c.promptMap[presuppose])
}

func (c *EnglishPrompt) InitPrompt() {
	c.promptMap = make(map[string]string)
	c.setPrompt()
}

func (c *EnglishPrompt) AddPrompt(key string, value string) {
	c.promptMap[key] = value
}
