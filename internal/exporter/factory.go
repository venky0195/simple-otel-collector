package exporter

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	typeStr = "exporter"
)

func NewFactory(logExporter *LogExporter) exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(func(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (exporter.Logs, error) {
			return createLogsExporter(ctx, set, cfg, logExporter)
		}, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
	logExporter *LogExporter,
) (exporter.Logs, error) {

	return exporterhelper.NewLogsExporter(ctx, set, cfg,
		func(ctx context.Context, ld plog.Logs) error {
			return logExporter.pushLogs(ctx, ld)
		},
	)
}
