include $(GOROOT)/src/Make.inc

#ONIG_CONFIG=$(shell which onig-config)
#CGO_LDFLAGS=$(shell $(ONIG_CONFIG) --libs)
#CGO_CFLAGS=$(shell $(ONIG_CONFIG) --cflags)

TARG=icu4go

CGOFILES=\
  charsetdetect.go\

CGO_OFILES=\
  helper.o\

CLEANFILES+=

include $(GOROOT)/src/Make.pkg

%.o: %.c
	gcc $(_CGO_CFLAGS_$(GOARCH)) -g -O2 -fPIC $(CFLAGS) -o $@ -c $^
