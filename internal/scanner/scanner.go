package scanner

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Result struct {
	IP       string
	MAC      *string
	Hostname *string
	OS       *string
}

func ScanNetwork(cidr string) ([]Result, error) {
	ips, err := hosts(cidr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CIDR: %w", err)
	}

	var wg sync.WaitGroup
	resultsChan := make(chan Result, len(ips))

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			mac := getMAC(ip, os.Getenv("NETWORK_INTERFACE"))
			if mac != nil {
				hostname := lookupHostname(ip)
				os := detectOS(ip) // stub, returns nil for now
				resultsChan <- Result{
					IP:       ip,
					MAC:      mac,
					Hostname: hostname,
					OS:       os,
				}
			}
		}(ip)
	}

	wg.Wait()
	close(resultsChan)

	var results []Result
	for res := range resultsChan {
		results = append(results, res)
	}
	return results, nil
}

func getMAC(ip, iface string) *string {
	// Try arping first
	mac := arpingWithInterface(ip, iface)
	if mac != nil {
		return mac
	}

	// fallback: ping + arp table lookup
	if err := exec.Command("ping", "-c", "1", "-W", "1", ip).Run(); err != nil {
		fmt.Printf("⚠️ ping error for %s: %v\n", ip, err)
		return nil
	}

	return arpGetMAC(ip)
}

func arpingWithInterface(ip, iface string) *string {
	cmd := exec.Command("arping", "-i", iface, "-c", "1", "-w", "1", ip)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("⚠️ arping error for %s: %v\nOutput: %s\n", ip, err, string(out))
		return nil
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Unicast reply from") {
			parts := strings.Split(line, "[")
			if len(parts) > 1 {
				macPart := parts[1]
				end := strings.Index(macPart, "]")
				if end != -1 {
					mac := macPart[:end]
					return &mac
				}
			}
		}
	}
	return nil
}

func arpGetMAC(ip string) *string {
	out, err := exec.Command("arp", "-n", ip).Output()
	if err != nil {
		fmt.Printf("⚠️ arp error for %s: %v\n", ip, err)
		return nil
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, ip) {
			fields := strings.Fields(line)
			for _, f := range fields {
				if isMAC(f) {
					return &f
				}
			}
		}
	}
	return nil
}

func isMAC(s string) bool {
	_, err := net.ParseMAC(s)
	return err == nil
}

func lookupHostname(ip string) *string {
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		return nil
	}
	hostname := strings.TrimSuffix(names[0], ".")
	return &hostname
}

// detectOS is just a stub now, always returns nil
func detectOS(ip string) *string {
	return nil
}

func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	return ips, nil
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
