package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

//flags used by more than 1 command
var (
	DebugFlag                  bool
	WatchFlag                  bool
	DeploymentOutputFlag       bool
	FollowFlag                 bool
	Name                       string
	ApplicationName            string
	ProjectName                string
	OrganizationName           string
	BranchName                 string
	ShowCredentials            bool
	OutputEnvironmentVariables bool
	Tail                       int
	ConfigurationDirectoryRoot string
)

func hasFlagChanged(cmd *cobra.Command) bool {
	flagChanged := false

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed && flag.Name != "watch" && flag.Name != "deployment-output" && flag.Name != "follow" && flag.Name != "tail" &&
			flag.Name != "credentials" && flag.Name != "debug" && flag.Name != "dotenv" {
			flagChanged = true
		}
	})

	return flagChanged
}
