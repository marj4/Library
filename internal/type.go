package internal

type Book struct {
	ID           int    `json:"id"`
	BookName     string `json:"name"`
	Author       string `json:"author"`
	YearRealease string `json:"year_realease"`
}
