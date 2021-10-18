package response

import (
	"time"

	"github.com/sirupsen/logrus"
)

func ErrorResponse(e int, m string, elapsed time.Time) logrus.Fields {

	start := time.Now()

	fields := logrus.Fields{
		"Status":    e,
		"Completed": false,
		"Message":   m,
		"Time":      start.Sub(elapsed).Milliseconds(),
	}

	return fields
}
