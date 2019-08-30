package main

import (
	"fmt"
	"os"

	"github.com/Permaweb/host"
	"github.com/Permaweb/host/internal/exec"
	"github.com/logrusorgru/aurora"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println(aurora.Bold("Permaweb Host :"), "Upload a Git repo to IPNS.")
	fmt.Println("Copyright © 2019 Nato Boram, Permaweb")
	fmt.Println(aurora.Bold("Contact :"), aurora.Blue("https://github.com/Permaweb/Host"))
	fmt.Println("")

	// User
	err := initUser()
	if err != nil {
		return
	}

	// IPFS
	err = initIPFS()
	if err != nil {
		return
	}

	// Git
	err = initGit()
	if err != nil {
		return
	}

}

func initUser() (err error) {

	host.RootConfig, err = os.UserConfigDir()
	if err != nil {
		fmt.Println("Couldn't get the default root directory to use for user-specific configuration data.")
		fmt.Println(err.Error())
	}

	host.RootCache, err = os.UserCacheDir()
	if err != nil {
		fmt.Println("Couldn't get the default root directory to use for user-specific cached data.")
		fmt.Println(err.Error())
	}

	return err
}

func initIPFS() (err error) {

	// Check for IPFS
	path, err := exec.LookPath("ipfs")
	if err != nil {
		fmt.Println("IPFS is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS :"), aurora.Blue(path))

	// Enable sharding
	exec.Command("ipfs", "config", "--json", "Experimental.ShardingEnabled", "true").Run()

	// Check for IPFS Cluster Service
	path, err = exec.LookPath("ipfs-cluster-service")
	if err != nil {
		fmt.Println("IPFS Cluster Service is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Service :"), aurora.Blue(path))

	// Check for IPFS Cluster Control
	path, err = exec.LookPath("ipfs-cluster-ctl")
	if err != nil {
		fmt.Println("IPFS Cluster Control is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Control :"), aurora.Blue(path))

	// Connect to Swarm
	initSwarm()

	fmt.Println("")
	return
}

func initSwarm() {
	for _, pg := range host.PublicGateways {
		exec.ipfsSwarmConnect(pg)
	}
}

func initGit() (err error) {

	// Git Directory
	err = os.MkdirAll(host.RootCache+host.DirGit, host.PermPrivateDirectory)
	if err != nil {
		fmt.Println("Couldn't create the git directory.")
		fmt.Println(err.Error())
		return
	}

	// Check for Git
	path, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("Git is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("Git :"), aurora.Blue(path))

	fmt.Println("")
	return
}
