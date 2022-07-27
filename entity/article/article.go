package article

func NewRequest(article *Article) *Request {
	return &Request{Article: *article}
}
