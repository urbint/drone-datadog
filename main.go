package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/urbint/drone-datadog/datadog"
	"github.com/urbint/drone-datadog/tags"
)

type Args struct {
	ApiKey  string   `envconfig:"dd_api_key"`
	Environ string   `envconfig:"dd_release_environment"`
	Version string   `envconfig:"dd_release_version"`
	Tags    []string `envconfig:"dd_event_tags"`
}

type DroneVars struct {
	BuildNumber   int    `envconfig:"build_number"`
	BuildFinished string `envconfig:"build_finished"`
	BuildStatus   string `envconfig:"build_status"`
	BuildLink     string `envconfig:"build_link"`
	CommitSha     string `envconfig:"commit_sha"`
	CommitBranch  string `envconfig:"commit_branch"`
	CommitAuthor  string `envconfig:"commit_author"`
	CommitLink    string `envconfig:"commit_link"`
	CommitMessage string `envconfig:"commit_message"`
	JobStarted    int64  `envconfig:"job_started"`
	Repo          string `envconfig:"build_link"`
	RepoLink      string `envconfig:"repo_link"`
	System        string
}

func main() {
	var (
		err   error
		vargs Args
		drone DroneVars
	)

	err = envconfig.Process("plugin", &vargs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = envconfig.Process("datadog", &drone)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := os.Stat(".tags"); err == nil {
		fmt.Println(".tags file found, parsing")
		vargs.Tags = append(vargs.Tags, tags.ParseFile(".tags")...)
	}

	// create the Datadog client
	client := datadog.NewClient(vargs.ApiKey)

	// generate the Datadog event
	msg := datadog.Event{
		Title:       "release-" + vargs.Environ + ": " + vargs.Version,
		Description: "Pushed " + vargs.Version + " to " + vargs.Environ,
		Tags:        vargs.Tags,
	}

	// sends the message
	if err := client.SendMessage(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
