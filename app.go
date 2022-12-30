package main

import (
	"log"
	"os"
	"os/exec"
)

type App struct {
	Repo       string
	Branch     string
	RepoFolder string
}

func (a *App) initRepo() error {
	_, err := os.Stat(a.RepoFolder)
	if os.IsExist(err) {
		log.Print("Repository already initialized")
		return nil
	}

	log.Print("Initializing repository")
	err = os.MkdirAll(a.RepoFolder, 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "--git-dir="+a.RepoFolder, "init")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "--git-dir="+a.RepoFolder, "remote", "add", "origin", a.Repo)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	log.Printf("Repository initialized")

	return nil
}

func (a *App) fetchChanges() error {
	_, err := os.Stat(a.RepoFolder + "stop.sh")
	if os.IsExist(err) {
		log.Print("Running stop.sh")
	} else {
		log.Print("No stop.sh file found in repository folder")
	}

	log.Print("Fetching changes")

	cmd := exec.Command("git", "--git-dir="+a.RepoFolder, "fetch", "-f", "origin", a.Branch+":"+a.Branch)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	log.Print("Finished fetching")

	_, err = os.Stat(a.RepoFolder + "start.sh")
	if os.IsExist(err) {
		log.Print("Running start.sh")
	} else {
		log.Print("No start.sh file found in repository folder")
	}

	return nil
}
