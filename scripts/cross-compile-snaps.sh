#!/bin/sh

snapcraft clean
snapcraft clean permaweb-host; snapcraft snap --target-arch=i386     --output build/permaweb-host_i386.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=amd64    --output build/permaweb-host_amd64.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=arm      --output build/permaweb-host_arm.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=arm64    --output build/permaweb-host_arm64.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=mips     --output build/permaweb-host_mips.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=mips64   --output build/permaweb-host_mips64.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=mips64le --output build/permaweb-host_mips64le.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=mipsle   --output build/permaweb-host_mipsle.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=ppc64    --output build/permaweb-host_ppc64.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=ppc64le  --output build/permaweb-host_ppc64le.snap
snapcraft clean permaweb-host; snapcraft snap --target-arch=s390x    --output build/permaweb-host_s390x.snap