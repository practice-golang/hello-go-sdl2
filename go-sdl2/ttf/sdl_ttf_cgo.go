//go:build !static
// +build !static

package ttf

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_ttf
// //#cgo linux freebsd darwin pkg-config: sdl2_ttf
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_ttf
import "C"
