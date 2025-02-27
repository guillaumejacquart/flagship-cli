/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"

	httprequest "github.com/Chadiii/flagship/utils/httpRequest"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this ldelete user",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		httprequest.HttpDeleteUsers(UserEmail)
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&UserEmail, "email", "e", "", "the email")

	if err := deleteCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	UserCmd.AddCommand(deleteCmd)
}
