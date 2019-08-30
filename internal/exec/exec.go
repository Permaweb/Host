package exec

import (
	"fmt"
	"os/exec"
)

func run(cmd *exec.Cmd, path string, errMessage string, cmdMessage ...interface{}) (out []byte, err error) {

	// Default path
	if path != "." {
		cmd.Dir = path
	}

	out, err = cmd.Output()
	if err != nil {
		fmt.Println(errMessage)
		fmt.Println(cmdMessage...)

		// Log the `ExitError`
		ee, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println(string(ee.Stderr))
		}

		fmt.Println(string(out))
		return
	}
	return
}
