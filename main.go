package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/net-byte/go-gateway"
)

func printMessageAndExit(msg string) {
	fmt.Println(msg)
	fmt.Print("Press enter key to exit")
	fmt.Scanln()
	os.Exit(0)
}

func main() {
	gateway, err := gateway.DiscoverGatewayIPv4()
	if err != nil {
		printMessageAndExit("Unable to discover Gateway IP")
	}

	fmt.Println("Gateway IP: ", gateway)

	var port int

	for {
		fmt.Print("Enter port number between 0 to 65535: (5555) ")
		var input string
		fmt.Scanln(&input)
		if input == "" {
			port = 5555
			break
		}
		port, err = strconv.Atoi(input)
		if err == nil && port >= 0 && port <= 65535 {
			break
		}
	}

	cmd := exec.Command("adb", "connect", fmt.Sprint(gateway, ":", port))
	output, err := cmd.CombinedOutput()
	if err != nil {
		printMessageAndExit(fmt.Sprint(err))
	}

	fmt.Println()
	fmt.Print(string(output))

	cmd = exec.Command("scrcpy")
	output, err = cmd.CombinedOutput()
	if err != nil {
		printMessageAndExit(fmt.Sprint(err))
	}

	fmt.Println()
	fmt.Print(string(output))

	printMessageAndExit("")
}
