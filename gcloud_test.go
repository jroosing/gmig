package main

import (
	"fmt"
	"testing"
)

func TestGcloudConfigSetProject(t *testing.T) {
	cc := new(commandCapturer)
	old := runCommand
	runCommand = cc.runCommand
	defer func() { runCommand = old }()
	cfg := Config{Project: "p", Region: "r", Zone: "z"}
	if err := gcloudConfigSetProject(cfg); err != nil {
		t.Fatal(err)
	}
	if got, want := len(cc.args), 3; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := fmt.Sprint(cc.args[0]), "[gcloud config set core/project p]"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := fmt.Sprint(cc.args[1]), "[gcloud config set compute/region r]"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := fmt.Sprint(cc.args[2]), "[gcloud config set compute/zone z]"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
