// Adapted copy from the OTLP text in the Opentelemetry collector

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textmarshaler

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// MarshalLogs plog.Logs to OTLP text.
func MarshalLogs(ld plog.Logs) ([]byte, error) {
	buf := dataBuffer{}
	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		buf.logEntry("ResourceLog #%d", i)
		rl := rls.At(i)
		buf.logEntry("Resource SchemaURL: %s", rl.SchemaUrl())
		buf.logAttributes("Resource attributes", rl.Resource().Attributes())
		ills := rl.ScopeLogs()
		for j := 0; j < ills.Len(); j++ {
			buf.logEntry("ScopeLogs #%d", j)
			ils := ills.At(j)
			buf.logEntry("ScopeLogs SchemaURL: %s", ils.SchemaUrl())
			buf.logInstrumentationScope(ils.Scope())

			logs := ils.LogRecords()
			for k := 0; k < logs.Len(); k++ {
				buf.logEntry("LogRecord #%d", k)
				lr := logs.At(k)
				buf.logEntry("ObservedTimestamp: %s", lr.ObservedTimestamp())
				buf.logEntry("Timestamp: %s", lr.Timestamp())
				buf.logEntry("SeverityText: %s", lr.SeverityText())
				buf.logEntry("SeverityNumber: %s(%d)", lr.SeverityNumber(), lr.SeverityNumber())
				buf.logEntry("Body: %s", valueToString(lr.Body()))
				buf.logAttributes("Attributes", lr.Attributes())
				buf.logEntry("Trace ID: %s", lr.TraceID())
				buf.logEntry("Span ID: %s", lr.SpanID())
				buf.logEntry("Flags: %d", lr.Flags())
			}
		}
	}

	return buf.buf.Bytes(), nil
}
