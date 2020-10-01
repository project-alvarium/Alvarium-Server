package handlers

import (
	"database-manager/configuration"
	"errors"
	"regexp"
	"database-manager/api"
	"database-manager/collections"
)

type InsightHandler struct {
}
var (
	UserCodeServerUrl string
)
const (
	/*valid url example: http://localhost:9496/api/handleMessage */
	urlRegex string = "^((https)|(http))(:\\/\\/((localhost)|(127.0.0.1)):([0-9]{1,5})(\\/))([a-zA-Z0-9]+([\\-\\.\\/]{1}[a-zA-Z0-9]+)*?)$"
)
func (insightHandler *InsightHandler) CreateInsight(data api.Data) (string, error) {
	collections.InsertOne(data.DataID)
	if match {
		config := configuration.Config
		UserCodeServerUrl = registerRequestModel.FullUrl
		orchestratorClient := clients.OrchestratorClient{Host: config.OrchestratorHost, Port: config.OrchestratorPort}
		err := orchestratorClient.Register()
		if err != nil {
			return "", errors.New("Registration Failed.")
		}
		return "successfully registered user code server: " + UserCodeServerUrl, nil
	}
	return "", errors.New("invalid user code server format " + registerRequestModel.FullUrl)
}

