package server

import (
	"FastSocketsTool/pkg/option"
	"FastSocketsTool/pkg/utils"
	"FastSocketsTool/prompt"
	"fmt"
	"github.com/zimolab/charsetconv"
	"net"
	"os"
	"strings"
	"sync"
)

type UdpServer struct {
	option        *option.OperationOption
	wg            sync.WaitGroup
	operationFunc map[string]func()
	listener      *net.UDPConn
	previousConn  *net.UDPAddr
}

func (u *UdpServer) StartServer() {
	u.Init()
	u.listener = u.getUDPConn()
	u.wg.Add(2)
	go u.listenToUserInput()
	go u.listenForConnections()
	u.wg.Wait()
}

func (u *UdpServer) Init() {
	u.operationFunc = make(map[string]func())
	u.operationFunc["exit"] = exit
	u.operationFunc["close"] = exit
}

func (u *UdpServer) getUDPConn() *net.UDPConn {
	ip := net.ParseIP(*u.option.ConnectionAddress)
	if ip == nil {
		prompt.Prompt("connection_address_is_invalid")
		os.Exit(1)
	}
	udp, err := net.ListenUDP(GetNetworkString("udp", *u.option.V6), &net.UDPAddr{IP: ip, Port: *u.option.Port})
	if err != nil {
		prompt.Prompt("unable_to_create_server")
		fmt.Println(err)
		os.Exit(1)
	}
	return udp
}

func (u *UdpServer) listenToUserInput() {
	defer u.wg.Done()
	for {
		input, err := utils.UserInput()
		if err != nil {
			prompt.Prompt("input_failure")
			break
		}
		operation := u.operationFunc[strings.Trim(input, "\n")]
		if operation != nil {
			operation()
		} else {
			u.preSend(input)
		}
	}

}

func (u *UdpServer) preSend(str string) {
	index := strings.Index(str, "@")
	if index == -1 {
		u.sendPreviousIp(str)
	} else if index == 0 {
		prefix := strings.TrimPrefix(str, "@")
		u.sendPreviousIp(prefix)
	} else {
		split := strings.Split(str, "@")
		u.sendSpecifiedIp(split[0], split[1])
	}
}

func (u *UdpServer) listenForConnections() {
	defer u.wg.Done()
	data := make([]byte, 1024)
	for {
		udp, u2, _ := u.listener.ReadFromUDP(data[:])
		u.previousConn = u2
		codedPrint := utils.AssignedCodedDecodeF(data[:udp], charsetconv.Charset(*u.option.ReceiveEncode))
		fmt.Printf("[%s]:%s", u2.IP.String(), codedPrint)
	}
}

func (u *UdpServer) sendPreviousIp(str string) {
	if u.previousConn == nil {
		prompt.Prompt("no_previousConn")
		return
	}
	encode := utils.AssignedCodedEncode(str, charsetconv.Charset(*u.option.SendEncode))
	_, err := u.listener.WriteToUDP([]byte(encode), u.previousConn)
	if err != nil {
		fmt.Printf("(%s) There's no such connection.\n", u.previousConn.IP.String())
	}
	fmt.Printf("[you]:%s\n", encode)
}

func (u *UdpServer) sendSpecifiedIp(ipStr string, str string) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		prompt.Prompt("invalid_ip_input")
		return
	}
	encode := utils.AssignedCodedEncode(str, charsetconv.Charset(*u.option.SendEncode))

	_, err := u.listener.WriteToUDP([]byte(encode), &net.UDPAddr{IP: ip, Port: *u.option.Port})
	if err != nil {
		fmt.Printf("(%s) There's no such connection.\n", str)
	}

}
