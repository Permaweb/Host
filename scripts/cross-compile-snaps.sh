#!/bin/sh

snapcraft clean --destructive-mode
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=386      snapcraft snap --target-arch=i386     --destructive-mode --output build/permaweb-host_i386.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=amd64    snapcraft snap --target-arch=amd64    --destructive-mode --output build/permaweb-host_amd64.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=arm      snapcraft snap --target-arch=arm      --destructive-mode --output build/permaweb-host_arm.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=arm64    snapcraft snap --target-arch=arm64    --destructive-mode --output build/permaweb-host_arm64.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=mips     snapcraft snap --target-arch=mips     --destructive-mode --output build/permaweb-host_mips.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=mips64   snapcraft snap --target-arch=mips64   --destructive-mode --output build/permaweb-host_mips64.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=mips64le snapcraft snap --target-arch=mips64le --destructive-mode --output build/permaweb-host_mips64le.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=mipsle   snapcraft snap --target-arch=mipsle   --destructive-mode --output build/permaweb-host_mipsle.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=ppc64    snapcraft snap --target-arch=ppc64    --destructive-mode --output build/permaweb-host_ppc64.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=ppc64le  snapcraft snap --target-arch=ppc64le  --destructive-mode --output build/permaweb-host_ppc64le.snap
snapcraft clean --destructive-mode permaweb-host; env GOOS=linux GOARCH=s390x    snapcraft snap --target-arch=s390x    --destructive-mode --output build/permaweb-host_s390x.snap