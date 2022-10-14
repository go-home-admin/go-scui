package gui

type Controller interface {
	Grid(v View)
	Form(f Form)
}

type View interface {
}

type Form interface {
}

type Style interface {
	Width(v string) Style
	Height(v string) Style
	MinWidth(v string) Style
	MaxWidth(v string) Style
	LineWidth(v string) Style
	MinHeight(v string) Style
	MaxHeight(v string) Style
	LineHeight(v string) Style
}

type Element interface {
	Style
}
