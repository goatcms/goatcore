package goaterr

import (
	"strconv"
	"strings"
)

func print(err error) (result string) {
	return "{" + printBody(err, "\n  ") + "\n}"
}

func printBody(err error, separator string) (result string) {
	var fields []string
	if messageError, ok := err.(MessageError); ok {
		fields = append(fields, `"message": `+strconv.Quote(messageError.Message()))
		if trackedError, ok := err.(TrackedError); ok {
			fields = append(fields, `"stack": `+strconv.Quote(trackedError.Stack()))
		}
		if errorsWrapper, ok := err.(ErrorsWrapper); ok {
			errs := errorsWrapper.UnwrapAll()
			if len(errs) != 0 {
				var steps []string
				for _, childErr := range errorsWrapper.UnwrapAll() {
					steps = append(steps, printBody(childErr, separator+"  "))
				}
				fields = append(fields, `"wraps": [{`+strings.Join(steps, separator+"}, {")+separator+`}]`)
			}
		} else {
			if errorWrapper, ok := err.(ErrorWrapper); ok {
				childErr := errorWrapper.Unwrap()
				if childErr != nil {
					fields = append(fields, `"wraps": [{`+printBody(childErr, separator+"}, {")+separator+`}]`)
				}
			}
		}
	} else {
		fields = append(fields, `"message": `+strconv.Quote(err.Error()))
	}
	return separator + strings.Join(fields, ","+separator)
}
