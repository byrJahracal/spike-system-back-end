package errorHandler

import "log"

func ErrorHandler(err error, failMsg string, successMsg string) {
	if err != nil {
		log.Fatalf("%s: %s", failMsg, err)
	} else if successMsg != "" {
		log.Println(successMsg)
	}
}

func ErrorWarning(err error, failMsg string, successMsg string) bool {
	if err != nil {
		log.Panic("%s: %s", failMsg, err)
		return true
	} else if successMsg != "" {
		log.Println(successMsg)
	}
	return false
}
