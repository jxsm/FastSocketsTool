package client

import (
	"FastSocketsTool/pkg/option"
	"FastSocketsTool/pkg/utils"
	"FastSocketsTool/prompt"
	"fmt"
	"github.com/zimolab/charsetconv"
	"net"
	"os"
	"sync"
)

type Client struct {
	option *option.OperationOption
	dial   net.Conn
	wg     sync.WaitGroup
}

func NewClient() *Client {
	return &Client{}
}

func (s *Client) Start(option *option.OperationOption) {
	s.wg = sync.WaitGroup{}
	s.option = option
	s.dial = connect(option.ConnectionType, option.GetConnectionAddress())
	s.wg.Add(2)
	go s.monitorServerPrint()
	go s.monitorUserInput()
	s.wg.Wait()
}

// connect to a server
//
// connect("tcp","127.0.0.1:8080")
func connect(typeStr string, address string) net.Conn {
	dia, err := net.Dial(typeStr, address)
	if err != nil {
		prompt.Prompt("failed_to_connect_to_the_server")
		fmt.Println(err)
		os.Exit(1)
	}
	return dia
}

func (s *Client) send(message []byte) {
	encodeString, err := charsetconv.DecodeToString(message, charsetconv.Charset(*s.option.SendEncode))
	if err != nil {
		prompt.Prompt("the_encoding_conversion_failed")
		fmt.Println("send(Err):", err)
		return
	}
	_, err = s.dial.Write([]byte(encodeString))
	if err != nil {
		prompt.Prompt("send_failure")
	}
}

func (s *Client) monitorServer(dataByte []byte) ([]byte, error) {
	read, err := s.dial.Read(dataByte)
	if err != nil {
		return nil, err
	}
	return dataByte[:read], nil
}

func (s *Client) monitorUserInput() {
	defer func(dial net.Conn) {
		s.wg.Done()
		_ = dial.Close()
	}(s.dial)
	for {
		input, err := utils.UserInput()
		if err != nil {
			prompt.Prompt("input_failure")
			break
		}
		bytesArray := []byte(input)
		s.send(bytesArray)
		fmt.Print("[you]:", input)
	}
}

func (s *Client) monitorServerPrint() {
	defer func(dial net.Conn) {
		s.wg.Done()
		_ = dial.Close()
	}(s.dial)

	dataByte := make([]byte, 1024)
	for {
		server, err := s.monitorServer(dataByte)
		if err != nil {
			os.Exit(1)
		}
		codedPrint := utils.AssignedCodedConversionsF(server, charsetconv.Charset(*s.option.ReceiveEncode))
		fmt.Print("[server]:", codedPrint)
	}
}
