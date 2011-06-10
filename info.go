package ogg

import "C"

type Info struct {
	Version        int
	Channels       int
	Rate           int32
	BitrateUpper   int32
	BitrateNominal int32
	BitrateLower   int32
	BitrateWindow  int32
}
