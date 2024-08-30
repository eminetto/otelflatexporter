package otelflatexporter

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type otelflatexporter struct {
}

func NewOtelFlatexporter() *otelflatexporter {
	return &otelflatexporter{}
}

func (s *otelflatexporter) pushLogs(_ context.Context, ld plog.Logs) error {
	return nil
}

func (s *otelflatexporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	return nil
}

func (s *otelflatexporter) pushTraces(_ context.Context, td ptrace.Traces) error {
	return nil
}
