package native

/*
#cgo CFLAGS: -I${SRCDIR}/zlib/
#cgo windows LDFLAGS: ${SRCDIR}/libs/winlibz.a
#cgo linux,amd64 LDFLAGS: ${SRCDIR}/libs/linuxlibz.a
#cgo linux,arm64 LDFLAGS: ${SRCDIR}/libs/linuxarm64libz.a
#cgo darwin LDFLAGS: ${SRCDIR}/libs/darwinlibz.a
*/
import "C"
