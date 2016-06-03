// This file is part of the Plan-B-ot package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package bot

import (
	"bytes"
	"errors"
	"net/http"
)

// sendToSlack sends a message to slack using an incoming web hook
func sendToSlack(title, message, color string) error {
	attachment := `{
		"attachments": [
			{
				"fallback": "` + title + `: ` + message + `",
				"color": "` + color + `",
				"fields": [
					{
						"title": "` + title + `",
						"value": "` + message + `",
						"short": false
					}
				]
			}
		]
	}`

	response, err := http.Post(
		Config.WebhookURL,
		"POST",
		bytes.NewBufferString(attachment),
	)

	if response.StatusCode != http.StatusOK {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(response.Body)
		err = errors.New(buffer.String())
	}

	return err
}
