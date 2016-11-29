package clarifai

type Query struct {
	QueryObject *QueryObject `json:"query"`
}

// QueryObject is a holder for query conditions (currently implemented only "AND").
type QueryObject struct {
	QueryAnds []*QueryFragment `json:"ands,omitempty"` // Collection of queries joined by an "AND" conditions.
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

type QueryInput struct {
	Data  *QueryData `json:"data,omitempty"`
	Input *Input     `json:"input,omitempty"` // used in reverse image search
}

type QueryData struct {
	Concepts []map[string]interface{} `json:"concepts,omitempty"`
	Image    *ImageData               `json:"image,omitempty"`
	Metadata interface{}              `json:"metadata,omitempty"`
}

// AddAndCondition adds a query part to a set of "AND" conditions.
func (q *Query) AddAndCondition(p *QueryFragment) {
	q.QueryObject.QueryAnds = append(q.QueryObject.QueryAnds, p)
}

type QueryConcept struct {
	Concept map[string]string
}

func NewQuery() *Query {
	return &Query{
		QueryObject: &QueryObject{},
	}
}

// SetMetadata adds metadata to q query input item ("input" -> "data" -> "metadata").
func (q *Input) SetMetadata(i interface{}) {
	if q.Data == nil {
		q.Data = &Image{}
	}
	q.Data.Metadata = i
}

func (q *QueryOutput) AddConcept(id string, value interface{}) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	m := make(map[string]interface{})
	m["id"] = id
	m["value"] = value

	q.Data.Concepts = append(q.Data.Concepts, m)
}

func (i *Input) AddConcept(id string, value interface{}) {

	if i.Data == nil {
		i.Data = &Image{}
	}

	m := make(map[string]interface{})
	m["id"] = id
	m["value"] = value

	i.Data.Concepts = append(i.Data.Concepts, m)
}

func (q *QueryOutput) SetConcepts(concepts map[string]interface{}) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	for n, v := range concepts {
		q.AddConcept(n, v)
	}

}

func (q *QueryOutput) AddConceptID(v string) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	m := make(map[string]interface{})
	m["id"] = v

	q.Data.Concepts = append(q.Data.Concepts, m)
}
