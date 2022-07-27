package article

type IBuilder interface {
	SetTagList(tagList []string) IBuilder
	SetTitle(title string) IBuilder
	SetDescription(description string) IBuilder
	SetBody(body string) IBuilder
	Build() *Article
}

type Builder struct {
	a Article
}

func NewBuilder() IBuilder {
	return &Builder{}
}

func (b Builder) SetTagList(tagList []string) IBuilder {
	b.a.TagList = tagList
	return b
}

func (b Builder) SetTitle(title string) IBuilder {
	b.a.Title = title
	return b
}

func (b Builder) SetDescription(description string) IBuilder {
	b.a.Description = description
	return b
}

func (b Builder) SetBody(body string) IBuilder {
	b.a.Body = body
	return b
}

func (b Builder) Build() *Article {
	return &b.a
}
