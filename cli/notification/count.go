/*
Copyright © 2021 Andy Zhou

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package notification

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count your notifications",
	Long:  `Output your notification count`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		count, err := GetNotificationCount(accessToken)
		cobra.CheckErr(err)
		fmt.Println(count)
	},
}

func GetNotificationCount(accessToken string) (count int, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u.URLFor("/api/v3/notification/all"), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = u.CheckStatusCode(resp, 200)
	if err != nil {
		return
	}
	var data []interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	count = len(data)
	return
}

func init() {
	notificationCmd.AddCommand(countCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
