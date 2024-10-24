package main

import (
	"net"
)

func identifyUserByMAC() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		msg := "cant get net.Interfaces()"
		dumpErrorToFile(msg)
		panic(msg)
	}

	var mac string
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && len(iface.HardwareAddr) == 6 {
			mac = iface.HardwareAddr.String()
			break
		}
	}

	var username string
	switch mac {
	case "":
		msg := "MAC string is empty"
		dumpErrorToFile(msg)
		panic(msg)
	case "18:03:73:cc:8d:8b":
		username = "tester"
	case "cc:52:af:41:22:8d":
		username = "tester"
		// co do zasady na roznych komputerach powinna byc rozna nazwa kopii (tu olewam bo chodzi testera)
	case "d4:be:d9:9c:40:91": // stara kornelia TODO wywalic po przelozeniu dysku funi do kompa po nikodemie
		username = "Kornelia"
	case "18:66:da:23:75:83": // stary nikodem, nowa kornelia
		username = "Kornelia"
	case "10:7c:61:4c:7e:7e": // nowy nikodem
		username = "Nikodem"
	default:
		msg := "unknown mac address: " + mac
		dumpErrorToFile(msg)
		panic(msg)
	}

	println("identified user is:", username)
	return username
}
