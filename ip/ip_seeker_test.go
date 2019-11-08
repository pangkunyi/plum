package ip

import (
	"testing"
)

func TestIP2Int64(t *testing.T) {
	val := IP2Int64("1.1.1.1")
	if val != 256*256*256+256*256+256+1 {
		t.Errorf("failed:%d", val)
	}
	val = IP2Int64("12.2.2.2, 1.1.1.1")
	if val != 256*256*256+256*256+256+1 {
		t.Errorf("failed:%d", val)
	}
	val = IP2Int64("255.255.255.255")
	if val != 256*256*256*255+256*256*255+256*255+1*255 {
		t.Errorf("failed:%d", val)
	}
	val = IP2Int64("255.255.a.255")
	if val != 0 {
		t.Errorf("failed:%d", val)
	}
}

func TestSeek(t *testing.T) {
	ip := "195.7.16.254"
	ipData := Seek(ip)
	if ipData.Shortcut != "it" {
		t.Errorf("seek ip[%s], failed:%#v", ip, ipData)
	}
}
