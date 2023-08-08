package dbModels

import "fmt"

type EventType int64

const (
	Welcome EventType = iota
	UserSubscribed
	UserUnsubscribed
	NewVideo
)

func (e EventType) String(userName string) string {
	switch e {
	case Welcome:
		return fmt.Sprintf("Welcome %s to our site", userName)
	case UserSubscribed:
		return fmt.Sprintf("%s has subscribed", userName)
	case UserUnsubscribed:
		return fmt.Sprintf("%s has unsubscribed", userName)
	case NewVideo:
		return fmt.Sprintf("%s has posted a new video", userName)
	}
	return "unknown"
}
