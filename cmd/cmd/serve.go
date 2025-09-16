/*
Copyright © 2025 Sameer Imtiaz <isameer3056@gmail.com>
*/
package cmd

import (
	"context"
	"time"

	"github.com/Isameer3056/Observability-go-project/pkg/adapter/instrumentation"
	"github.com/Isameer3056/Observability-go-project/pkg/adapter/metric"
	"github.com/Isameer3056/Observability-go-project/pkg/adapter/rest"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start an OpenTelemetry-instrumented HTTP service",
	Long: `Starts an HTTP service that is instrumented with OpenTelemetry for observability.

This service exposes metrics, traces, and logs that can be collected and analyzed
by OpenTelemetry collectors. The service responds with its name and demonstrates
distributed tracing capabilities when called through the Traefik reverse proxy.

Example usage:
  Observability-go-project serve`,
	Run: func(cmd *cobra.Command, args []string) {
		shutdown, err := instrumentation.Initialize(cmd.Context())
		if err != nil {
			panic(err)
		}

		defer func() {
			ctx, cancel := context.WithTimeout(cmd.Context(), time.Second*5)
			defer cancel()
			if err := shutdown(ctx); err != nil {
				panic("failed to shutdown TracerProvider")
			}
		}()

		metric.Initialize(otel.GetMeterProvider().Meter(instrumentation.Name))

		rest.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
