package ipSeeker

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pangkunyi/plum/files"
)

const (
	DEFAULT_DATA_FILE = "ip.seeker.dat"
)

var (
	MUL_NUM   = []int64{int64(256 * 256 * 256), int64(256 * 256), int64(256), int64(1)}
	DATA_FILE = DEFAULT_DATA_FILE
	ipDatas   = make([]*IpData, 0)
)

func InitIpData() {
	if err := files.ScanFile(DATA_FILE, func(line string) error {
		ipDatas = append(ipDatas, NewIpData(line))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	if len(ipDatas) < 1 {
		log.Fatal(fmt.Errorf("failed to load ip datas"))
	}
}

type IpData struct {
	Start    int64
	End      int64
	Shortcut string
	Mcc      string
	Mnc      string
	Carrier  string
}

func (this *IpData) Compare(ipValue int64) int {
	if this.Start > ipValue {
		return 1
	} else if this.End < ipValue {
		return -1
	}
	return 0
}

func Ip2Int64(ip string) int64 {
	if ip == "" {
		return 0
	}
	ips := strings.Split(ip, ",")
	if len(ips) < 1 {
		return 0
	}
	fields := strings.Split(strings.TrimSpace(ips[len(ips)-1]), ".")
	if len(fields) != 4 {
		return 0
	}
	value := int64(0)
	for i := 0; i < 4; i++ {
		if num, err := strconv.ParseInt(fields[i], 10, 64); err != nil {
			return 0
		} else {
			value = value + num*MUL_NUM[i]
		}
	}
	return value
}

func NewIpData(line string) *IpData {
	fields := strings.Fields(line)
	if len(fields) != 6 {
		log.Fatal(fmt.Errorf("invalid ip data line: %s", line))
	}
	startIp, endIp, shortcut, mcc, mnc, carrier := fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]
	return &IpData{Start: Ip2Int64(startIp), End: Ip2Int64(endIp), Shortcut: shortcut, Mcc: mcc, Mnc: mnc, Carrier: carrier}
}

func Seek(ip string) *IpData {
	ipValue := Ip2Int64(ip)
	if ipValue == 0 {
		return nil
	}
	l := len(ipDatas)
	start := 0
	end := l - 1
	for start <= end {
		mid := (end + start) / 2
		ipData := ipDatas[mid]
		switch ipData.Compare(ipValue) {
		case 0:
			return ipData
		case 1:
			end = mid - 1
		case -1:
			start = mid + 1
		}
	}
	return nil
}
