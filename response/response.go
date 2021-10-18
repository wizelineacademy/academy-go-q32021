package response

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewResponse(response map[string]string, elapsed time.Time) logrus.Fields {

	start := time.Now()

	fields := logrus.Fields{
		"Body":      response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}

func NewResponseIndex(response map[int][]string, elapsed time.Time) logrus.Fields {

	start := time.Now()

	fields := logrus.Fields{
		"Response":  response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}

func NewResponseInterface(response map[string]interface{}, elapsed time.Time) logrus.Fields {

	start := time.Now()

	fields := logrus.Fields{
		"Body":      response,
		"Completed": true,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}
