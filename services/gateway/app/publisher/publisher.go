package publisher

import "context"

type CommandPublisher interface {
	Publish(ctx context.Context, body interface{}, key string) error
}
