package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command will start the migration",
	Long:  `Start command will start the app and generate migration stats report.`,
	Run: func(cmd *cobra.Command, args []string) {

		source, _ := cmd.Flags().GetString("source")
		projectId, _ := cmd.Flags().GetString("projectId")
		githubOrg, _ := cmd.Flags().GetString("org")

		fmt.Printf("input flags values are %s %s %s", source, projectId, githubOrg)
		//task := utils.Task{
		//	Title:       taskTitle,
		//	Description: taskDescription,
		//}
		//fmt.Printf("Creating task %+v\n", task)
		//
		//// Storing task in backend calling my-todos REST API
		//resp := utils.WriteAPI(task)
		//fmt.Println("Task created with ID:", resp.TaskID)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("source", "s", "", "specify the repositories source, GitLab or BitBucket")
	startCmd.Flags().StringP("project", "p", "", "specify the project id for GitLab")
	startCmd.Flags().StringP("org", "o", "", "specify the GitHub organization")
}
