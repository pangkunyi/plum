package stock

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/pangkunyi/plum/encoding"
	"github.com/pangkunyi/plum/http"
)

const (
	SINA_F_URL        = "http://hq.sinajs.cn/list=%s"
	SINA_F_FETCH_SIZE = 100
)

type SinaService struct {
	Service
}

func NewSinaService() *SinaService {
	return &SinaService{}
}

func code2SinaCode(code string) string {
	if strings.HasPrefix(code, "6") {
		return "sh" + code
	} else {
		return "sz" + code
	}
}

func sinaCode2Code(code string) string {
	return code[2:]
}

func (s *SinaService) GetQuote(code string) (Quote, error) {
	qMap, err := s.GetQuotes([]string{code})
	if err != nil {
		return EMPTY_QUOTE, err
	}
	q, ok := qMap[code]
	if !ok {
		return EMPTY_QUOTE, fmt.Errorf("no quote for code %s", code)
	}
	return q, nil
}

func (s *SinaService) GetQuotes(codes []string) (map[string]Quote, error) {
	qMap := make(map[string]Quote)
	idx := 0
	var codesStr string
	for _, code := range codes {
		codesStr = codesStr + code2SinaCode(code) + ","
		idx++
		if idx%SINA_F_FETCH_SIZE == 0 {
			quotes, err := getQuotes(fmt.Sprintf(SINA_F_URL, codesStr))
			if err != nil {
				return qMap, err
			}
			for c, quote := range quotes {
				qMap[c] = quote
			}
			idx = 0
			codesStr = ""
		}
	}
	if idx > 0 {
		quotes, err := getQuotes(fmt.Sprintf(SINA_F_URL, codesStr))
		if err != nil {
			return qMap, err
		}
		for c, quote := range quotes {
			qMap[c] = quote
		}
	}
	return qMap, nil
}

func getQuotes(url string) (quotes map[string]Quote, err error) {
	quotes = make(map[string]Quote)
	var body []byte
	if body, err = http.Url2Bytes(url); err != nil {
		return
	}
	if body, err = encoding.GbkToUtf8(body); err != nil {
		return
	}
	r := bufio.NewReader(bytes.NewReader(body))
	for {
		var line []byte
		if line, _, err = r.ReadLine(); err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return
		}
		var q Quote
		if q, err = parseQuoteLine(string(line)); err != nil {
			return
		}
		quotes[q.Code] = q
	}
	return
}

func parseQuoteLine(line string) (Quote, error) {
	if !strings.HasPrefix(line, `var hq_str_`) {
		return EMPTY_QUOTE, fmt.Errorf("error sina quote line:%s", line)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error sina quote line:%s\n", line)
			panic(err)
		}
	}()
	var q Quote
	q.Code = sinaCode2Code(line[11:19])
	fields := strings.Split(line, ",")
	if len(fields) < 20 {
		return q, nil
	}
	q.Date = time.Now().Format("2006-01-02")
	q.Name = fields[0][strings.Index(fields[0], "\"")+1:]
	q.Open, _ = strconv.ParseFloat(fields[1], 64)
	q.Close, _ = strconv.ParseFloat(fields[2], 64)
	q.Cur, _ = strconv.ParseFloat(fields[3], 64)
	q.High, _ = strconv.ParseFloat(fields[4], 64)
	q.Low, _ = strconv.ParseFloat(fields[5], 64)
	q.Volume, _ = strconv.Atoi(fields[8])
	q.BuyVol1, _ = strconv.Atoi(fields[10])
	q.BuyP1, _ = strconv.ParseFloat(fields[11], 64)
	q.BuyVol2, _ = strconv.Atoi(fields[12])
	q.BuyP2, _ = strconv.ParseFloat(fields[13], 64)
	q.BuyVol3, _ = strconv.Atoi(fields[14])
	q.BuyP3, _ = strconv.ParseFloat(fields[15], 64)
	q.BuyVol4, _ = strconv.Atoi(fields[16])
	q.BuyP4, _ = strconv.ParseFloat(fields[17], 64)
	q.BuyVol5, _ = strconv.Atoi(fields[18])
	q.BuyP5, _ = strconv.ParseFloat(fields[19], 64)
	q.SaleVol1, _ = strconv.Atoi(fields[20])
	q.SaleP1, _ = strconv.ParseFloat(fields[21], 64)
	q.SaleVol2, _ = strconv.Atoi(fields[22])
	q.SaleP2, _ = strconv.ParseFloat(fields[23], 64)
	q.SaleVol3, _ = strconv.Atoi(fields[24])
	q.SaleP3, _ = strconv.ParseFloat(fields[25], 64)
	q.SaleVol4, _ = strconv.Atoi(fields[26])
	q.SaleP4, _ = strconv.ParseFloat(fields[27], 64)
	q.SaleVol5, _ = strconv.Atoi(fields[28])
	q.SaleP5, _ = strconv.ParseFloat(fields[29], 64)
	return q, nil
}
