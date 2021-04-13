package handlers

import (
	"database-manager/api/models"
	"database-manager/subscriber"
	 "database-manager/annotator"
)

func CreateContent(data models.Data) (string, error) {
	data = models.Data{"123","content"}
	annotator.StoreAnnotation(subscriber.AnnSubscriber, data)
	return "ok", nil
}
// TODO : return data
func GetContent(insightID string) (string, error) {
	return "ok", nil
}