package otelflatexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var (
	Type = component.MustNewType("otelflatexporter")
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		Type,
		createDefaultConfig,
		// exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
		// exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
	)
}

// func createTracesExporter(
// 	ctx context.Context,
// 	set exporter.Settings,
// 	config component.Config) (exporter.Traces, error) {

// 	cfg := config.(*Config)
// 	s := NewOtelFlatexporter()
// 	return exporterhelper.NewTracesExporter(ctx, set, cfg, s.pushTraces)
// }

// func createMetricsExporter(
// 	ctx context.Context,
// 	set exporter.Settings,
// 	config component.Config) (exporter.Metrics, error) {

// 	cfg := config.(*Config)
// 	s := NewOtelFlatexporter()
// 	return exporterhelper.NewMetricsExporter(ctx, set, cfg, s.pushMetrics)
// }

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	config component.Config) (exporter.Logs, error) {

	cfg := config.(*Config)
	s := NewOtelFlatexporter()
	return exporterhelper.NewLogsExporter(ctx, set, cfg, s.pushLogs)
}

type Config struct {
}

func createDefaultConfig() component.Config {
	return &Config{}
}
