package main

import (
	"flag"
	"fmt"
	"net/http/pprof"
	"os"

	fasthttpprometheus "github.com/flf2ko/fasthttp-prometheus"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/sabrusrin/wb_go_task/pkg/service"
	"github.com/sabrusrin/wb_go_task/pkg/service/httpserver"
)

type configuration struct {
	Port               string `envconfig:"PORT" default:"8080"`
	Debug              bool   `envconfig:"DEBUG" default:"true"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB

	MetricsNamespace string `envconfig:"METRICS_NAMESPACE" default:"wb"`
	MetricsSubsystem string `envconfig:"METRICS_SUBSYSTEM" default:"supply_service"`

	MetricsNameCount    string `envconfig:"METRICS_NAME_COUNT" default:"request_count"`
	MetricsNameDuration string `envconfig:"METRICS_NAME_DURATION" default:"request_duration"`
	MetricsHelpCount    string `envconfig:"METRICS_HELP_COUNT" default:"request count"`
	MetricsHelpDuration string `envconfig:"METRICS_HELP_DURATION" default:"request duration"`

	MetricsSvcNameCount        string `envconfig:"METRICS_SVC_NAME_COUNT" default:"svc_request_count"`
	MetricsSvcNameDuration     string `envconfig:"METRICS_SVC_NAME_DURATION" default:"svc_request_duration"`
	MetricsSvcHelpCount        string `envconfig:"METRICS_SVC_HELP_COUNT" default:"svc request count"`
	MetricsSvcHelpDuration     string `envconfig:"METRICS_SVC_HELP_DURATION" default:"svc request duration"`
	AscKeyForGoodIncomeService string `envconfig:"ASC_KEY_FOR_GOOD_INCOME_SERVICE" default:"asc"`
}

var serviceVersion = "dev"
var methodError = []string{"method", "error"}

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(serviceVersion)
		os.Exit(0)
	}

	// logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	// configuration
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	svc := service.NewService()

	svc = service.NewLoggingMiddleware(logger, svc)
	svc = service.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(prometheus.CounterOpts{
			Namespace: cfg.MetricsNamespace,
			Subsystem: cfg.MetricsSubsystem,
			Name:      cfg.MetricsSvcNameCount,
			Help:      cfg.MetricsSvcHelpCount,
		}, methodError),
		kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
			Namespace: cfg.MetricsNamespace,
			Subsystem: cfg.MetricsSubsystem,
			Name:      cfg.MetricsSvcNameDuration,
			Help:      cfg.MetricsSvcHelpDuration,
		}, methodError),
		svc,
	)

	router := httpserver.NewPreparedServer(svc)
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))

	p := fasthttpprometheus.NewPrometheus(cfg.MetricsSubsystem)
	fasthttpServer := &fasthttp.Server{
		Handler:            p.WrapHandler(router),
		MaxRequestBodySize: cfg.MaxRequestBodySize,
	}

	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)

}
