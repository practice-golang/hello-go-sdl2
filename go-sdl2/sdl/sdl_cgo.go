//go:build !static
// +build !static

package sdl

//#cgo LDFLAGS: -lSDL2
// //#cgo windows LDFLAGS: -lSDL2
// //#cgo linux freebsd darwin openbsd pkg-config: sdl2
import "C"
