package ip

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pangkunyi/plum/files"
)

const (
	defaultDataFile = "ip.seeker.dat"
)

var (
	multipleNums = []int64{int64(256 * 256 * 256), int64(256 * 256), int64(256), int64(1)}
	datas        = make([]*Data, 0)
)

//InitIPSeeker init ip seeker
func InitIPSeeker(dataFile string) error {
	if dataFile == "" {
		dataFile = defaultDataFile
	}
	if err := files.ScanFile(dataFile, func(line string) error {
		datas = append(datas, newIPData(line))
		return nil
	}); err != nil {
		return err
	}
	if len(datas) < 1 {
		return fmt.Errorf("empty ip datas, please check ip data file[%s]", dataFile)
	}
	return nil
}

//Data ip data struct
type Data struct {
	Start    int64
	End      int64
	Shortcut string
	Mcc      string
	Mnc      string
	Carrier  string
}

//Compare with two ip values
func (data *Data) Compare(ipValue int64) int {
	if data.Start > ipValue {
		return 1
	} else if data.End < ipValue {
		return -1
	}
	return 0
}

//IP2Int64 convert ip string ip numberical value
func IP2Int64(ip string) int64 {
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
		num, err := strconv.ParseInt(fields[i], 10, 64)
		if err != nil {
			return 0
		}
		value = value + num*multipleNums[i]
	}
	return value
}

func newIPData(line string) *Data {
	fields := strings.Split(line, "\x01")
	if len(fields) != 6 {
		log.Fatal(fmt.Errorf("invalid ip data line: %s", line))
	}
	startIP, err1 := strconv.ParseInt(fields[0], 10, 64)
	endIP, err2 := strconv.ParseInt(fields[1], 10, 64)
	if err1 != nil || err2 != nil {
		log.Fatal(fmt.Errorf("invalid ip data line: %s", line))
	}
	shortcut, mcc, mnc, carrier := fields[2], fields[3], fields[4], fields[5]
	return &Data{Start: startIP, End: endIP, Shortcut: shortcut, Mcc: mcc, Mnc: mnc, Carrier: carrier}
}

//Seek seek ip data with an ip
func Seek(ip string) *Data {
	ipValue := IP2Int64(ip)
	if ipValue == 0 {
		return nil
	}
	l := len(datas)
	start := 0
	end := l - 1
	for start <= end {
		mid := (end + start) / 2
		data := datas[mid]
		switch data.Compare(ipValue) {
		case 0:
			return data
		case 1:
			end = mid - 1
		case -1:
			start = mid + 1
		}
	}
	return nil
}
