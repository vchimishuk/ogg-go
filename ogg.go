package ogg

// #include <stdio.h>
// #include <stdlib.h>
// #include <vorbis/codec.h>
// #include <vorbis/vorbisfile.h>
// #include "comment_hlp.h"
import "C"

import (
	"os"
	"unsafe"
)

type File struct {
	cOggFile C.OggVorbis_File
}

func New(filename string) (file *File, err os.Error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	file = new(File)

	r := C.ov_fopen(cFilename, &(file.cOggFile))
	if r != 0 {
		return nil, os.NewError("Failed to open file")
	}

	return file, nil
}

func (file *File) Comment() *Comment {
	cComment := C.ov_comment(&(file.cOggFile), -1)

	comment := new(Comment)
	comment.UserComments = make([]string, cComment.comments)
	for i := 0; i < int(cComment.comments); i++ {
		cUc := C.comment_hlp_get_user_comment(cComment, _Ctype_int(i))
		comment.UserComments[i] = C.GoString(cUc)
	}
	comment.Vendor = C.GoString(cComment.vendor)

	return comment
}

func (file *File) Info() *Info {
	cInfo := C.ov_info(&(file.cOggFile), -1)

	info := new(Info)
	info.Version = int(cInfo.version)
	info.Channels = int(cInfo.channels)
	info.Rate = int32(cInfo.rate)
	info.BitrateUpper = int32(cInfo.bitrate_upper)
	info.BitrateNominal = int32(cInfo.bitrate_nominal)
	info.BitrateLower = int32(cInfo.bitrate_lower)
	info.BitrateWindow = int32(cInfo.bitrate_window)

	return info
}

func (file *File) Close() {
	C.ov_clear(&(file.cOggFile))
}
