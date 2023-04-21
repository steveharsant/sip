package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unable to get hostname"
	}

	fmt.Print(hostname, "\n\n")

	// Get interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Unable to get interfaces")
		return
	}

	// Iterate over interfaces, if the interface is 'up', get the IPv4 address
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 {

			ips, err := iface.Addrs()
			if err != nil {
				fmt.Println("Failed to get IP addresses for this device:", err)
				return
			}

			for _, ip := range ips {
				ipNet, ok := ip.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {

					re := regexp.MustCompile(`\(([^)]+)\)`)
					matches := re.FindStringSubmatch(iface.Name)

					if len(matches) > 1 {
						ifaceName := matches[1]
						fmt.Printf("%-16s: %s\n", ifaceName, ipNet.IP.String())
					} else {
						ifaceName := iface.Name
						fmt.Printf("%-16s: %s\n", ifaceName, ipNet.IP.String())
					}
				}
			}
		}
	}
}
