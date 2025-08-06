package html

type ArticleElement struct {
	BaseElement
}

func Article() *ArticleElement {
	return &ArticleElement{InitBaseElement("article")}
}
