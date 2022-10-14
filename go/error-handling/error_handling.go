package erratum

type ErrorHandler struct {
	Error error
}

func (e ErrorHandler) isTransient() (TransientError, bool) {
	var i interface{} = e.Error
	var err, transient_errors = i.(TransientError)
	return err, transient_errors
}

func Use(opener ResourceOpener, input string) (ret error) {
	a := ErrorHandler{}
	res, err := opener()
	a.Error = err

	defer func() {
		if r := recover(); r != nil {
			fe, isFrobError := r.(FrobError)
			ret = fe // update return value with FrobError
			if isFrobError {
				res.Defrob(fe.defrobTag)
			} else {
				nfe, _ := r.(error)
				ret = nfe // update return value with NonFrobError
			}
			res.Close()
		}
	}()

	for _, isTransient := a.isTransient(); isTransient; {
		res, err = opener()
		a.Error = err
		if err == nil {
			break
		}
		_, isTransient = a.isTransient()
	}

	if a.Error == nil {
		res.Frob(input)
		res.Close()
	}

	return a.Error
}
