/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package campaign

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this delete campaign",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpDeleteCampaign(CampaignID)
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&CampaignID, "id", "i", "", "delete the campaign")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.
	CampaignCmd.AddCommand(deleteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
