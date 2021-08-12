package watcher

import (
	sentry "github.com/getsentry/sentry-go"
)

func CaptureMessage(msg string) {
	sentry.CaptureMessage(msg)
}

func CaptureError(err error) {
	sentry.CaptureException(err)
}

func CaptureRecover() {
	sentry.Recover()
}

func CaptureEvent(msg string, extra map[string]string) {
	event := sentry.NewEvent()
	event.Message = msg

	for k, v := range extra {
		event.Extra[k] = v
	}

	sentry.CaptureEvent(event)

}
