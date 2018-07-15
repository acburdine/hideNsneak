package deployer

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func cidrHosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func ParseIPFile(path string) ([]string, error) {
	var ipList []string
	var cidrList []string
	var endNum int
	file, err := os.Open(path)
	if err != nil {
		return ipList, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, ip := range lines {
		ip = strings.TrimSpace(ip)
		if _, _, err := net.ParseCIDR(ip); err == nil {
			fmt.Println(err)
			cidrList, _ = cidrHosts(ip)
			ipList = append(ipList, cidrList...)
		} else if net.ParseIP(ip) != nil {
			ipList = append(ipList, ip)
		} else if strings.Contains(ip, "-") {

			ipRangeList := strings.Split(ip, "-")
			digitList := strings.Split(ipRangeList[0], ".")
			threeNumbers := strings.Join(digitList[:3], ".")
			lastDigit := digitList[3]
			startNum, _ := strconv.Atoi(lastDigit)

			if net.ParseIP(ipRangeList[1]) != nil {
				digitList = strings.Split(ipRangeList[1], ".")
				endNum, _ = strconv.Atoi(digitList[3])
			} else {
				endNum, _ = strconv.Atoi(ipRangeList[1])
			}
			for i := startNum; i <= endNum; i++ {
				incrementToString := strconv.Itoa(i)
				ipList = append(ipList, threeNumbers+"."+incrementToString)
			}
		} else {
			return ipList, fmt.Errorf("Incorrectly formatted ip, ip range, or subnet in file: %s", ip)
		}
	}
	return ipList, nil
}

func normalizeTargets(targets []string) string {
	return strings.Join(targets, " ")
}

func ValidatePorts(ports string) (allports []string, err error) {
	portArray := strings.Split(ports, ",")
	var validatePort, startPort, endPort int

	for _, portString := range portArray {
		if strings.Contains(portString, "-") {
			portRange := strings.Split(portString, "-")
			if len(portRange) != 2 {
				err = fmt.Errorf("Incorrectly formatted port string")
				return
			}
			startPort, err = strconv.Atoi(portRange[0])
			if err != nil {
				err = fmt.Errorf("Incorrectly formatted port string")
				return
			}
			endPort, err = strconv.Atoi(portRange[1])
			if err != nil {
				err = fmt.Errorf("Incorrectly formatted port string")
				return
			}
			if startPort > endPort {
				err = fmt.Errorf("Incorrectly formatted port string")
				return
			}
			for i := startPort; i <= endPort; i++ {
				allports = append(allports, strconv.Itoa(i))
			}
		} else {
			validatePort, err = strconv.Atoi(portString)
			if err != nil {
				err = fmt.Errorf("Incorrectly formatted port string")
				return
			}

			allports = append(allports, strconv.Itoa(validatePort))
		}
	}
	return
}

func generateIPPortList(targets []string, ports []string) []string {
	var ipPortList []string
	for _, port := range ports {
		for _, ip := range targets {
			ipPortList = append(ipPortList, ip+":"+port)
		}
	}
	return ipPortList
}

func splitNmapCommand(ports string, hostFile string, command string, count int, evasive bool) (commandList []string) {
	hosts, _ := ParseIPFile(hostFile)

	if evasive {
		fmt.Println("Dividing hosts and ports for evasion...")
		r := rand.New(rand.NewSource(time.Now().Unix()))
		portList, _ := ValidatePorts(ports)
		ipPortList := generateIPPortList(hosts, portList)

		tempSlice := ipPortList

		targetMap := make(map[int]map[string][]string)

		for range ipPortList {

			tempIndex := r.Intn(len(tempSlice))
			p := tempIndex % count
			if targetMap[p] == nil {
				targetMap[p] = make(map[string][]string)
			}

			portSlice := strings.Split(ipPortList[tempIndex], ":")

			targetMap[p][portSlice[1]] = append(targetMap[p][portSlice[1]], portSlice[0])

			tempSlice = append(tempSlice[:tempIndex], tempSlice[tempIndex+1:]...)
		}
		for _, portIP := range targetMap {
			for port, ips := range portIP {
				tempCommand := command + " -p " + port + " " + normalizeTargets(ips)
				commandList = append(commandList, tempCommand)
			}
		}

	} else {
		hostsPerServer := len(hosts) / count
		remainderForServers := len(hosts) % count

		for i := 0; i < count; i++ {
			var targetHosts []string
			remainder := 0
			if remainderForServers > 0 {
				remainder = 1
				remainderForServers = remainderForServers - 1
			}
			if i == count-1 {
				targetHosts = hosts[hostsPerServer*i:]
			} else {
				targetHosts = hosts[hostsPerServer*i : hostsPerServer*(i+1)+remainder]
			}

			tempCommand := command + " -p " + ports + " " + normalizeTargets(targetHosts)
			commandList = append(commandList, tempCommand)
		}
	}
	return
}
