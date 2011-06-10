include $(GOROOT)/src/Make.inc

TARG=ogg

CGOFILES=\
        ogg.go\
	comment.go\
	info.go\

CGO_LDFLAGS=-lvorbisfile -lvorbis -lm -logg

CGO_OFILES=\
        comment_hlp.o\

include $(GOROOT)/src/Make.pkg

format:
	find . -type f -name '*.go' -exec gofmt -w {} \;
