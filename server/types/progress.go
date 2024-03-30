package types

import "time"

func ProgresNow(t ProgressType, content string) *Progress {
	return &Progress{Timestamp: time.Now().Unix(), Type: t, Content: content}
}
