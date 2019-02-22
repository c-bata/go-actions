package actions

import (
	"os"
	"path/filepath"
)

const (
	prefixGithubRef = "refs/heads/"
)

// GetHome returns the path to the GitHub home directory used to store user data. Value: /github/home.
func GetHome() string {
	return os.Getenv("HOME")
}

// GetGithubWorkflow returns the name of the workflow.
func GetGithubWorkflow() string {
	return os.Getenv("GITHUB_WORKFLOW")
}

// GetGithubAction returns the name of action.
func GetGithubAction() string {
	return os.Getenv("GITHUB_ACTION")
}

// GetGithubActor returns the name of the person or app that initiated the workflow.
func GetGithubActor() string {
	return os.Getenv("GITHUB_ACTOR")
}

// GetGithubRepository returns the owner and repository name.
func GetGithubRepository() string {
	return os.Getenv("GITHUB_REPOSITORY")
}

// GetGithubEventName returns the webhook name of the event that triggered the workflow.
func GetGithubEventName() string {
	return os.Getenv("GITHUB_EVENT_NAME")
}

// GetGithubEventPath returns the path to a file that contains the payload of the event that triggered the workflow.
func GetGithubEventPath() string {
	return os.Getenv("GITHUB_EVENT_PATH")
}

// GetGithubWorkspace returns the GitHub workspace path.
func GetGithubWorkspace() string {
	return os.Getenv("GITHUB_WORKSPACE")
}

// GetGithubSHA returns the commit SHA that triggered the workflow.
func GetGithubSHA() string {
	return os.Getenv("GITHUB_SHA")
}

// GetDirContainsCopyOfRepository the directory path contains a copy of the repository.
func GetDirContainsCopyOfRepository() string {
	return filepath.Join(GithubWorkspace, GetGithubSHA())
}

// GetGithubRef returns the branch or tag ref that triggered the workflow.
func GetGithubRef() string {
	return os.Getenv("GITHUB_REF")
}

// GetTagOrBranchName returns the branch or tag name that triggered the workflow.
func GetTagOrBranchName() string {
	ref := GetGithubRef()
	if ref == "" {
		return ""
	}
	return ref[len(prefixGithubRef):]
}

// GetGithubToken returns a GitHub App installation token scoped to the repository containing the workflow file.
func GetGithubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}
