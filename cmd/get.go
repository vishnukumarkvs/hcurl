/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Make a GET request to the specified URL",
	Long:  `Make a GET request to the specified URL`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Please provide the url")
			return
		}
		url := args[0]

		var wg sync.WaitGroup
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					fmt.Println("Error creating request:", err)
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

				fmt.Printf("Response status: %s\nResponse Body: %s\n", response.Status, string(respBody))
			}()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
