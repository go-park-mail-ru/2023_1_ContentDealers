package domain

import "time"

type ViewsOptions struct {
	UserID       uint64
	SortDate     string
	Limit        uint32
	Offset       uint32
	TypeView     string
	ViewProgress float32
}

type View struct {
	UserID     uint64
	ContentID  uint64
	DateAdding time.Time
	StopView   time.Duration
	Duration   time.Duration
}

type Views struct {
	IsLast bool
	Views  []View
}

type HasView struct {
	View    View
	HasView bool
}
