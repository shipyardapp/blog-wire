package apm

type APM interface {
	CaptureError(err error)
}
