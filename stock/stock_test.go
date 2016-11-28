package stock

import "testing"

func TestSinaService(t *testing.T) {
	s := NewSinaService()
	q, err := s.GetQuote("000001")
	if err != nil {
		t.Errorf("failed:%s", err)
	}
	t.Logf("quote:%#v\n", q)
	qs, err := s.GetQuotes([]string{"000001", "005241"})
	if err != nil {
		t.Errorf("failed:%s", err)
	}
	t.Logf("quotes:%#v\n", qs)
}
