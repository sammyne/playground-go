# Forwarded declared C struct shared the same memory address in cgo

Workaround wanted, any sugguestion?

## Environment
- OS: ubuntu: 18.04
- CMake: >=3.10
- Compiler: g++-7.4

> This environment can be reproduced by the [docker/Dockerfile](docker/Dockerfile)

## Reproduction

```bash
mkdir cpp/build && cd cpp/build
cmake ..
make 

cd - 
go run world.go
```

## Expect 

- Different address for `x` and `y`
- Different address for `a` and `b`

Something likes

```bash
x: 0x777280
y: 0x777284

a: 0xc00001a080
b: 0xc00001a084
```

## Got

- Same address for `x` and `y`
- Different address for `a` and `b`

```bash
x: 0x777280
y: 0x777280

a: 0xc00001a080
b: 0xc00001a084
```

`unsafe.Sizeof(x) == 0`，I think this is the root cause, so any workaround?