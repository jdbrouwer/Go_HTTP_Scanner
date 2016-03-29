package main

//this is a practice example from the book "network programming with go"
import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)

	result, err := ioutil.ReadAll(conn)

	fmt.Println(string(result))
	os.Exit(1)

}

//CheckError is used for some basic error handeling, nothing that is extremely importante
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		os.Exit(1)
	}
}
