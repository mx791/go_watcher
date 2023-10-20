package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var targetDirPath = "./target"                     // os.Getenv("TARGET_DIR_PATH")
var gitUrl = "https://github.com/mx791/go_watcher" // os.Getenv("GIT_URL")
var period = 15                                    // os.Getenv("PERIOD")

func main() {
	fmt.Println("Starting GoWatcher")

	// supression du dossier si re-clone
	if gitUrl != "" {
		cmd := exec.Command("rm", "-rf", targetDirPath)
		cmd.Run()
	}

	// création du dossier si absent
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

	// clone du repertoire si gitUrl spécifié
	if gitUrl != "" {
		cmd := exec.Command("git", "clone", gitUrl)
		cmd.Dir = targetDirPath
		err = cmd.Run()

		if err != nil {
			fmt.Println("Cannot fetch repository " + gitUrl)
			fmt.Println(err)
		} else {
			// on remonte tous les fichiers d'un niveau
			cmd = exec.Command("cp", "*/*", ".")
			cmd.Dir = targetDirPath
			cmd.Run()
		}
	}

	for true {

		fmt.Println("Scanning repository...")

		cmd := exec.Command("git", "pull")
		cmd.Dir = targetDirPath
		out, _ := cmd.CombinedOutput()
		outString := string(out)

		if strings.Contains(outString, "Already up to date") {
			fmt.Println("Nothing to update")
		} else {
			fmt.Println("Updating repository")
		}

		time.Sleep(time.Duration(period) * time.Second)
	}

}
