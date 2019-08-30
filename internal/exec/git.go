package exec

import (
	"os/exec"

	"github.com/Permaweb/host"
	"github.com/logrusorgru/aurora"
)

func gitClone(link string, uuid string) (out []byte, err error) {
	return run(
		exec.Command("git", "clone", link, uuid),
		host.RootCache+host.DirGit,
		"Couldn't clone the repository.",
		aurora.Bold("Command :"), "git", "clone", aurora.Blue(link), uuid,
	)
}

func gitPull(uuid string) (out []byte, err error) {
	return run(
		exec.Command("git", "-C", uuid, "pull"),
		host.RootCache+host.DirGit,
		"Couldn't pull the repository.",
		aurora.Bold("Command :"), "git", "-C", uuid, "pull",
	)
}

func ipfsSwarmConnect(address string) (out []byte, err error) {
	return run(
		exec.Command("ipfs", "swarm", "connect", address),
		".",
		"Couldn't connect to "+address+".",
		aurora.Bold("Command :"), "ipfs", "swarm", "connect", aurora.Cyan(address),
	)
}
