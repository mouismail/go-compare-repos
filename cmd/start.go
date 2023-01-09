package cmd

import (
	"fmt"
	"log"
	"migrator/build"
	"migrator/pkg/options"
	"strconv"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command will start the migration",
	Long:  `Start command will start the app and generate migration stats report.`,
	Run: func(cmd *cobra.Command, args []string) {

		var hasProject = true

		source, _ := cmd.Flags().GetString("source")
		projectId, _ := cmd.Flags().GetString("projectId")
		githubOrg, _ := cmd.Flags().GetString("org")

		i, _ := strconv.Atoi(projectId)

		if projectId == "" {
			hasProject = false
		}

		s := build.SourceType{
			Source:     source,
			HasProject: hasProject,
		}

		o := options.NewOptions(s, githubOrg, i)

		err := o.Check()

		if err != nil {
			log.Fatal(err)
		}

		options.Generate(o)

		fmt.Printf("input flags values are %s %s %s", source, projectId, githubOrg)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("source", "s", "", "specify the repositories source, GitLab or BitBucket")
	startCmd.Flags().StringP("project", "p", "", "specify the project id for GitLab")
	startCmd.Flags().StringP("org", "o", "", "specify the GitHub organization")
}
