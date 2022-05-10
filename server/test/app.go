package test

import (
	"github.com/kubeshop/tracetest/app"
	"github.com/kubeshop/tracetest/config"
	"github.com/kubeshop/tracetest/tracedb"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configtls"
)

func GetTestingApp(demoApp *DemoApp) (*app.App, error) {
	db, err := GetTestingDatabase("file://../migrations")

	if err != nil {
		return nil, err
	}

	config := config.Config{
		JaegerConnectionConfig: &configgrpc.GRPCClientSettings{
			Endpoint: demoApp.JaegerEndpoint(),
			TLSSetting: configtls.TLSClientSetting{
				Insecure: true,
			},
		},
	}

	tracedb, err := tracedb.New(config)
	if err != nil {
		return nil, err
	}

	return app.New(config, db, tracedb)
}
