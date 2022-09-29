package grid_scui

type Form struct {
	*View

	LabelWidth    string      `json:"labelWidth"`
	LabelPosition string      `json:"labelPosition"`
	Size          string      `json:"size"`
	FormItems     []FormItems `json:"formItems"`
}

type FormItems struct {
	Label          string                 `json:"label,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Value          string                 `json:"value,omitempty"`
	Component      string                 `json:"component,omitempty"`
	Span           int                    `json:"span"`
	Options        map[string]interface{} `json:"options,omitempty"`
	Rules          map[string]string      `json:"rules,omitempty"`
	RequiredHandle string                 `json:"required_handle,omitempty"`
	HideHandle     string                 `json:"hide_handle,omitempty"`
}

func NewForm() *Form {
	view := NewView("form.vue")
	return &Form{
		View:          view,
		LabelWidth:    "130px",
		LabelPosition: "right",
	}
}
