package goaterr

import (
	"strconv"
	"strings"
)

func printJSON(err error) (result string) {
	return "{" + printJSONBody(err, "\n  ") + "\n}"
}

func printJSONBody(err error, separator string) (result string) {
	var fields []string
	fields = append(fields, `"message": `+strconv.Quote(err.Error()))
	if trackedError, ok := err.(TrackedError); ok {
		fields = append(fields, `"stack": `+strconv.Quote(trackedError.Stack()))
	}
	if errorsWrapper, ok := err.(ErrorsWrapper); ok {
		errs := errorsWrapper.UnwrapAll()
		if len(errs) != 0 {
			var steps []string
			for _, childErr := range errorsWrapper.UnwrapAll() {
				steps = append(steps, printJSONBody(childErr, separator+"  "))
			}
			fields = append(fields, `"wraps": [{`+strings.Join(steps, separator+"}, {")+separator+`}]`)
		}
	} else {
		if errorWrapper, ok := err.(ErrorWrapper); ok {
			childErr := errorWrapper.Unwrap()
			if childErr != nil {
				fields = append(fields, `"wraps": [{`+printJSONBody(childErr, separator+"}, {")+separator+`}]`)
			}
		}
	}
	return separator + strings.Join(fields, ","+separator)
}
