package errb

type ErrBin interface {
	Error() string

	ErrorRU() string
	WithErrorRU(ruStr string) ErrBin

	Code() string
	WithCode(code string) ErrBin

	Details() interface{}
	WithDetails(details interface{}) ErrBin
}

type Err struct {
	str     string
	ruStr   string
	code    string
	details interface{}
}

func (e Err) Error() string {
	return e.str
}

func (e Err) ErrorRU() string {
	if e.ruStr != "" {
		return e.ruStr
	}
	return e.str
}

func (e Err) WithErrorRU(ruStr string) ErrBin {
	return Err{
		str:     e.str,
		ruStr:   ruStr,
		code:    e.code,
		details: e.details,
	}
}

func (e Err) Code() string {
	return e.code
}

func (e Err) WithCode(code string) ErrBin {
	return Err{
		str:     e.str,
		ruStr:   e.ruStr,
		code:    code,
		details: e.details,
	}
}

func (e Err) Details() interface{} {
	return e.details
}

func (e Err) WithDetails(details interface{}) ErrBin {
	return Err{
		str:     e.str,
		ruStr:   e.ruStr,
		code:    e.code,
		details: details,
	}
}

func New(str string) ErrBin {
	return &Err{
		str: str,
	}
}

func Is(e error, code string) bool {
	if eb, ok := e.(ErrBin); ok && eb.Code() == code {
		return true
	}

	return false
}

// TODO: wrap function
