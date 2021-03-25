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

func IP4bins(ipStr string) []string {
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
	return bins
}

func main() {
	var (
		ipStr string
	)
	flag.StringVar(&ipStr, "ip", "", "IPv4 address or CIDR")
	flag.Parse()

	if ipStr == "" {
		usage()
	}

	if ip, _, err := net.ParseCIDR(ipStr); err == nil {
		cidrs := strings.Split(ipStr, "/")
		bins := IP4bins(ip.String())
		fmt.Printf("%s.%s.%s.%s/%s\n", bins[0], bins[1], bins[2], bins[3], cidrs[1])
	} else if ip := net.ParseIP(ipStr); ip != nil {
		bins := IP4bins(ipStr)
		fmt.Printf("%s.%s.%s.%s\n", bins[0], bins[1], bins[2], bins[3])
	} else {
		usage()
	}
}
