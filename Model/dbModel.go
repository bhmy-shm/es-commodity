package Model

type Books struct{
	BookID int `gorm:"column:book_id"`
	BookName string `gorm:"column:book_name"`
	BookIntr string `gorm:"column:book_intr"`
	BookPrice1 float64 `gorm:"column:book_price1"`
	BookPrice2 float64 `gorm:"column:book_price2"`
	BookAuthor string `gorm:"column:book_author"`
	BookPress string `gorm:"column:book_press"`
	BookDate string `gorm:"column:book_date"`
	BookKind int `gorm:"column:book_kind"`
}

type BookSlice []*Books
