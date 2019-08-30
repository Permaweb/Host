package exec

import (
	"os/exec"

	"github.com/Permaweb/host"
	"github.com/logrusorgru/aurora"
)

func ipfsClusterAdd(link string, rmin string, rmax string, uuid string) (out []byte, err error) {
	return run(
		exec.Command("ipfs-cluster-ctl", "add", "--recursive", "--quieter", "--chunker=rabin", "--cid-version=1", "--name", link, "--replication-min", rmin, "--replication-max", rmax, uuid),
		host.RootCache+host.DirGit,
		"Couldn't add the repository to IPFS.",
		aurora.Bold("Command :"), "ipfs-cluster-ctl", "add", "--recursive", "--quieter", "--chunker=rabin", "--cid-version=1", "--name", aurora.Blue(link), "--replication-min", aurora.Bold(rmin), "--replication-max", aurora.Bold(rmax), uuid,
	)
}

func ipfsClusterRm(ipfs string) (out []byte, err error) {
	return run(
		exec.Command("ipfs-cluster-ctl", "pin", "rm", ipfs),
		".",
		"Couldn't remove the repository from IPFS.",
		aurora.Bold("Command :"), "ipfs-cluster-ctl", "pin", "rm", aurora.Cyan(ipfs),
	)
}
