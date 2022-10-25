package gui

type SelectOptions []Options

type Options struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}
