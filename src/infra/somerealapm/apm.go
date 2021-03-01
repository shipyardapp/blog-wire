package somerealapm

import "github.com/shipyardapp/blog-wire/src/config"

type Config struct {
	ServiceCaptureEndpoint string
}

func NewConfig(enver config.Enver) (Config, error) {
	var err error
	c := Config{}

	if c.ServiceCaptureEndpoint, err = config.RequiredString(enver, "BLOGWIRE_APM_CAPTURE_ENDPOINT"); err != nil {
		return Config{}, err
	}

	return c, nil
}

type APM struct {
	// someClient
}

func New(config Config) (*APM, error) {
	return &APM{
		// someClient using config.ServiceCaptureEndpoint.
	}, nil
}

func (a *APM) Close() error {
	// return a.someClient.Close()

	return nil
}

func (a *APM) CaptureError(err error) {
	// This would actually call some service in the ral world.
}
