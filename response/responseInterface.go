package response

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewResponseInterface(response map[string]interface{}, elapsed time.Time) logrus.Fields {
	// t := Todo{Result: response, Completed: true, Due: time.Now()}

	start := time.Now()

	fields := logrus.Fields{
		"Body":      response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}
