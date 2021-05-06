package statsd

import "fmt"

type Publisher interface {
	PackageDuration(duration int64)
}

type publisher struct {
	client Client
	labels Labels
}

type Labels struct {
	ProjectPath string
	PackageName string
}

func (p *publisher) PackageDuration(duration int64) {
	bucket := fmt.Sprintf("duration.timer.%s", p.formattedLabels())
	p.client.Timing(bucket, duration)
}

func (p *publisher) formattedLabels() string {
	return fmt.Sprintf("%s.%s",
		p.labels.ProjectPath,
		p.labels.PackageName,
	)
}

func NewPublisher(client Client, labels Labels) (publisherInstance Publisher) {
	publisherInstance = &publisher{
		client: client,
		labels: labels,
	}
	return
}
