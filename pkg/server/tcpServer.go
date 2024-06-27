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

type TCPServer struct {
	option        *option.OperationOption
	wg            sync.WaitGroup
	listener      *net.TCPListener
	connectList   map[string]*net.TCPConn
	operationFunc map[string]func()
}

func (s *TCPServer) StartServer() {
	s.Init()
	s.listener = s.getTCPListener()
	s.wg.Add(2)
	go s.startListening()
	go s.listenToUserInput()
	s.wg.Wait()
}

func (s *TCPServer) Init() {
	s.connectList = make(map[string]*net.TCPConn)
	s.operationFunc = make(map[string]func())
	s.operationFunc["list"] = s.showList
	s.operationFunc["exit"] = s.exit
	s.operationFunc["close"] = s.exit
}

func (s *TCPServer) getTCPListener() *net.TCPListener {
	ip := net.ParseIP(*s.option.ConnectionAddress)
	if ip == nil {
		prompt.Prompt("connection_address_is_invalid")
		os.Exit(1)
	}
	tcp, err := net.ListenTCP(GetNetworkString("tcp", *s.option.V6), &net.TCPAddr{IP: ip, Port: *s.option.Port})
	if err != nil {
		prompt.Prompt("unable_to_create_server")
		fmt.Println(err)
		os.Exit(1)
	}
	return tcp
}

func (s *TCPServer) startListening() {
	defer s.wg.Done()
	errNo := 0
	for {
		tcp, err := s.listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			errNo += 1
			if errNo >= 10 {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		s.wg.Add(1)
		go s.listenForConnections(tcp)
		errNo = 0
	}
}

func (s *TCPServer) monitorForConnections(conn *net.TCPConn, dataByte []byte) ([]byte, error) {
	read, err := conn.Read(dataByte)
	if err != nil {
		return nil, err
	}
	return dataByte[:read], nil
}

// listenForConnections Listen to the client's information
func (s *TCPServer) listenForConnections(conn *net.TCPConn) {
	ipConn := getIpConn(conn)
	defer func(conn *net.TCPConn) {
		_ = conn.Close()
		s.wg.Done()
		delete(s.connectList, ipConn)
	}(conn)
	s.wg.Add(1)
	s.connectList[strings.Split(conn.LocalAddr().String(), ":")[0]] = conn
	fmt.Printf("[in]:%s\n", ipConn)
	dataByte := make([]byte, 1024)
	for {
		connections, err := s.monitorForConnections(conn, dataByte)
		if err != nil {
			break
		}
		codedPrint := utils.AssignedCodedConversionsF(connections, charsetconv.Charset(*s.option.ReceiveEncode))
		fmt.Printf("[%s]:%s", getIpConn(conn), codedPrint)
	}
	fmt.Printf("[out]:%s\n", ipConn)
}

func (s *TCPServer) listenToUserInput() {
	defer s.wg.Done()
	for {
		input, err := utils.UserInput()
		if err != nil {
			prompt.Prompt("input_failure")
			break
		}
		operation := s.operationFunc[strings.Trim(input, "\n")]
		if operation != nil {
			operation()
		} else {
			s.preSend(input)
		}
	}

}

// preSend Processing of information before it is sent
func (s *TCPServer) preSend(str string) {

	index := strings.Index(str, "@")
	if index == -1 {
		s.sendAll(str)
	} else if index == 0 {
		prefix := strings.TrimPrefix(str, "@")
		s.sendAll(prefix)
	} else {
		split := strings.Split(str, "@")
		s.sendSpecifiedIp(split[0], split[1])
	}
}

func (s *TCPServer) sendAll(message string) {
	conversions := utils.AssignedCodedConversions([]byte(message), charsetconv.Charset(*s.option.SendEncode))
	var errList []net.Conn
	fmt.Println("[sendAll]:", message)
	for _, conn := range s.connectList {
		_, err := conn.Write([]byte(conversions))
		if err != nil {
			errList = append(errList, conn)
		}
	}
	for _, conn := range errList {
		fmt.Println("send to(" + getIpConn(conn) + ")failure")
	}
}

func (s *TCPServer) sendSpecifiedIp(ip string, message string) {
	conn := s.connectList[getIp(ip)]
	if conn == nil {
		fmt.Println("(" + ip + ") There's no such connection.")
		return
	}

	conversions := utils.AssignedCodedConversions([]byte(message), charsetconv.Charset(*s.option.SendEncode))
	_, err := conn.Write([]byte(conversions))
	ipConn := getIpConn(conn)
	if err != nil {
		fmt.Println("send to(" + ipConn + ")failure")
	}
	fmt.Println("send to["+ipConn+"]:", message)
}

func (s *TCPServer) showList() {
	fmt.Println("Online connection list:")
	for s2 := range s.connectList {
		fmt.Println("--", getIp(s2))
	}
}

func (s *TCPServer) exit() {
	os.Exit(0)
}

func getIp(addr string) string {
	return strings.Split(addr, ":")[0]
}

func getIpConn(conn net.Conn) string {
	return getIp(conn.LocalAddr().String())
}
