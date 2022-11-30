package tracedb

import (
	"context"
	"errors"
	"fmt"

	"github.com/kubeshop/tracetest/server/config"
	"github.com/kubeshop/tracetest/server/model"
	"github.com/kubeshop/tracetest/server/traces"
)

var ErrTraceNotFound = errors.New("trace not found")

const (
	JAEGER_BACKEND     string = "jaeger"
	TEMPO_BACKEND      string = "tempo"
	OPENSEARCH_BACKEND string = "opensearch"
	SIGNALFX           string = "signalfx"
	OTLP               string = "otlp"
)

type TraceDB interface {
	GetTraceByID(ctx context.Context, traceID string) (traces.Trace, error)
	Close() error
}

type noopTraceDB struct{}

func (db *noopTraceDB) GetTraceByID(ctx context.Context, traceID string) (t traces.Trace, err error) {
	return
}
func (db *noopTraceDB) Close() error { return nil }

var ErrInvalidTraceDBProvider = fmt.Errorf("invalid traceDB provider: available options are (jaeger, tempo)")

func New(c config.Config, repository model.RunRepository) (db TraceDB, err error) {
	selectedDataStore, err := c.DataStore()
	if err != nil {
		return &noopTraceDB{}, nil
	}

	err = ErrInvalidTraceDBProvider

	switch {
	case selectedDataStore.Type == JAEGER_BACKEND:
		db, err = newJaegerDB(&selectedDataStore.Jaeger)
	case selectedDataStore.Type == TEMPO_BACKEND:
		db, err = newTempoDB(&selectedDataStore.Tempo)
	case selectedDataStore.Type == OPENSEARCH_BACKEND:
		db, err = newOpenSearchDB(selectedDataStore.OpenSearch)
	case selectedDataStore.Type == SIGNALFX:
		db, err = newSignalFXDB(selectedDataStore.SignalFX)
	case selectedDataStore.Type == OTLP:
		db, err = newCollectorDB(repository)
	}

	return
}
