package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"

	"github.com/net-byte/go-gateway"
)

func exit() {
	fmt.Print("Press enter key to exit")
	fmt.Scanln()
	os.Exit(0)
}

func main() {
	ip, err := gateway.DiscoverGatewayIPv4()
	if err != nil {
		fmt.Println("Unable to discover Gateway IP")
		fmt.Println()
		exit()
	}

	for {
		fmt.Print(fmt.Sprint("Device IP: (", ip, ") "))
		var input string
		fmt.Scanln(&input)
		fmt.Println()
		if input == "" {
			break
		}
		ip = net.ParseIP(input)
		if ip != nil {
			break
		}
	}

	var port int

	for {
		fmt.Print("Enter port number between 0 to 65535: (5555) ")
		var input string
		fmt.Scanln(&input)
		fmt.Println()
		if input == "" {
			port = 5555
			break
		}
		port, err = strconv.Atoi(input)
		if err == nil && port >= 0 && port <= 65535 {
			break
		}
	}

	// run following command to list all connected devices with their identifiers
	// adb devices

	cmd := exec.Command("adb", "-s", "<device_serial>", "tcpip", fmt.Sprint(port))
	output, _ := cmd.CombinedOutput()
	fmt.Print(string(output))
	fmt.Println()

	cmd = exec.Command("scrcpy", fmt.Sprint("--tcpip=", ip, ":", port))
	output, _ = cmd.CombinedOutput()
	fmt.Print(string(output))
	fmt.Println()

	exit()
}
