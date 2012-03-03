include $(GOROOT)/src/Make.inc

PKG_CONFIG=$(shell which pkg-config)
CGO_LDFLAGS=$(shell $(PKG_CONFIG) --libs icu-i18n)
CGO_CFLAGS=$(shell $(PKG_CONFIG) --cflags icu-i18n)

TARG=icu4go

CGOFILES=\
  charsetdetect.go\

CGO_OFILES=\
  helper.o\

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%.o: %.c
	gcc $(_CGO_CFLAGS_$(GOARCH)) -g -O2 -fPIC $(CFLAGS) -o $@ -c $^
