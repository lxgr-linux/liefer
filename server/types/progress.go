package types

import "time"

func ProgresNow(content string) *Progress {
	return &Progress{Timestamp: uint64(time.Now().Unix()), Content: content}
}
