package main

var (
	// The full version string
	Version = "2.0.0.0"

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse --short=8 HEAD)"
	GitCommit string
)

func GetVersion() string {
	if GitCommit != "" {
		return Version + "-" + GitCommit
	}
	return Version
}