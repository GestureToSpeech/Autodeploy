package main

import (
	"log"
	"net/http"

	"github.com/pelletier/go-toml"
)

func main() {
	cfg, err := toml.LoadFile("config.tml")
	catch(err)

	repo, ok := cfg.Get("app.repo").(string)
	if !ok {
		log.Fatal("app.repo not defined")
	}

	branch, ok := cfg.Get("app.branch").(string)
	if !ok {
		log.Fatal("app.branch not defined")
	}

	key, ok := cfg.Get("hook.key").(string)
	if !ok {
		log.Fatal("hook.key not defined")
	}

	app := &App{
		Repo:   repo,
		Branch: branch,
	}

	http.Handle("/hook", NewHookHandler(&HookOptions{
		App:    app,
		Secret: key,
	}))

	addr, ok := cfg.Get("core.addr").(string)
	if !ok {
		log.Fatal("core.addr not defined")
	}

	err = http.ListenAndServe(addr, nil)
	catch(err)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
