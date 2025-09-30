package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

// Has check the form has the specific field value or not
func (form *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}

// Valid returns the validation state of the form
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

// Required will check empty value valition for multiple fields
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, fmt.Sprintf("The %s cannot be blank.\n", field))
		}
	}
}

func (form *Form) MinLength(field string, length int, r *http.Request) bool {
	value := r.Form.Get(field)
	if len(value) < length {
		form.Errors.Add(field, fmt.Sprintf("This %s must be at least %d characters long.", field, length))
		return false
	}
	return true
}

func (form *Form) IsEmail(field string) {
	if !govalidator.IsEmail(form.Get(field)) {
		form.Errors.Add(field, "Invalid email address.")
	}
}
