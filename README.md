* build
```sh
go build -x -ldflags "-linkmode external" .

go build -o bin/
```

* go env -w ...
```makefile
CGO_CFLAGS=-O2 -g -I${PROJECT_ROOT}/sdl2/x86_64-w64-mingw32/include
# if gcc -> CGO_LDFLAGS=-g -L${PROJECT_ROOT}/sdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2 -lSDL2main
CGO_LDFLAGS=-g -L${PROJECT_ROOT}/sdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2 -lSDL2main -lole32 -lgdi32 -lwinmm -lversion -lsetupapi -limm32 -loleaut32 -mwindows
CC=zig cc
```

* main.c_ -> main.c
```sh
# gcc -o main main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2main -lSDL2 -lole32 -lgdi32 -lversion -lsetupapi -limm32 -mwindows
gcc -o main main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2main -lSDL2 -mwindows
zig cc -target x86_64-windows-gnu main.c -Isdl2/x86_64-w64-mingw32/include/SDL2 -Lsdl2/x86_64-w64-mingw32/lib -lmingw32 -lSDL2 -lSDL2main -lole32 -lgdi32 -lwinmm -lversion -lsetupapi -limm32 -loleaut32 -mwindows
```

* Refs.
    * https://github.com/veandco/go-sdl2
    * https://github.com/libsdl-org/SDL
    * https://github.com/abhirag/go-sdl2-tutorial?tab=readme-ov-file
    * https://lazyfoo.net/tutorials/SDL
