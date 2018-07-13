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

func ParseIPFile(path string) []string {
	var ipList []string
	var cidrList []string
	var endNum int
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, ip := range lines {
		if _, _, err := net.ParseCIDR(ip); err == nil {
			cidrList, _ = cidrHosts(ip)
			ipList = append(ipList, cidrList...)
		}
		if net.ParseIP(ip) != nil {
			ipList = append(ipList, ip)
		}
		if strings.Contains(ip, "-") {

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
		}
	}
	return ipList
}

func normalizeTargets(targets []string) string {
	return strings.Join(targets, " ")
}

func GenerateIPPortList(targets []string, ports []string) []string {
	var ipPortList []string
	for _, port := range ports {
		for _, ip := range targets {
			ipPortList = append(ipPortList, ip+":"+port)
		}
	}
	return ipPortList
}

//This is for splitting up hosts more granualarly for stealthier scans
func RandomizeIPPortsToHosts(scannerCount int, ipPortList []string) map[int]map[int][]string {
	nmapTargeting := make(map[int]map[int][]string)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	tempSlice := ipPortList
	for range ipPortList {
		tempIndex := r.Intn(len(tempSlice))
		p := tempIndex % scannerCount
		splitArray := strings.Split(ipPortList[tempIndex], ":")
		port, _ := strconv.Atoi(splitArray[1])
		targetIP := splitArray[0]
		if nmapTargeting[p] == nil {
			ipSlice := []string{targetIP}
			ipSlice = ipSlice[:]
			portMap := make(map[int][]string)
			portMap[port] = ipSlice
			nmapTargeting[p] = portMap
		} else {
			nmapTargeting[p][port] = append(nmapTargeting[p][port], targetIP)
		}
		tempSlice = append(tempSlice[:tempIndex], tempSlice[tempIndex+1:]...)
	}
	return nmapTargeting
}

//This is for splitting up hosts straight up for less stealthy scans
// func splitIPsToHosts(Instances map[int]*Instance, portList []string, ipList []string) {
// 	count := len(Instances)
// 	splitNum := len(ipList) / count
// 	for i := range Instances {
// 		Instances[i].NmapTargets = make(map[string][]string)
// 		Instances[i].NmapTargets = make(map[string][]string)
// 		for _, port := range portList {
// 			if i != count-1 {
// 				Instances[i].NmapTargets[port] = ipList[i*splitNum : (i+1)*splitNum]
// 			} else {
// 				Instances[i].NmapTargets[port] = ipList[i*splitNum:]
// 			}
// 		}
// 	}
// }

// func (instance Instance) parseNmapTargets() (portList []string, ipList []string) {
// 	for _, ipPort := range instance.Nmap.NmapTargets{
// 		splitArray := strings.Split(ipPort, ":")
// 		ipList = removeDuplicateStrings(append(ipList, splitArray[0]))
// 		portList = removeDuplicateStrings(append(portList, splitArray[1]))
// 	}
// 	return
// }
