package response

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewResponse(response map[string]string, elapsed time.Time) logrus.Fields {
	// t := Todo{Result: response, Completed: true, Due: time.Now()}

	start := time.Now()

	fields := logrus.Fields{
		"Body":      response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}

func NewResponseIndex(response map[int][]string, elapsed time.Time) logrus.Fields {
	// t := Todo{Result: response, Completed: true, Due: time.Now()}

	start := time.Now()

	fields := logrus.Fields{
		"Response":  response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}
