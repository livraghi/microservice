#!/usr/bin/env bash

SEP="------------------------------------------------------------------"
echo -e "$SEP"
ENGINE=${ENGINE:-docker}
OTEL_VERSION=${OTEL_VERSION:-latest}
OTEL_IMAGE=otel/opentelemetry-collector-contrib:$OTEL_VERSION
echo -e "Using container engine: \t\t\t\t$ENGINE"
echo -e "OTLP image version: \t\t\t\t\t$OTEL_VERSION"
echo -e "$SEP"

PPROF_PORT=${PPROF_PORT:-1888} # pprof extension
PROMETHEUS_COLLECTOR_PORT=${PROMETHEUS_COLLECTOR_PORT:-8888} # Prometheus metrics exposed by the Collector
PROMETHEUS_EXPORTER_PORT=${PROMETHEUS_EXPORTER_PORT:-8889} # Prometheus exporter metrics
HEALTH_CHECK_PORT=${HEALTH_CHECK_PORT:-13133} # health_check extension
GRPC_PORT=${GRPC_PORT:-4317} # OTLP gRPC receiver
HTTP_PORT=${HTTP_PORT:-4318} # OTLP http receiver
ZPAGES_PORT=${ZPAGES_PORT:-55679} # zpages extension

echo -e "pprof extension port:\t\t\t\t\t$PPROF_PORT"
echo -e "Prometheus metrics exposed by the Collector port: \t$PROMETHEUS_COLLECTOR_PORT"
echo -e "Prometheus exporter metrics port: \t\t\t$PROMETHEUS_EXPORTER_PORT"
echo -e "health_check extension port:\t\t\t\t$HEALTH_CHECK_PORT"
echo -e "OTLP gRPC receiver port:\t\t\t\t$GRPC_PORT"
echo -e "OTLP http receiver port:\t\t\t\t$HTTP_PORT"
echo -e "zpages extension port:\t\t\t\t\t$ZPAGES_PORT"
echo -e "$SEP"

if [ "$ENGINE" == "podman" ]; then
  $ENGINE pull docker://$OTEL_IMAGE
  echo -e "$SEP"
fi
echo -e "Starting OTLP Collector..."
$ENGINE run --rm \
  -p $PPROF_PORT:1888 \
  -p $PROMETHEUS_COLLECTOR_PORT:8888 \
  -p $PROMETHEUS_EXPORTER_PORT:8889 \
  -p $HEALTH_CHECK_PORT:13133 \
  -p $GRPC_PORT:4317 \
  -p $HTTP_PORT:4318 \
  -p $ZPAGES_PORT:55679 \
  otel/opentelemetry-collector-contrib