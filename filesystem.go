package actions

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/google/go-github/github"
)

const (
	// GithubHome is a directory containing user-related data.
	GithubHome = "/github/home"
	// GithubWorkspace is the working directory of the DockerContainer.
	GithubWorkspace = "/github/workspace"
	// GithubPostEventJSONPATH specifies the path of the POST response of the webhook event.
	GithubPostEventJSONPATH = "/github/workflow/event.jso"
)

var (
	cachedPushEvent   *github.PushEvent
	cachedPushEventMu sync.Mutex
)

// LoadPushEvent returns the POST response of the webhook event that triggered the workflow.
// The result is globally cached. It's no problem because this cache is cleared at ends of actions.
func LoadPushEvent() (*github.PushEvent, error) {
	cachedPushEventMu.Lock()
	defer cachedPushEventMu.Unlock()
	if cachedPushEvent != nil {
		return cachedPushEvent, nil
	}

	f, err := os.Open(GithubPostEventJSONPATH)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cachedPushEvent = new(github.PushEvent)
	err = json.NewDecoder(f).Decode(cachedPushEvent)
	return cachedPushEvent, err
}
