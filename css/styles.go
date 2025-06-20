package css

func Color(v string) Style {
	return Style{Name: "color", Value: v}
}

func BackgroundColor(v string) Style {
	return Style{Name: "background-color", Value: v}
}
