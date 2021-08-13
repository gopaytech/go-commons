package statsd_test

import (
	"testing"

	"github.com/gopaytech/go-commons/pkg/statsd"
)

type publisherTestContext struct {
	publisher  statsd.Publisher
	clientMock *statsd.ClientMock
}

func (ctx *publisherTestContext) setUp(t *testing.T) {
	ctx.clientMock = &statsd.ClientMock{}
	labels := statsd.Labels{
		ProjectPath: "project-path",
		PackageName: "package-name",
	}
	ctx.publisher = statsd.NewPublisher(ctx.clientMock, labels)
}

func (ctx *publisherTestContext) tearDown() {
}

func TestPublisher_PackageDuration(t *testing.T) {
	ctx := &publisherTestContext{}
	ctx.setUp(t)
	defer ctx.tearDown()

	ctx.clientMock.On("Timing", "duration.timer.project-path.package-name", int64(10000)).Once()

	ctx.publisher.PackageDuration(10000)

	ctx.clientMock.AssertExpectations(t)
}
