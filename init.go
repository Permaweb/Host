package host

import (
	"fmt"
	"os"
)

func init() {

	var err error

	RootConfig, err = os.UserConfigDir()
	if err != nil {
		fmt.Println("Couldn't get the default root directory to use for user-specific configuration data.")
		panic(err.Error())
	}

	RootCache, err = os.UserCacheDir()
	if err != nil {
		fmt.Println("Couldn't get the default root directory to use for user-specific cached data.")
		panic(err.Error())
	}
}
