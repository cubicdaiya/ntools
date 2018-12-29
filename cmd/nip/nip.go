package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func usage() {
	flag.Usage()
	os.Exit(1)
}

func ten2bin(octet string) (string, error) {
	n, err := strconv.Atoi(octet)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%08b", n), nil
}

func main() {
	var (
		ipStr string
	)
	flag.StringVar(&ipStr, "ip", "", "IP address")
	flag.Parse()

	if ipStr == "" {
		usage()
	}

	if ip := net.ParseIP(ipStr); ip == nil {
		usage()
	}

	octets := strings.Split(ipStr, ".")
	bins := make([]string, net.IPv4len)
	for i, octet := range octets {
		sb, err := ten2bin(octet)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		bins[i] = sb

	}
	fmt.Printf("%s.%s.%s.%s\n", bins[0], bins[1], bins[2], bins[3])
}
