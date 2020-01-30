package goaterr

import "strconv"

func print(err error) (result string) {
	return printStep(err, "  \n")
}

func printStep(err error, separator string) (result string) {
	result = `{`
	if messageError, ok := err.(MessageError); ok {
		result += separator + `message: ` + strconv.Quote(messageError.Message())
	} else {
		result += separator + `message: ` + strconv.Quote(err.Error())
	}
	if trackedError, ok := err.(TrackedError); ok {
		result += separator + `stack: ` + strconv.Quote(trackedError.Stack())
	}
	if errorsWrapper, ok := err.(ErrorsWrapper); ok {
		errs := errorsWrapper.UnwrapAll()
		if len(errs) != 0 {
			result += separator + `wraps: [`
			for _, childErr := range errorsWrapper.UnwrapAll() {
				result += printStep(childErr, separator+"  ")
			}
			result += `]`
		}
	} else {
		if errorWrapper, ok := err.(ErrorWrapper); ok {
			childErr := errorWrapper.Unwrap()
			if childErr != nil {
				result += separator + `wraps: [` + printStep(childErr, separator+"  ") + `]`
			}
		}
	}
	return result + `}`
}
