package stock

//Quote stock quote
type Quote struct {
	Code     string
	Name     string
	Date     string
	Volume   int
	Open     float64
	Close    float64
	Cur      float64
	High     float64
	Low      float64
	BuyVol1  int
	BuyP1    float64
	BuyVol2  int
	BuyP2    float64
	BuyVol3  int
	BuyP3    float64
	BuyVol4  int
	BuyP4    float64
	BuyVol5  int
	BuyP5    float64
	SaleVol1 int
	SaleP1   float64
	SaleVol2 int
	SaleP2   float64
	SaleVol3 int
	SaleP3   float64
	SaleVol4 int
	SaleP4   float64
	SaleVol5 int
	SaleP5   float64
}

var emptyQuote = Quote{}

//Service stock quote service
type Service interface {
	GetQuote(code string) (Quote, error)
	GetQuotes(codes []string) (map[string]Quote, error)
}
