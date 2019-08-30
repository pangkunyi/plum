package email

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//Validate an email address is a valid email address or not
func Validate(email string, from string) error {
	idx := strings.Index(email, "@")
	if idx < 1 {
		return fmt.Errorf("[%s]invalid email address", email)
	}
	domain := email[idx+1:]
	if domain == "" {
		return fmt.Errorf("[%s]invalid email address", email)
	}

	mxs, err := net.LookupMX(domain)
	if err != nil {
		return fmt.Errorf("[%s]failed to lookup MX record, cuase by %v", email, err)
	}
	if len(mxs) < 1 {
		return fmt.Errorf("[%s]no MX record found", email)
	}
	mx := mxs[len(mxs)-1].Host

	//connect to smtp port 25
	conn, err := net.Dial("tcp", mx+":25")
	if err != nil {
		return fmt.Errorf("[%s]failed to connect mx server[%s] on port 25", email, mx)
	}
	defer conn.Close()
	br := bufio.NewReader(conn)
	status, err := readReply(br)
	//HELO
	conn.Write([]byte("EHLO " + from + "\r\n"))
	status, err = readReply(br)
	if err != nil || !strings.HasPrefix(status, "250") {
		return fmt.Errorf("[EHLO]failed validate email[%s] with mx[%s], status:%s, err:%v", email, mx, status, err)
	}
	//MAIL FROM
	conn.Write([]byte("MAIL FROM: <" + from + ">\r\n"))
	status, err = readReply(br)
	if err != nil || !strings.HasPrefix(status, "250") {
		return fmt.Errorf("[MAIL FROM]failed validate email[%s] with mx[%s], status:%s, err:%v", email, mx, status, err)
	}
	//RCPT TO
	conn.Write([]byte("RCPT TO: <" + email + ">\r\n"))
	status, err = readReply(br)
	if err != nil || !strings.HasPrefix(status, "250") {
		return fmt.Errorf("[RCPT TO]failed validate email[%s] with mx[%s], status:%s, err:%v", email, mx, status, err)
	}
	conn.Write([]byte("RSET\r\n"))
	readReply(br)
	conn.Write([]byte("QUIT\r\n"))
	readReply(br)
	return nil
}
func readReply(br *bufio.Reader) (string, error) {
	var buf strings.Builder
	for {
		content, err := br.ReadBytes('\n')
		if err != nil {
			return "", err
		}
		if len(content) < 4 {
			return "", fmt.Errorf("recieve error msg:%s", string(content))
		}
		buf.Write(content)
		if content[3] == ' ' {
			break
		}
	}
	return buf.String(), nil
}
