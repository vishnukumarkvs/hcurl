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
		if len(args) < 1 {
			fmt.Println("Please provide the url")
			return
		}

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

				req, err := http.NewRequest("POST", url, strings.NewReader(string(requestBody)))

				if err != nil {
					fmt.Println("Error creating request: ", err)
					return
				}

				req.Header.Set("Content-Type", "application/json")

				headers := strings.Split(headers, ",")
				for _, h := range headers {
					h = strings.TrimSpace(h)
					parts := strings.SplitN(h, ":", 2)
					if len(parts) == 2 {
						key := strings.TrimSpace(parts[0])
						value := strings.TrimSpace(parts[1])
						req.Header.Set(key, value)
					}
				}

				client := &http.Client{}

				response, err := client.Do(req)
				if err != nil {
					fmt.Println("Not able to send request: ", err)
					return
				}

				defer response.Body.Close()

				respBody, err := io.ReadAll(response.Body)

				fmt.Printf("Response status: %s\n Response Body: %s\n", response.Status, string(respBody))
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
