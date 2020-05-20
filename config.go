package sqsworker

const defaultMaxNumberOfMessage = 10
const defaultWaitTimeSecond = 20
const defaultVisibilityTimeout = 20

// Config struct
type Config struct {
	MaxNumberOfMessage int64
	QueueURL           string
	WaitTimeSecond     int64
	VisibilityTimeout  int64
	WorkerSize         int64
	Logger             Logger
}

func (config *Config) populateDefaultValues() {
	if config.MaxNumberOfMessage == 0 {
		config.MaxNumberOfMessage = defaultMaxNumberOfMessage
	}

	if config.WaitTimeSecond == 0 {
		config.WaitTimeSecond = defaultWaitTimeSecond
	}

	if config.VisibilityTimeout == 0 {
		config.VisibilityTimeout = defaultVisibilityTimeout
	}

	if config.WorkerSize == 0 {
		config.WorkerSize = 1
	}

	if config.Logger == nil {
		config.Logger = &logger{}
	}
}
