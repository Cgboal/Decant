package main

import (
	"bufio"
	"os"
	"net"
	"fmt"
	"log"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func expandCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	ipList := []string{}
	if err != nil {
		return ipList, err
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ipList = append(ipList, ip.String())
	}

	return ipList, nil
}

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ipRange, err := expandCIDR(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, ipAddr := range ipRange {
			fmt.Println(ipAddr)
		}
	}
}
