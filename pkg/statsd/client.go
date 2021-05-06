package statsd

import "gopkg.in/alexcesaro/statsd.v2"

type Client interface {
	Flush()
	Increment(bucket string)
	Timing(bucket string, value interface{})
}

type ClientConfig struct {
	Host     string
	Protocol string
	Prefix   string
	Muted    bool
}

func (config ClientConfig) StatsdOptions() (options []statsd.Option) {
	options = []statsd.Option{}
	if config.Host != "" {
		options = append(options, statsd.Address(config.Host))
	}
	if config.Protocol != "" {
		options = append(options, statsd.Network(config.Protocol))
	}
	if config.Prefix != "" {
		options = append(options, statsd.Prefix(config.Prefix))
	}
	options = append(options, statsd.Mute(config.Muted))
	return options
}

func NewClient(config ClientConfig) (client Client, err error) {
	client, err = statsd.New(config.StatsdOptions()...)
	return
}
