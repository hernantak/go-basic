package entity

type Product struct {
	ID    int
	Nama  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	var Status string
	if p.Stock < 3 {
		Status = "Stock hampir habis"
	} else if p.Stock < 10 {
		Status = "Stock terbatas"
	}

	return Status
}
