package observability

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func SetUpObservability(serviceName string, serviceVersion string, opts ...TracingOption) (func(context.Context) error, error) {
	ctx := context.Background()
	cfg := newTracingConfig(opts...)

	res, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(serviceVersion),
	), resource.WithContainer())
	if err != nil {
		return nil, fmt.Errorf("failed to create resources: %w", err)
	}

	traceExporter, err := cfg.exporterFactory(ctx, cfg.exporterOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

type NewTraceExporter func(context.Context, ...ExporterOption) (sdktrace.SpanExporter, error)

func newGRpcOtelTraceExporter(ctx context.Context, opts ...ExporterOption) (sdktrace.SpanExporter, error) {
	cfg := newExporterConfig(opts...)
	ctx, cancel := context.WithTimeout(ctx, cfg.timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx, fmt.Sprintf("%s:%d", cfg.host, cfg.grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	return traceExporter, nil
}

func newHttpOtelTraceExporter(ctx context.Context, opts ...ExporterOption) (sdktrace.SpanExporter, error) {
	cfg := newExporterConfig(opts...)
	ctx, cancel := context.WithTimeout(ctx, cfg.timeout)
	defer cancel()

	traceExporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint(fmt.Sprintf("%s:%d", cfg.host, cfg.httpPort)), otlptracehttp.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	return traceExporter, nil
}
