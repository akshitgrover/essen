package essen

//EssenError type which interfcaes with native Error type
type EssenError struct {
	errortype string
	message   string
	nilval    bool
}

//Error function to return string representation of error occured.
func (err EssenError) Error() string {
	if err.IsNil() {
		return ""
	}
	return err.errortype + ": " + err.message
}

//Message function to return unexported message field of EssenError
func (err EssenError) Message() string {
	return err.message
}

//Type function to return unexported errortype field of EssenError
func (err EssenError) Type() string {
	return err.errortype
}

//IsNil function to check if EssenError instance is nil or not.
//
//Used to check if error occured or not.
func (err EssenError) IsNil() bool {
	return err.nilval
}
