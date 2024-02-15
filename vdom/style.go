package vdom

// StylePair represent name value pair
type StylePair struct {
	Name  string
	Value string
}

type StylePairs []StylePair

type StyleItem struct {
	Name  string
	Value string

	Delayed StylePairs
	Remove  StylePairs
}

type VNodeStyle []StyleItem
