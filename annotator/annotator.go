package annotator
import (
	"database-manager/iota"
	"database-manager/api/models"
);
type Annotator struct {
	sub iota.Subscriber
	count int
	ids []string
}
 
func NewAnnotator(sub iota.Subscriber, ids []string) Annotator {
	return Annotator{sub, 0, ids}
}

func StoreAnnotation(sub iota.Subscriber, data models.Data) {
	an := iota.Annotation{"id", "id", 2, "2", "2", 1.0,data.DataID, data.Content, 1}
	annotationMessage := iota.NewAnnotation("000", an)
	sub.SendMessage(annotationMessage)
}