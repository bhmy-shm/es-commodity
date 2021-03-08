package Model

const (
	OrderByPriceASC = 1 //价格从低到高
	OrderByPriceDESC = 2 //价格从高到低
)
type SearchModel struct {
	BookName string  `json:"book_name"`
	BookPress  string `json:"book_press"`
	//RangeQuery起始金额，结束金额
	BookPrice1Start float64 `json:"book_price1_start" binding:"required,gte=0,lt=10000"`
	BookPrice2End	float64 `json:"book_price1_end" binding:"required,gte=0,lt=10000,gtefield=BookPrice1Start`
	//排序
	OrderSet struct{
		Score bool `json:"score"`
		PriceOrder int `json:"price_order" binding:"oneof=0 1 2"`
	}`json:"OrderSet" binding:"required,dive"`
	//分页
	Current int `json:"current" binding:"gte=1"`
	Size int `json:"size" binding:"oneof=10 20 50"`
}

func NewSearchModel() *SearchModel{
	return &SearchModel{}
}
