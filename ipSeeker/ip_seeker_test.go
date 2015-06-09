package ipSeeker

import (
	"testing"
)

func init() {
	DATA_FILE = "/data/ip.seeker.dat"
	InitIpData()
}

func TestIp2Int64(t *testing.T) {
	val := Ip2Int64("1.1.1.1")
	if val != 256*256*256+256*256+256+1 {
		t.Errorf("failed:%d", val)
	}
	val = Ip2Int64("12.2.2.2, 1.1.1.1")
	if val != 256*256*256+256*256+256+1 {
		t.Errorf("failed:%d", val)
	}
	val = Ip2Int64("255.255.255.255")
	if val != 256*256*256*255+256*256*255+256*255+1*255 {
		t.Errorf("failed:%d", val)
	}
	val = Ip2Int64("255.255.a.255")
	if val != 0 {
		t.Errorf("failed:%d", val)
	}
}

func TestSeek(t *testing.T) {
	ipData := Seek("195.7.9.254")
	if ipData.Start != 3272015872 || ipData.End != 3272019967 || ipData.Country != "罗马尼亚" || ipData.Shortcut != "ro" {
		t.Errorf("failed:%d", ipData)
	}
}
