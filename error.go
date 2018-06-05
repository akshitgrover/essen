package essen

type EssenError struct {
	errortype string
	message   string
	nilval    bool
}

func (err EssenError) Error() string {
	if err.IsNil() {
		return ""
	}
	return err.errortype + ": " + err.message
}

func (err EssenError) Message() string {
	return err.message
}

func (err EssenError) Type() string {
	return err.errortype
}

func (err EssenError) IsNil() bool {
	return err.nilval
}
