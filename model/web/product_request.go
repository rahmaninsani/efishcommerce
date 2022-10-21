package web

type ProductFilterRequest struct {
	Categories []string
	MinPrice   uint64
	MaxPrice   uint64
}
