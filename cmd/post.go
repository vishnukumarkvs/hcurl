/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var body string

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post [URL]",
	Short: "Make a POST request to the specified URL",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		var requestBody []byte
		if strings.HasPrefix(body, "@") {
			filePath := strings.TrimPrefix(body, "@")
			fmt.Println("Path of file: ", filePath)
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Unable to read the file: ", err)
				return
			}
			requestBody = fileContent
		} else {
			requestBody = []byte(body)
		}
		var wg sync.WaitGroup
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				response, err := http.Post(url, "application/json", strings.NewReader(string(requestBody)))
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				defer response.Body.Close()

				respBody, err := io.ReadAll(response.Body)
				if err != nil {
					fmt.Println("Error reading response: ", err)
					return
				}
				fmt.Println("Response status: ", response.Status)
				fmt.Println("Response Body", string(respBody))
			}()
		}
		wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// postCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// postCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	postCmd.Flags().StringVarP(&body, "request-body", "b", "", "Pass request body of type json either in single quotes or pass the json file path starting with @")
}
