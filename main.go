package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var targetDirPath = os.Getenv("TARGET_DIR_PATH")
var gitUrl = os.Getenv("GIT_URL")
var period, _ = strconv.Atoi(os.Getenv("PERIOD"))
var postUpdate = os.Getenv("POST_UPDATE")

func Log(msg string) {
	dt := time.Now()
	fmt.Println(dt.String(), msg)
}

func main() {

	Log("Starting GoWatcher with :")
	Log("TARGET_DIR_PATH = " + targetDirPath)
	Log("GIT_URL = " + gitUrl)
	Log("POST_UPDATE = " + postUpdate)

	// supression du dossier si re-clone
	if gitUrl != "" {
		Log("Removing directory " + targetDirPath)
		cmd := exec.Command("rm", "-rf", targetDirPath)
		cmd.Run()
	}

	// création du dossier si absent
	_, err := os.Stat(targetDirPath)
	if err != nil {
		Log("Cannot find directory " + targetDirPath)
		cmd := exec.Command("mkdir", targetDirPath)
		err = cmd.Run()
		if err != nil {
			Log("Cannot create directory " + targetDirPath)
			return
		}
	}

	// clone du repertoire si gitUrl spécifié
	if gitUrl != "" {
		cmd := exec.Command("git", "clone", gitUrl, targetDirPath)
		cmd.Dir = targetDirPath
		err = cmd.Run()

		if err != nil {
			Log("Cannot fetch repository " + gitUrl)
			fmt.Println(err)
		} else {
			// on remonte tous les fichiers d'un niveau
			cmd = exec.Command("cp", "*/*", ".")
			cmd.Dir = targetDirPath
			cmd.Run()
		}
	}

	for true {

		time.Sleep(time.Duration(period) * time.Second)

		cmd := exec.Command("git", "pull")
		cmd.Dir = targetDirPath
		out, err := cmd.CombinedOutput()
		outString := string(out)

		if err != nil {
			Log("Something went wrong while pulling repository")
			continue
		}

		if strings.Contains(outString, "Already up to date") {
			Log("Nothing to update")
			continue
		}

		Log("Updating repository")
		if postUpdate != "" {
			cmd = exec.Command("sh", "-c", postUpdate)
			cmd.Dir = targetDirPath
			err = cmd.Run()

			if err == nil {
				Log("Post-update ran successfully")
			} else {
				Log("Error with post-update script : " + err.Error())
			}
		}

	}

}
