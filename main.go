package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var targetDirPath = "./target"                     // os.Getenv("TARGET_DIR_PATH")
var gitUrl = "https://github.com/mx791/go_watcher" // os.Getenv("GIT_URL")
var period = 15                                    // os.Getenv("PERIOD")

func main() {
	fmt.Println("Starting GoWatcher")

	_, err := os.Stat(targetDirPath)
	if err != nil {
		fmt.Println("Cannot find directory " + targetDirPath)
		cmd := exec.Command("mkdir", targetDirPath)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Cannot create directory " + targetDirPath)
			return
		}
	}

	if gitUrl != "" {
		cmd := exec.Command("git", "clone", gitUrl)
		cmd.Dir = targetDirPath
		err = cmd.Run()

		if err != nil {
			fmt.Println("Cannot fetch repository " + gitUrl)
			fmt.Println(err)
			//return
		}

		cmd = exec.Command("cp", "*/*", ".")
		cmd.Dir = targetDirPath
		cmd.Run()
	}

	for true {

		cmd := exec.Command("git", "pull")
		cmd.Dir = targetDirPath
		out, _ := cmd.CombinedOutput()
		outString := string(out)

		fmt.Println(outString)

		time.Sleep(60 * time.Second)
	}

}
