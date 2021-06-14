package handlers

import (
	"github.com/project-alvarium/Alvarium-Server/annotator"
	"github.com/project-alvarium/Alvarium-Server/api/models"
	"github.com/project-alvarium/Alvarium-Server/subscriber"
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