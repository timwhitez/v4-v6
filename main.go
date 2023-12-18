package main
import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <ipv4_address>")
		os.Exit(1)
	}
	ipv4Str := os.Args[1]
	ipv6 := mapIPv4ToIPv6(ipv4Str)
	fmt.Println("IPv6 Address:", ipv6)
}

func mapIPv4ToIPv6(ipv4Str string) string {
	// Parse the IPv4 address
	ipv4Addr := net.ParseIP(ipv4Str)
	if ipv4Addr == nil {
		fmt.Println("Invalid IPv4 address")
		os.Exit(1)
	}

	// To4 converts the IPv4 address to 4-byte representation.
	// If the IP address is not IPv4, To4 returns nil.
	ip := ipv4Addr.To4()
	if ip == nil {
		fmt.Println("Invalid IPv4 address")
		os.Exit(1)
	}

	// Creating the IPv4-mapped IPv6 address by appending "::ffff:"
	// and the hexadecimal representation of the IPv4 address.
	return fmt.Sprintf("::ffff:%02X%02X:%02X%02X", ip[0], ip[1], ip[2], ip[3])
}

func intToHex(n int) string {
	return strings.ToUpper(fmt.Sprintf("%x", n))
}
