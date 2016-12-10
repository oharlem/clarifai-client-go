package clarifai

// QueryObject is a holder for query conditions (currently implemented only "AND").
type QueryObject struct {
	Ands []*QueryFragment `json:"ands,omitempty"` // Collection of queries joined by an "AND" conditions.
	// QueryOrs, etc.
}

// QueryFragment is a self-contained part of a conditional clause.
type QueryFragment struct {
	Output *QueryOutput `json:"output,omitempty"`
	Input  *Input       `json:"input,omitempty"`
}

type QueryOutput struct {
	Data  *QueryData `json:"data,omitempty"`
	Input *Input     `json:"input,omitempty"` // used in reverse image search
}

type QueryData struct {
	Concepts []map[string]interface{} `json:"concepts,omitempty"`
	Image    *ImageData               `json:"image,omitempty"`
	Metadata interface{}              `json:"metadata,omitempty"`
}

func (q *QueryOutput) AddConcept(id string, value interface{}) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	q.Data.Concepts = append(q.Data.Concepts, map[string]interface{}{
		"name":  id,
		"value": value,
	})
}
