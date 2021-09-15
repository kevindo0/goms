#### 1、run
```bash
# mac 动态加载so库 添加env
cd include
gcc -shared -o ../lib/libvideo.so video.c
go run -exec "env DYLD_LIBRARY_PATH=/Users/kkdu/mygit/goms/pkgtest/cgo/lib" main.go
go run -exec "env DYLD_LIBRARY_PATH=`pwd`/cgo/lib" main.go
# 只修改c代码不修改go代码时会使用编译的缓存
# -a 强制重新编译
go run -a main.go
```