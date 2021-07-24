package server

var (
	StatHitsKey    int
	StatMissesKey  int
	StatExpiredKey int
)

const (
	LookupNone     = 0
	LookupNoTouch  = 1 << 0
	LookupNoNotify = 1 << 1
)
