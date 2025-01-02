package ports

type LogPublisher interface {
	PublishLog(logData []byte) error
}
