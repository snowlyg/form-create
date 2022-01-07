package form_create

import (
	"fmt"
	"strings"
)

// Form
type Form struct {
	Rule    []Rule                   `json:"rule"`
	Action  string                   `json:"action"`
	Method  string                   `json:"method"`
	Title   string                   `json:"title"`
	Config  Config                   `json:"config,omitempty"`
	Headers []map[string]interface{} `json:"headers,omitempty"`
}

// Option
type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Children []Option    `json:"children"`
}

// Config
type Config struct{}

// Rule
type Rule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Col      map[string]interface{}   `json:"col,omitempty"`
	Options  []Option                 `json:"options,omitempty"`
	Controls []Control                `json:"control,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}

// ControlRule
type ControlRule struct {
	Title    string                   `json:"title"`
	Type     string                   `json:"type"`
	Field    string                   `json:"field"`
	Info     string                   `json:"info"`
	Value    interface{}              `json:"value"`
	Props    map[string]interface{}   `json:"props"`
	Options  []Option                 `json:"options,omitempty"`
	Validate []map[string]interface{} `json:"validate,omitempty"`
}

// Control
type Control struct {
	Value interface{} `json:"value"`
	Rule  []Rule      `json:"rule"`
}

// TransData
func (r *Rule) TransData(rule string, token []byte) {
	switch r.Type {
	case "input":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "text",
		}
	case "textarea":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
			"type":        "textarea",
		}
		r.Type = "input"
	case "number":
		r.Props = map[string]interface{}{
			"placeholder": "请输入" + r.Title,
		}
		r.Type = "inputNumber"
	case "radio":
		r.Props = map[string]interface{}{}
		rules := strings.Split(rule, ";")
		for _, ru := range rules {
			rus := strings.Split(ru, ":")
			if len(rus) == 2 {
				r.Options = append(r.Options, Option{Label: rus[1], Value: rus[0]})
			}
		}
	case "file":
		// seitURL, _ := param.GetSeitURL()
		r.Props = map[string]interface{}{
			"action": fmt.Sprintf("%s/v1/admin/media/upload", ""),
			"data":   map[string]interface{}{},
			"headers": map[string]interface{}{
				"Authorization": "Bearer " + string(token),
			},

			"limit":      1,
			"uploadType": "file",
		}
		r.Type = "upload"
	case "image":
		r.Props = map[string]interface{}{
			"footer":    false,
			"height":    "480px",
			"maxLength": 1,
			"modal":     map[string]interface{}{"modal": false},
			"src":       "/admin/setting/uploadPicture?field=" + r.Field + "&type=1",
			"title":     "请选择" + r.Title,
			"type":      r.Type,
			"width":     "896px",
		}
		r.Type = "frame"
	}

}

func (rule *Rule) AddValidator(validator ...map[string]interface{}) *Rule {
	rule.Validate = append(rule.Validate, validator...)
	return rule
}

func (rule *Rule) AddOption(opt ...Option) *Rule {
	rule.Options = append(rule.Options, opt...)
	return rule
}

func (rule *Rule) AddControl(control ...Control) *Rule {
	rule.Controls = append(rule.Controls, control...)
	return rule
}

func (rule *Rule) AddProps(props map[string]interface{}) *Rule {
	rule.Props = props
	return rule
}

// NewRadio
func NewRadio(title, field, info string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "radio",
		Field: field,
		Value: value,
		Info:  info,
	}
}

// NewCascader
func NewCascader(title, field, info string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "cascader",
		Field: field,
		Value: value,
		Info:  info,
	}
}

// NewTextarea
func NewTextarea(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "input",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"type":        "textarea",
			"placeholder": placeholder,
		},
	}
}

// NewInput
func NewInput(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "input",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"type":        "text",
			"placeholder": placeholder,
		},
	}
}

// NewDatePicker
func NewDatePicker(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "input",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"type":        "date",
			"placeholder": placeholder,
			"editable":    false,
		},
	}
}

// NewHidden
func NewHidden(field string, value interface{}) *Rule {
	return &Rule{
		Type:  "hidden",
		Field: field,
		Value: value,
	}
}

// NewInputNumber
func NewInputNumber(title, field, placeholder string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "inputNumber",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"placeholder": placeholder,
		},
	}
}

// NewFrame
func NewFrame(title, field string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "frame",
		Field: field,
		Value: value,
	}
}

// NewRate
func NewRate(title, field string, span int64, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "rate",
		Field: field,
		Value: value,
		Col: map[string]interface{}{
			"span": 8,
		},
		Props: map[string]interface{}{
			"max": 5,
		},
	}
}

// NewSelect
func NewSelect(title, field, placeholder string, value interface{}, multiple bool) *Rule {
	return &Rule{
		Title: title,
		Type:  "select",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"multiple":    multiple,
			"placeholder": placeholder,
		},
	}
}

// NewSwitch
func NewSwitch(title, field string, value interface{}) *Rule {
	return &Rule{
		Title: title,
		Type:  "switch",
		Field: field,
		Value: value,
		Props: map[string]interface{}{
			"activeValue":   1,
			"inactiveValue": 2,
			"inactiveText":  "关闭",
			"activeText":    "开启",
		},
	}
}

func (form *Form) AddRule(rule Rule) *Form {
	form.Rule = append(form.Rule, rule)
	return form
}

func (form *Form) SetAction(uri string) {
	form.Action = SetUrl(uri)
}

func SetUrl(uri string) string {
	// if multi_gin.IsAdmin(ctx) {
	// 	return g.TENANCY_CONFIG.System.AdminPreix + uri
	// } else if multi_gin.IsTenancy(ctx) {
	// 	return g.TENANCY_CONFIG.System.ClientPreix + uri
	// }
	return ""
}
