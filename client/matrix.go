package client

import (
	"hummingbard/gomatrix"
	"log"
)

func (c *Client) TempMatrixClient(userID, accessToken string) (*gomatrix.Client, error) {

	fu, us := c.IsFederated(userID)
	//port is only for my dev environment, this needs to go
	serverName := c.Config.Matrix.HomeserverURL

	//if federation user, we query homeserver at the /well-known endpoint
	//for full server path
	if fu {
		wk, err := WellKnown("https://" + us.ServerName)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		serverName = "https://" + wk.ServerName
	}

	matrix, err := gomatrix.NewClient(serverName, "", "")

	if accessToken != "" {
		matrix.SetCredentials(userID, accessToken)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return matrix, nil
}
