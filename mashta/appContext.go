package mashta

type ApplicationContext struct {
	HttpServer      HttpServer
	Retrier         MessageRetrier
	Config          Config
	StreamRouter    *StreamRouter
	MetricPublisher MetricPublisher
}
