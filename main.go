package main

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/tobiashort/clap-go"
	"github.com/tobiashort/th-utils/pkg/random"
	"github.com/tobiashort/utils-go/must"
)

type Args struct {
	Dir string `clap:"default-value='$HOME/.shelllog',description='Location where logs are stored'"`
}

func main() {
	args := Args{}
	clap.Parse(&args)

	dir := os.ExpandEnv(args.Dir)
	must.Do(os.MkdirAll(dir, 0700))

	width := 8
	seed := time.Now().UnixNano()
	alphabet := random.Uppercase + random.Lowercase + random.Numbers
	sessionID := random.String(alphabet, width, seed)

	logFile :=
		must.Do2(
			os.OpenFile(
				filepath.Join(dir, sessionID+".log"),
				os.O_CREATE|os.O_RDWR|os.O_APPEND,
				0600))
	defer logFile.Close()

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/bash"
	}

	cmd := exec.Command(shell)
	cmd.Stdin = os.Stdin
	cmd.Stdout = io.MultiWriter(os.Stdout, logFile)
	cmd.Stderr = io.MultiWriter(os.Stderr, logFile)
	must.Do(cmd.Run())
}
