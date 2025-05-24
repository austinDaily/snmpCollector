package collector

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"time"
)

func CollectBasicInfo(target string, community string) error {
	g := &gosnmp.GoSNMP{
		Target:    target,
		Port:      161,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   3,
	}

	err := g.Connect()
	if err != nil {
		return fmt.Errorf("Connect error: %v", err)
	}
	defer g.Conn.Close()

	// Query some standard OIDs
	oids := []string{
		".1.3.6.1.2.1.1.1.0", // sysDescr
		".1.3.6.1.2.1.1.3.0", // sysUpTime
		".1.3.6.1.2.1.1.5.0", // sysName
	}

	result, err := g.Get(oids)
	if err != nil {
		return fmt.Errorf("SNMP GET error: %v", err)
	}

	for _, variable := range result.Variables {
		fmt.Printf("%s = ", variable.Name)

		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("%s\n", string(variable.Value.([]byte)))
		default:
			fmt.Printf("%v\n", variable.Value)
		}
	}

	return nil
}
