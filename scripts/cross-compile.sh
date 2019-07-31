#!/bin/sh

./scripts/build.sh

env GOOS=aix        GOARCH=ppc64    go build -o build/permaweb-host-aix-ppc64
env GOOS=android    GOARCH=386      go build -o build/permaweb-host-android-386
env GOOS=android    GOARCH=amd64    go build -o build/permaweb-host-android-amd64
env GOOS=android    GOARCH=arm      go build -o build/permaweb-host-android-arm
env GOOS=android    GOARCH=arm64    go build -o build/permaweb-host-android-arm64
env GOOS=darwin     GOARCH=386      go build -o build/permaweb-host-darwin-386
env GOOS=darwin     GOARCH=amd64    go build -o build/permaweb-host-darwin-amd64
env GOOS=darwin     GOARCH=arm      go build -o build/permaweb-host-darwin-arm
env GOOS=darwin     GOARCH=arm64    go build -o build/permaweb-host-darwin-arm64
env GOOS=dragonfly  GOARCH=amd64    go build -o build/permaweb-host-dragonfly-amd64
env GOOS=freebsd    GOARCH=386      go build -o build/permaweb-host-freebsd-386
env GOOS=freebsd    GOARCH=amd64    go build -o build/permaweb-host-freebsd-amd64
env GOOS=freebsd    GOARCH=arm      go build -o build/permaweb-host-freebsd-arm
env GOOS=js         GOARCH=wasm     go build -o build/permaweb-host-js-wasm
env GOOS=linux      GOARCH=386      go build -o build/permaweb-host-linux-386
env GOOS=linux      GOARCH=amd64    go build -o build/permaweb-host-linux-amd64
env GOOS=linux      GOARCH=arm      go build -o build/permaweb-host-linux-arm
env GOOS=linux      GOARCH=arm64    go build -o build/permaweb-host-linux-arm64
env GOOS=linux      GOARCH=mips     go build -o build/permaweb-host-linux-mips
env GOOS=linux      GOARCH=mips64   go build -o build/permaweb-host-linux-mips64
env GOOS=linux      GOARCH=mips64le go build -o build/permaweb-host-linux-mips64le
env GOOS=linux      GOARCH=mipsle   go build -o build/permaweb-host-linux-mipsle
env GOOS=linux      GOARCH=ppc64    go build -o build/permaweb-host-linux-ppc64
env GOOS=linux      GOARCH=ppc64le  go build -o build/permaweb-host-linux-ppc64le
env GOOS=linux      GOARCH=s390x    go build -o build/permaweb-host-linux-s390x
env GOOS=nacl       GOARCH=386      go build -o build/permaweb-host-nacl-386
env GOOS=nacl       GOARCH=amd64p32 go build -o build/permaweb-host-nacl-amd64p32
env GOOS=nacl       GOARCH=arm      go build -o build/permaweb-host-nacl-arm
env GOOS=netbsd     GOARCH=386      go build -o build/permaweb-host-netbsd-386
env GOOS=netbsd     GOARCH=amd64    go build -o build/permaweb-host-netbsd-amd64
env GOOS=netbsd     GOARCH=arm      go build -o build/permaweb-host-netbsd-arm
env GOOS=openbsd    GOARCH=386      go build -o build/permaweb-host-openbsd-386
env GOOS=openbsd    GOARCH=amd64    go build -o build/permaweb-host-openbsd-amd64
env GOOS=openbsd    GOARCH=arm      go build -o build/permaweb-host-openbsd-arm
env GOOS=plan9      GOARCH=386      go build -o build/permaweb-host-plan9-386
env GOOS=plan9      GOARCH=amd64    go build -o build/permaweb-host-plan9-amd64
env GOOS=plan9      GOARCH=arm      go build -o build/permaweb-host-plan9-arm
env GOOS=solaris    GOARCH=amd64    go build -o build/permaweb-host-solaris-amd64
env GOOS=windows    GOARCH=386      go build -o build/permaweb-host-windows-386.exe
env GOOS=windows    GOARCH=amd64    go build -o build/permaweb-host-windows-amd64.exe
env GOOS=windows    GOARCH=arm      go build -o build/permaweb-host-windows-arm.exe