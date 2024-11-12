package main

import (
	"context"

	"github.com/goplugin/plugin-libocr/offchainreporting2plus/ocr3types"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/loop"
	"github.com/goplugin/plugin-common/pkg/services"
	"github.com/goplugin/plugin-common/pkg/types"
	v1 "github.com/goplugin/plugin-common/pkg/types/mercury/v1"
	v2 "github.com/goplugin/plugin-common/pkg/types/mercury/v2"
	v3 "github.com/goplugin/plugin-common/pkg/types/mercury/v3"
	v4 "github.com/goplugin/plugin-common/pkg/types/mercury/v4"

	ds_v1 "github.com/goplugin/plugin-data-streams/mercury/v1"
	ds_v2 "github.com/goplugin/plugin-data-streams/mercury/v2"
	ds_v3 "github.com/goplugin/plugin-data-streams/mercury/v3"
	ds_v4 "github.com/goplugin/plugin-data-streams/mercury/v4"
)

type Plugin struct {
	loop.Plugin
	stop services.StopChan
}

func NewPlugin(lggr logger.Logger) *Plugin {
	return &Plugin{Plugin: loop.Plugin{Logger: lggr}, stop: make(services.StopChan)}
}

func (p *Plugin) NewMercuryV1Factory(ctx context.Context, provider types.MercuryProvider, dataSource v1.DataSource) (types.MercuryPluginFactory, error) {
	var ctxVals loop.ContextValues
	ctxVals.SetValues(ctx)
	lggr := logger.With(p.Logger, ctxVals.Args()...)

	factory := ds_v1.NewFactory(dataSource, lggr, provider.OnchainConfigCodec(), provider.ReportCodecV1())

	s := &mercuryPluginFactoryService{lggr: logger.Named(lggr, "MercuryV1PluginFactory"), MercuryPluginFactory: factory}

	p.SubService(s)

	return s, nil
}

func (p *Plugin) NewMercuryV2Factory(ctx context.Context, provider types.MercuryProvider, dataSource v2.DataSource) (types.MercuryPluginFactory, error) {
	var ctxVals loop.ContextValues
	ctxVals.SetValues(ctx)
	lggr := logger.With(p.Logger, ctxVals.Args()...)

	factory := ds_v2.NewFactory(dataSource, lggr, provider.OnchainConfigCodec(), provider.ReportCodecV2())

	s := &mercuryPluginFactoryService{lggr: logger.Named(lggr, "MercuryV2PluginFactory"), MercuryPluginFactory: factory}

	p.SubService(s)

	return s, nil
}

func (p *Plugin) NewMercuryV3Factory(ctx context.Context, provider types.MercuryProvider, dataSource v3.DataSource) (types.MercuryPluginFactory, error) {
	var ctxVals loop.ContextValues
	ctxVals.SetValues(ctx)
	lggr := logger.With(p.Logger, ctxVals.Args()...)

	factory := ds_v3.NewFactory(dataSource, lggr, provider.OnchainConfigCodec(), provider.ReportCodecV3())

	s := &mercuryPluginFactoryService{lggr: logger.Named(lggr, "MercuryV3PluginFactory"), MercuryPluginFactory: factory}

	p.SubService(s)

	return s, nil
}

func (p *Plugin) NewMercuryV4Factory(ctx context.Context, provider types.MercuryProvider, dataSource v4.DataSource) (types.MercuryPluginFactory, error) {
	var ctxVals loop.ContextValues
	ctxVals.SetValues(ctx)
	lggr := logger.With(p.Logger, ctxVals.Args()...)

	factory := ds_v4.NewFactory(dataSource, lggr, provider.OnchainConfigCodec(), provider.ReportCodecV4())

	s := &mercuryPluginFactoryService{lggr: logger.Named(lggr, "MercuryV4PluginFactory"), MercuryPluginFactory: factory}

	p.SubService(s)

	return s, nil
}

type mercuryPluginFactoryService struct {
	services.StateMachine
	lggr logger.Logger
	ocr3types.MercuryPluginFactory
}

func (r *mercuryPluginFactoryService) Name() string { return r.lggr.Name() }

func (r *mercuryPluginFactoryService) Start(ctx context.Context) error {
	return r.StartOnce("ReportingPluginFactory", func() error { return nil })
}

func (r *mercuryPluginFactoryService) Close() error {
	return r.StopOnce("ReportingPluginFactory", func() error { return nil })
}

func (r *mercuryPluginFactoryService) HealthReport() map[string]error {
	return map[string]error{r.Name(): r.Healthy()}
}
