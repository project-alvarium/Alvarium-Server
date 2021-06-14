package handlers

import (
	"github.com/project-alvarium/Alvarium-Server/api/models"
	"github.com/project-alvarium/Alvarium-Server/collections"
)

// InsightHandler is astuct to export the struct
type InsightHandler struct {
}


// CreateInsight creates an insight
func CreateInsight(data models.Data) (string, error) {
	res, err := collections.InsertOne(data.DataID)
	if err != nil {
		return "", err
	}
	return res, nil
}

// GetInsight gets an insight from the db by id
func GetInsight(insightID string) (*collections.Insight, error) {
	res, err := collections.FindOne(insightID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
