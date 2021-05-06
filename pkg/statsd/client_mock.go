package statsd

import "github.com/stretchr/testify/mock"

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Flush() {
	c.Called()
}

func (c *ClientMock) Increment(bucket string) {
	c.Called(bucket)
}

func (c *ClientMock) Timing(bucket string, value interface{}) {
	c.Called(bucket, value)
}
