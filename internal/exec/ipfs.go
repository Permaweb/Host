package exec

import (
	"net/url"
	"os/exec"

	"github.com/logrusorgru/aurora"
)

func ipfsKeyGen(link string) (out []byte, err error) {
	escaped := url.PathEscape(link)
	return run(
		exec.Command("ipfs", "key", "gen", "--type", "ed25519", escaped),
		".",
		"Couldn't generate a new key.",
		aurora.Bold("Command :"), "ipfs", "key", "gen", "--type", "ed25519", aurora.Blue(escaped),
	)
}

func ipfsKeyRm(key string) (out []byte, err error) {
	return run(
		exec.Command("ipfs", "key", "rm", key),
		".",
		"Couldn't remove a key.",
		aurora.Bold("Command :"), "ipfs", "key", "rm", key,
	)
}

func ipfsKeyRmName(name string) (out []byte, err error) {
	return ipfsKeyRm(url.PathEscape(name))
}

func ipfsNamePublish(key string, ipfs string) (out []byte, err error) {
	return run(
		exec.Command("ipfs", "name", "publish", "--key", key, "--quieter", "/ipfs/"+ipfs),
		".",
		"Couldn't publish on IPNS.",
		aurora.Bold("Command :"), "ipfs", "name", "publish", "--key", key, "--quieter", aurora.Cyan("/ipfs/"+ipfs),
	)
}
