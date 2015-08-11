//--------------------------------------------
// ping.go
//
// A tool to ping an IP or web address to test
// if there is a functional connection.
//--------------------------------------------

package tools

import (
	"NetworkObserver/configuration"
	"golang.org/x/net/icmp"
	"golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
)

type pingInfo struct {
	pingDelay   string
	externalip  string
	internalip  string
	externalurl string
}

var file *os.File
var device_ip string

func init() {
	file, _ = os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
}

func Cleanup() {
	file.Close()
}

func Ping(pi pingInfo) pingResponse {
	c, err := icmp.ListenPacket("udp4", configuration.GetDeviceIP())

	if err != nil {
		log.Print(err)
		log.Println()
		return pingResponse{}
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Print(err)
		log.Println()
		return pingResponse{}
	}

	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP("8.8.8.8"), Zone: "en0"}); err != nil {
		log.Print(err)
		log.Println()
		return pingResponse{}
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Print(err)
		log.Println()
		return pingResponse{}
	}

	rm, err := icmp.ParseMessage(iana.ProtocolICMP, rb[:n])
	if err != nil {
		log.Print(err)
		log.Println()
		return pingResponse{}
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v", peer)
	default:
		log.Printf("got %+v; want echo reply", rm)
	}

	return pingResponse{}
}
