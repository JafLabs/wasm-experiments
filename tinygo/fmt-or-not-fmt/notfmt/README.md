## Example not using "fmt" package

```
tinygo build -o file.wasm -target=wasi -wasm-abi=generic -gc=leaking -no-debug
wasm-opt -O file.wasm -o file.wasm
```

Output 7.9k
