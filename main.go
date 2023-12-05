package main

import (
	"fmt"
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
	gateway, err := gateway.DiscoverGatewayIPv4()
	if err != nil {
		fmt.Println("Unable to discover Gateway IP")
		fmt.Println()
		exit()
	}

	fmt.Println("Gateway IP: ", gateway)
	fmt.Println()

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

	cmd = exec.Command("scrcpy", fmt.Sprint("--tcpip=", gateway, ":", port))
	output, _ = cmd.CombinedOutput()
	fmt.Print(string(output))
	fmt.Println()

	exit()
}
