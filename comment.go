package ogg

import "C"

// The Comment structure defines an Ogg Vorbis comment.
type Comment struct {
	UserComments []string
	Vendor       string
}
