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

func printError(err error, max int) string {
	return strings.Join(unwrapErrorMessages(err, " *", max), "\n")
}

func unwrapErrorMessages(err error, spaces string, max int) (rows []string) {
	if max <= 0 {
		return
	}
	if errorWrapper, ok := err.(MessageError); ok {
		rows = []string{
			spaces + errorWrapper.ErrorMessage(),
		}
	} else {
		rows = []string{
			spaces + err.Error(),
		}
	}
	if errorsWrapper, ok := err.(ErrorsWrapper); ok {
		for _, child := range errorsWrapper.UnwrapAll() {
			rows = append(rows, unwrapErrorMessages(child, "  "+spaces, max-1)...)
		}
	} else {
		if errorWrapper, ok := err.(ErrorWrapper); ok {
			child := errorWrapper.Unwrap()
			if child != nil {
				rows = append(rows, unwrapErrorMessages(child, "  "+spaces, max-1)...)
			}
		}
	}
	return
}
