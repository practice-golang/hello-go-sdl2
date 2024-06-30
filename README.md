Taste `go-sdl2` cross compilation using `zig cc`


## Test c using [MinGW](https://github.com/brechtsanders/winlibs_mingw) or [Zig](https://github.com/ziglang/zig)
* Change main.c_ -> main.c then run below
```sh
# gcc -o main main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2main -lSDL2 -lole32 -lgdi32 -lversion -lsetupapi -limm32 -mwindows
gcc -o main main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2main -lSDL2 -mwindows
zig cc -target x86_64-windows-gnu main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2 -lSDL2main -lole32 -lgdi32 -lwinmm -lversion -lsetupapi -limm32 -loleaut32 -mwindows
```

## go-sdl2
is Added to `go.work` and remarked `pkg-config` using in all go source
* from 
```go
//#cgo linux ... pkg-config: sdl2
```
* to
```go
// //#cgo linux ... pkg-config: sdl2
```

## env location
```sh
go env GOENV
```

## go env
* `SRCDIR` = Project root

* windows with `GOARCH=amd64` and `GOOS=windows`
```sh
CGO_ENABLED=1
CGO_CFLAGS=-O2 -g -I${SRCDIR}/go-sdl2/_libs/include -I${SRCDIR}/go-sdl2/_libs/include/SDL2
CGO_LDFLAGS=-g -L${SRCDIR}/lib/windows_amd64_mingw -lmingw32 -lSDL2 -lSDL2main -lole32 -lgdi32 -lwinmm -lversion -lsetupapi -limm32 -loleaut32 -mwindows
CC=zig cc -target x86_64-windows-gnu
```

* linux amd64 with `GOARCH=amd64` and `GOOS=linux`
```sh
CGO_ENABLED=1
CGO_CFLAGS=-O2 -g -I${SRCDIR}/go-sdl2/_libs/include -I${SRCDIR}/go-sdl2/_libs/include/SDL2
CGO_LDFLAGS=-g -static -L${SRCDIR}/lib/linux_amd64 -lSDL2 -lSDL2main
CC=zig cc -target x86_64-linux-gnu
```

## build
```sh
# go build -x -ldflags "-linkmode external" .
# go build -ldflags "-linkmode external" .
# go build -x .
go build .
```

* Refs.
    * https://github.com/veandco/go-sdl2
    * https://github.com/libsdl-org/SDL
    * https://ziglang.org/download
    * https://github.com/abhirag/go-sdl2-tutorial?tab=readme-ov-file
    * https://lazyfoo.net/tutorials/SDL
