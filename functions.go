package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
)

func formatOutput(key, value string) {
	if key == "" {
		fmt.Println(value)
	} else {
		fmt.Printf("%-16s: %s\n", key, value)
	}
}

func getExternalIP() string {

	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "Failed to get external IP"
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "Failed to read response body"
	}

	return string(body)
}

func getHostName() string {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unable to get hostname"
	}

	return hostname
}

func getInterfaceIPs(filter string) {

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

					var ifaceName string
					if len(matches) > 1 {
						ifaceName = matches[1]
					} else {
						ifaceName = iface.Name
					}

					regex := regexp.MustCompile(filter)
					if regex.MatchString(ifaceName) {
						formatOutput(ifaceName, ipNet.IP.String())
					}
				}
			}
		}
	}
}

func getMainIP() string {
	hostname := getHostName()

	addresses, err := net.LookupIP(hostname)
	if err != nil {
		return "Failed to lookup IP"
	}

	var mainIP net.IP
	for _, address := range addresses {
		if ipv4 := address.To4(); ipv4 != nil {
			mainIP = ipv4
			break
		}
	}

	if mainIP != nil {
		return mainIP.String()
	} else {
		return "Unable to get main IP"
	}
}
