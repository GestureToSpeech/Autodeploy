package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type App struct {
	Repo       string
	Branch     string
	MainFolder string
	RepoFolder string
}

func NewApp(repo string, branch string, mainFolder string) *App {
	repoSSHParts := strings.Split(repo, "/")
	repoName := repoSSHParts[len(repoSSHParts)-1]
	repoName = strings.TrimSuffix(repoName, ".git")

	a := &App{
		Repo:       repo,
		Branch:     branch,
		MainFolder: mainFolder,
		RepoFolder: mainFolder,
	}

	return a
}

func (a *App) initRepo() error {
	_, err := os.Stat(a.RepoFolder)
	if os.IsExist(err) {
		log.Print("Repository already initialized")
		return nil
	}

	log.Print("Initializing repository")
	cmd := exec.Command("bash", "-C", a.MainFolder, "\"git clone "+a.Repo+"\"")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("bash", "-C", a.RepoFolder, "\"git checkout "+a.Branch+"\"")
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

	cmd := exec.Command("bash", "-C", a.RepoFolder, "\"git fetch -f origin\"")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("bash", "-C", a.RepoFolder, "\"git checkout "+a.Branch+"\"")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	log.Printf("Repository initialized")
	log.Print("Finished fetching")

	_, err = os.Stat(a.RepoFolder + "start.sh")
	if os.IsExist(err) {
		log.Print("Running start.sh")
	} else {
		log.Print("No start.sh file found in repository folder")
	}

	return nil
}
