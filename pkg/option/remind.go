package option

import "fmt"

var remindMap = make(map[string]string)

func Remind(key string) {
	fmt.Println(remindMap[key])
}

func init() {
	//AddRemind("ConnectionTypeInputError", "Connection type input error can only be tcp or udp")
	AddRemind("ConnectionUrlIsEmpty", "Connection url is empty")
	AddRemind("PortIsEmpty", "Port is empty")
	AddRemind("PortIsOutOfRange", "PortIsOutOfRange")
}

func AddRemind(key string, value string) {
	remindMap[key] = value
}
