package validator

func NewValidator() *Validator {
	return &Validator{}
}

type Validator struct {
	errors Errors
}

func (p *Validator) Error() error {
	if p.errors == nil {
		return nil
	}
	return p.errors
}

func (p *Validator) AddError(field string, msg string) {
	if p.errors == nil {
		p.errors = NewErrors()
	}
	p.errors.Add(field, msg)
}

func (p *Validator) getError(field string) string {
	if p.errors == nil {
		return ""
	}

	return p.errors.Get(field)
}

func (p *Validator) CheckError(field string) bool {
	if p.getError(field) != "" {
		return true
	}
	return false
}

func (p *Validator) HasError() bool {
	return p.errors != nil
}

func (p *Validator) Validate(field, label string, value ...interface{}) *Flow {
	flow := new(Flow)
	flow.label = label
	flow.field = field
	flow.values = value
	flow.validator = p
	return flow
}
