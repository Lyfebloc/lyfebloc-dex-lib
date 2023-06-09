package entity

type PriceSource string

const (
	PriceSourcelyfebloc  PriceSource = "lyfebloc"
	PriceSourceCoingecko PriceSource = "coingecko"
)

type Price struct {
	Address           string      `json:"address"`
	Price             float64     `json:"price"`
	Liquidity         float64     `json:"liquidity"`
	LpAddress         string      `json:"lpAddress"`
	MarketPrice       float64     `json:"marketPrice"`
	PreferPriceSource PriceSource `json:"preferPriceSource"`
}

// GetPreferredPrice returns the preferred price + if the value is market price or not
// Default is price from Coingecko price source
func (p Price) GetPreferredPrice() (float64, bool) {
	// We don't always have market price, so it's better to have this fallback
	if p.MarketPrice == 0 {
		return p.Price, false
	}

	switch p.PreferPriceSource {
	case PriceSourcelyfebloc:
		return p.Price, false
	case PriceSourceCoingecko:
		return p.MarketPrice, true
	default:
		return p.MarketPrice, true
	}
}
