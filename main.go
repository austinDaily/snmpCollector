package main

import (
	"fmt"
	"github.com/austinDaily/snmpCollector/collector"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: snmp_collector <ip> <community>")
		return
	}

	ip := os.Args[1]
	community := os.Args[2]

	err := collector.CollectBasicInfo(ip, community)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
