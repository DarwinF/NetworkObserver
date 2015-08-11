//--------------------------------------------
// ping.go
//
// A tool to ping an IP or web address to test
// if there is a functional connection.
//--------------------------------------------

package tools

import (
	"NetworkObserver/configuration"
	"errors"
	"golang.org/x/net/icmp"
	"golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type pingInfo struct {
	pingDelay   string
	externalip  string
	internalip  string
	externalurl string
}

var file *os.File
var sequence int = 1

func init() {
	file, _ = os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Println("-+--------------------------------------------+-")
}

func Cleanup() {
	file.Close()
}

func Ping(pi pingInfo) pingResponse {
	pr := pingResponse{}
	pr.err = nil

	conn, err := icmp.ListenPacket("udp4", configuration.GetDeviceIP())

	if err != nil {
		msg := "Could not create a packet endpoint to listen on."
		log.Println(msg)
		pr.err = errors.New(msg)
		return pr
	}

	defer conn.Close()

	testNetwork(&pr, pi, conn)
	testInternet(&pr, pi, conn)
	testDNS(&pr, pi, conn)

	// Just a little seperator
	log.Println("-+--------------------------------------------+-")

	sequence++
	return pr
}

// This checks if we can ping an internal IP
func testNetwork(pr *pingResponse, pi pingInfo, conn *icmp.PacketConn) {
	start := time.Now()
	msg := createMessage()

	msg_bytes, err := msg.Marshal(nil)
	if err != nil {
		emsg := "Could not marshal the message to []byte."
		log.Println(emsg)
		pr.err = errors.New(emsg)
		return
	}

	if _, err := conn.WriteTo(msg_bytes, &net.UDPAddr{IP: net.ParseIP(pi.internalip), Zone: "en0"}); err != nil {
		emsg := "Could not write to the internal ip address: " + pi.internalip
		log.Println(emsg)
		pr.internal = false
		pr.err = errors.New(emsg)
		return
	}

	pr.internal = true

	response := make([]byte, 1500)
	count, peer, err := conn.ReadFrom(response)
	if err != nil {
		emsg := "Could not read the response."
		log.Println(emsg)
		pr.internal = false
		pr.err = errors.New(emsg)
		return
	}

	_, err = icmp.ParseMessage(iana.ProtocolICMP, response[:count])
	if err != nil {
		emsg := "Could not parse the message received."
		log.Println(emsg)
		pr.internal = false
		pr.err = errors.New(emsg)
		return
	}

	log.Println("Response " + strconv.Itoa(sequence) + " received from " + peer.String() +
		" after " + time.Now().Sub(start).String())
}

// This checks if we can convert an URL to an IP
func testDNS(pr *pingResponse, pi pingInfo, conn *icmp.PacketConn) {
	start := time.Now()
	msg := createMessage()

	msg_bytes, err := msg.Marshal(nil)
	if err != nil {
		emsg := "Could not marshal the message to []byte."
		log.Println(emsg)
		pr.err = errors.New(emsg)
		return
	}

	ip, _ := net.LookupHost(pi.externalurl)

	if _, err := conn.WriteTo(msg_bytes, &net.UDPAddr{IP: net.ParseIP(ip[0]), Zone: "en0"}); err != nil {
		emsg := "Could not write to the internal ip address: " + ip[0]
		log.Println(emsg)
		pr.external_url = false
		pr.err = errors.New(emsg)
		return
	}

	pr.external_url = true

	response := make([]byte, 1500)
	count, peer, err := conn.ReadFrom(response)
	if err != nil {
		emsg := "Could not read the response."
		log.Println(emsg)
		pr.external_url = false
		pr.err = errors.New(emsg)
		return
	}

	_, err = icmp.ParseMessage(iana.ProtocolICMP, response[:count])
	if err != nil {
		emsg := "Could not parse the message received."
		log.Println(emsg)
		pr.external_url = false
		pr.err = errors.New(emsg)
		return
	}

	log.Println("Response " + strconv.Itoa(sequence) + " received from " + peer.String() +
		" after " + time.Now().Sub(start).String())
}

// This checks if we can ping an external IP
func testInternet(pr *pingResponse, pi pingInfo, conn *icmp.PacketConn) {
	start := time.Now()
	msg := createMessage()

	msg_bytes, err := msg.Marshal(nil)
	if err != nil {
		emsg := "Could not marshal the message to []byte."
		log.Println(emsg)
		pr.err = errors.New(emsg)
		return
	}

	if _, err := conn.WriteTo(msg_bytes, &net.UDPAddr{IP: net.ParseIP(pi.externalip), Zone: "en0"}); err != nil {
		emsg := "Could not write to the external ip address: " + pi.externalip
		log.Println(emsg)
		pr.external_ip = false
		pr.err = errors.New(emsg)
		return
	}

	pr.external_ip = true

	response := make([]byte, 1500)
	count, peer, err := conn.ReadFrom(response)
	if err != nil {
		emsg := "Could not read the response."
		log.Println(emsg)
		pr.external_ip = false
		pr.err = errors.New(emsg)
		return
	}

	_, err = icmp.ParseMessage(iana.ProtocolICMP, response[:count])
	if err != nil {
		emsg := "Could not parse the message received."
		log.Println(emsg)
		pr.external_ip = false
		pr.err = errors.New(emsg)
		return
	}

	log.Println("Response " + strconv.Itoa(sequence) + " received from " + peer.String() +
		" after " + time.Now().Sub(start).String())
}

func createMessage() icmp.Message {
	return icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  sequence,
			Data: []byte(time.Now().String()),
		},
	}
}
