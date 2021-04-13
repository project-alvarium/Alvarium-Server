package iota

import (
	"encoding/json"
	"fmt"
)
type Annotation struct {
	Iss string  `json:"iss" bson:"iss"`
	Sub string  `json:"sub" bson:"sub"`
	Iat int64   `json:"iat" bson:"iat"`
	Jti string  `json:"jti" bson:"jti"`
	Ann string  `json:"ann" bson:"ann"`
	Avl float64 `json:"avl" bson:"avl"`
	DataId string `json:"dataId" bson:"dataId"`
	Content string `json:"content" bson:"content"`
	Score float64 `json:"score" bson:"score"`
	
}
type TangleMessage struct {
	message string
}

func NewReading(sensorId string, readingId string, data string) TangleMessage {
	message := "{ \"sensor_id\": \"" + sensorId + "\", \"reading_id\": \"" + readingId +
		"\", \"data\": \"" + data + "\" }"
	return TangleMessage{ message }
}

func NewAnnotation(readingId string, ann Annotation) TangleMessage {
	message := "{ \"reading_id\": \"" + readingId + "\", \"annotation\": " + AnnotationToString(ann) + " }"
	return TangleMessage{ message }
}

func AnnotationToString(ann Annotation) string {
	return "{ \"header\": " + Header("RS256", "JWT") +
		", \"payload\": " + Payload(ann) +
		", \"signature\": " + Signature("A signature") + " }"
}

func Header(alg string, typ string) string {
	return "{ \"alg\": \"" + alg + "\", \"typ\": \"" + typ + "\" }"
}

func Signature(sig string) string {
	return "\"" + sig + "\""
}

func Payload(ann Annotation) string {
	j, err := json.Marshal(ann)
	if err != nil {
		fmt.Println("Error marshalling annotation json")
	}
	return string(j)
}
