package somerealapm

import "github.com/shipyardapp/blog-wire/src/config"

type Config struct {
	ServiceCaptureEndpoint string
}

func NewConfig(enver config.Enver) (Config, error) {
	// TODO
	return Config{}, nil
}

type APM struct {
	// someClient
}

func New(config Config) (*APM, error) {
	return &APM{
		// someClient from config.ServiceCaptureEndpoint.
	}, nil
}

func (a *APM) Close() error {
	// return a.someClient.Close()

	return nil
}

func (a *APM) CaptureError(err error) {
	// This would actually call some service in the ral world.
}
