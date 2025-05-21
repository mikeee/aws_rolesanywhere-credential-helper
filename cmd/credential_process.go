package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	helper "github.com/mikeee/aws_rolesanywhere-credential-helper/aws_signing_helper"
	"github.com/spf13/cobra"
)

func init() {
	initCredentialsSubCommand(credentialProcessCmd)
}

var credentialProcessCmd = &cobra.Command{
	Use:   "credential-process [flags]",
	Short: "Retrieve AWS credentials in the appropriate format for external credential processes",
	Long: `To retrieve AWS credentials in the appropriate format for external
credential processes, as determined by the SDK/CLI. More information can be
found at: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html`,
	Run: func(cmd *cobra.Command, args []string) {
		err := PopulateCredentialsOptions()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		helper.Debug = credentialsOptions.Debug

		signer, signingAlgorithm, err := helper.GetSigner(&credentialsOptions)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		defer signer.Close()
		credentialProcessOutput, err := helper.GenerateCredentials(&credentialsOptions, signer, signingAlgorithm)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		buf, _ := json.Marshal(credentialProcessOutput)
		fmt.Print(string(buf[:]))
	},
}
