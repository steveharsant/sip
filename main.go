package main

import (
	"flag"
	"fmt"
)

const version = "1.1.0"

func main() {

	allFlag := flag.Bool("a", false, "Get all IP and host information")
	externalFlag := flag.Bool("e", false, "Get external IP")
	filterFlag := flag.String("f", "", "Filter interface names. Regex supported (use with -i)")
	hostnameFlag := flag.Bool("h", false, "Get hostname")
	interfaceFlag := flag.Bool("i", false, "Get all interface IPs")
	versionFlag := flag.Bool("v", false, "Print version")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		return
	}

	if flag.NFlag() == 0 {
		formatOutput("Hostname", getHostName())
		formatOutput("Address", getMainIP())
		return
	}

	if *allFlag {
		formatOutput("Hostname", getHostName())
		formatOutput("Address", getMainIP())
		formatOutput("External IP", getExternalIP())
		fmt.Println("\n-----------Interfaces-----------")
		getInterfaceIPs(*filterFlag)
		return
	}

	if *hostnameFlag {
		formatOutput("", getHostName())
	}

	if *interfaceFlag {
		getInterfaceIPs(*filterFlag)
	}

	if *externalFlag {
		formatOutput("", getExternalIP())
	}
}
