// Copyright 2024 Humanitec
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHelp(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"--help"})
	assert.NoError(t, err)
	assert.Equal(t, `SCORE is a specification for defining environment agnostic configuration for cloud based workloads.
This tool produces a docker-compose configuration file from the SCORE specification.
Complete documentation is available at https://score.dev

Usage:
  score-ecs [command]

Available Commands:
  check-version Assert that the version of score-ecs matches the required constraint
  completion    Generate the autocompletion script for the specified shell
  generate      Convert one or more Score files into a Docker compose manifest
  help          Help about any command
  init          Initialise a new score-ecs project with local state directory and score file
  provisioners  Subcommands related to provisioners
  resources     Subcommands related to provisioned resources

Flags:
  -h, --help            help for score-ecs
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times
      --version         version for score-ecs

Use "score-ecs [command] --help" for more information about a command.
`, stdout)
	assert.Equal(t, "", stderr)
}

func TestRootVersion(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"--version"})
	assert.NoError(t, err)
	pattern := regexp.MustCompile(`^score-ecs 0.0.0 \(build: \S+, sha: \S+\)\n$`)
	assert.Truef(t, pattern.MatchString(stdout), "%s does not match: '%s'", pattern.String(), stdout)
	assert.Equal(t, "", stderr)
}

func TestRootCompletion(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion"})
	assert.NoError(t, err)
	assert.Equal(t, `Generate the autocompletion script for score-ecs for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  score-ecs completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Global Flags:
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times

Use "score-ecs completion [command] --help" for more information about a command.
`, stdout)
	assert.Equal(t, "", stderr)

	stdout2, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"help", "completion"})
	assert.NoError(t, err)
	assert.Equal(t, stdout, stdout2)
	assert.Equal(t, "", stderr)
}

func TestRootCompletionBashHelp(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "bash", "--help"})
	assert.NoError(t, err)
	assert.Equal(t, `Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(score-ecs completion bash)

To load completions for every new session, execute once:

#### Linux:

	score-ecs completion bash > /etc/bash_completion.d/score-ecs

#### macOS:

	score-ecs completion bash > $(brew --prefix)/etc/bash_completion.d/score-ecs

You will need to start a new shell for this setup to take effect.

Usage:
  score-ecs completion bash

Flags:
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions

Global Flags:
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times
`, stdout)
	assert.Equal(t, "", stderr)
}

func TestRootCompletionBash(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "bash"})
	assert.NoError(t, err)
	assert.Contains(t, stdout, "# bash completion V2 for score-ecs")
	assert.Equal(t, "", stderr)
}

func TestRootCompletionFishHelp(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "fish", "--help"})
	assert.NoError(t, err)
	assert.Equal(t, `Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	score-ecs completion fish | source

To load completions for every new session, execute once:

	score-ecs completion fish > ~/.config/fish/completions/score-ecs.fish

You will need to start a new shell for this setup to take effect.

Usage:
  score-ecs completion fish [flags]

Flags:
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions

Global Flags:
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times
`, stdout)
	assert.Equal(t, "", stderr)
}

func TestRootCompletionFish(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "fish"})
	assert.NoError(t, err)
	assert.Contains(t, stdout, "# fish completion for score-ecs")
	assert.Equal(t, "", stderr)
}

func TestRootCompletionZshHelp(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "zsh", "--help"})
	assert.NoError(t, err)
	assert.Equal(t, `Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(score-ecs completion zsh)

To load completions for every new session, execute once:

#### Linux:

	score-ecs completion zsh > "${fpath[1]}/_score-ecs"

#### macOS:

	score-ecs completion zsh > $(brew --prefix)/share/zsh/site-functions/_score-ecs

You will need to start a new shell for this setup to take effect.

Usage:
  score-ecs completion zsh [flags]

Flags:
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions

Global Flags:
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times
`, stdout)
	assert.Equal(t, "", stderr)
}

func TestRootCompletionZsh(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "zsh"})
	assert.NoError(t, err)
	assert.Contains(t, stdout, "# zsh completion for score-ecs")
	assert.Equal(t, "", stderr)
}

func TestRootCompletionPowershellHelp(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "powershell", "--help"})
	assert.NoError(t, err)
	assert.Equal(t, `Generate the autocompletion script for powershell.

To load completions in your current shell session:

	score-ecs completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

Usage:
  score-ecs completion powershell [flags]

Flags:
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions

Global Flags:
      --quiet           Mute any logging output
  -v, --verbose count   Increase log verbosity and detail by specifying this flag one or more times
`, stdout)
	assert.Equal(t, "", stderr)
}

func TestRootCompletionPowershell(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"completion", "powershell"})
	assert.NoError(t, err)
	assert.Contains(t, stdout, "# powershell completion for score-ecs")
	assert.Equal(t, "", stderr)
}

func TestRootUnknown(t *testing.T) {
	stdout, stderr, err := executeAndResetCommand(context.Background(), rootCmd, []string{"unknown"})
	assert.EqualError(t, err, "unknown command \"unknown\" for \"score-ecs\"")
	assert.Equal(t, "", stdout)
	assert.Equal(t, "", stderr)
}
