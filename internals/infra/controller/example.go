package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Isameer3056/Observability-go-project/internals/domain/core/domain"
	"github.com/Isameer3056/Observability-go-project/pkg/adapter/instrumentation"
	"github.com/Isameer3056/Observability-go-project/pkg/adapter/metric"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type exampleController struct {
	exampleUseCase domain.ExampleUseCase
}

// GetExample implements domain.ExampleController.
func (controller *exampleController) GetExample(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	propagator := otel.GetTextMapPropagator()
	ctx = propagator.Extract(ctx, propagation.HeaderCarrier(request.Header))

	ctx, span := instrumentation.Tracer.Start(ctx, "controller.GetExample")
	defer span.End()

	start := time.Now()
	metric.RequestInFlight.Add(ctx, 1)
	metric.RequestCounter.Add(ctx, 1)
	defer func() {
		metric.RequestInFlight.Add(ctx, -1)
		metric.RequestDuration.Record(ctx, time.Since(start).Seconds())
	}()

	instrumentation.Logger.InfoContext(ctx, "Processing request",
		"method", request.Method,
		"path", request.URL.Path,
	)

	example, err := controller.exampleUseCase.GetExample(ctx)
	if err != nil {
		instrumentation.Logger.ErrorContext(ctx, "Failed to get example",
			"error", err,
			"method", request.Method,
			"path", request.URL.Path,
			"status", http.StatusInternalServerError,
		)
		span.RecordError(err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	metric.LetterCounter.Add(ctx, 1)
	metric.WordLengthGauge.Record(ctx, int64(len(example.Word)))

	response.WriteHeader(http.StatusOK)
	err = json.NewEncoder(response).Encode(example)
	if err != nil {
		instrumentation.Logger.ErrorContext(ctx, "Failed to encode response",
			"error", err,
			"method", request.Method,
			"path", request.URL.Path,
			"status", http.StatusInternalServerError,
		)
		span.RecordError(err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
}

func NewExampleController(exampleUseCase domain.ExampleUseCase) domain.ExampleController {
	return &exampleController{
		exampleUseCase: exampleUseCase,
	}
}
