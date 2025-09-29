package forms

import (
	"net/http"
	"net/url"
)

// Form is custom data types with url values and errors
type Form struct {
	url.Values
	Errors errors
}

// New initialize new form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (form *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}
