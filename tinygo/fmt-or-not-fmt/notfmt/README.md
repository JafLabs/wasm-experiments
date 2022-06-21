## Example not using "fmt" package

```
tinygo build -o file.wasm -target=wasi -wasm-abi=generic -gc=leaking -no-debug
wasm-opt -O file.wasm -o file.wasm
```

tinygo v0.23.0
wasm-opt v107

Output 7.9k
