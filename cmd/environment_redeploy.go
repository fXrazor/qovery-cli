package cmd

import (
	"context"
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"time"

	"github.com/qovery/qovery-cli/utils"
	"github.com/qovery/qovery-client-go"
	"github.com/spf13/cobra"
)

var environmentRedeployCmd = &cobra.Command{
	Use:   "redeploy",
	Short: "Redeploy an environment",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Capture(cmd)

		tokenType, token, err := utils.GetAccessToken()
		if err != nil {
			utils.PrintlnError(err)
			os.Exit(1)
			panic("unreachable") // staticcheck false positive: https://staticcheck.io/docs/checks#SA5011
		}

		client := utils.GetQoveryClient(tokenType, token)
		_, _, envId, err := getOrganizationProjectEnvironmentContextResourcesIds(client)

		if err != nil {
			utils.PrintlnError(err)
			os.Exit(1)
			panic("unreachable") // staticcheck false positive: https://staticcheck.io/docs/checks#SA5011
		}

		// wait until service is ready
		for {
			if utils.IsEnvironmentInATerminalState(envId, client) {
				break
			}

			utils.Println(fmt.Sprintf("Waiting for environment %s to be ready..", pterm.FgBlue.Sprintf(envId)))
			time.Sleep(5 * time.Second)
		}

		_, _, err = client.EnvironmentActionsAPI.RedeployEnvironment(context.Background(), envId).Execute()

		if err != nil {
			utils.PrintlnError(err)
			os.Exit(1)
			panic("unreachable") // staticcheck false positive: https://staticcheck.io/docs/checks#SA5011
		}

		utils.Println("Environment is redeploying!")

		if watchFlag {
			utils.WatchEnvironment(envId, qovery.STATEENUM_DEPLOYED, client)
		}
	},
}

func init() {
	environmentCmd.AddCommand(environmentRedeployCmd)
	environmentRedeployCmd.Flags().StringVarP(&organizationName, "organization", "", "", "Organization Name")
	environmentRedeployCmd.Flags().StringVarP(&projectName, "project", "", "", "Project Name")
	environmentRedeployCmd.Flags().StringVarP(&environmentName, "environment", "", "", "Environment Name")
	environmentRedeployCmd.Flags().BoolVarP(&watchFlag, "watch", "w", false, "Watch environment status until it's ready or an error occurs")
}
