package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // anonymous type
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
