package clarifai

type Query struct {
	QueryObject *QueryObject `json:"query"`
}

// QueryObject is a holder for query conditions (currently implemented only "AND").
type QueryObject struct {
	QueryAnds []*QueryFragment `json:"ands"` // Collection of queries joined by an "AND" conditions.
	// QueryOrs, etc.
}

// QueryPart is a part of the query that may be separated by conditions, x. "AND".
type QueryFragment struct {
	Output *QueryOutput `json:"output,omitempty"`
	Input  *QueryInput  `json:"input,omitempty"`
}

type QueryOutput struct {
	Data  *QueryData `json:"data,omitempty"`
	Input *Input     `json:"input,omitempty"` // used in reverse image search
}

type QueryInput struct {
	Data *QueryData `json:"data,omitempty"`
}

type QueryData struct {
	Concepts []map[string]string `json:"concepts,omitempty"`
	Metadata interface{}         `json:"metadata,omitempty"`
}

// AddAndCondition adds a query part to a set of "AND" conditions.
func (q *Query) AddAndCondition(p *QueryFragment) {
	q.QueryObject.QueryAnds = append(q.QueryObject.QueryAnds, p)
}

func (qp *QueryFragment) SetOutput(qo *QueryOutput) {
	qp.Output = qo
}

func (qp *QueryFragment) SetInput(qi *QueryInput) {
	qp.Input = qi
}

type QueryConcept struct {
	Concept map[string]string
}

func NewQuery() *Query {
	return &Query{
		QueryObject: &QueryObject{},
	}
}

func (q *QueryOutput) SetInput(i *Input) {
	q.Input = i
}

// SetMetadata adds metadata to q query input item ("input" -> "data" -> "metadata").
func (q *QueryInput) SetMetadata(i interface{}) {
	if q.Data == nil {
		q.Data = &QueryData{}
	}
	q.Data.Metadata = i
}

func (q *QueryOutput) AddConcept(v string) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	m := make(map[string]string)
	m["name"] = v

	q.Data.Concepts = append(q.Data.Concepts, m)
}

func (q *QueryInput) AddConcept(v string) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	m := make(map[string]string)
	m["name"] = v

	q.Data.Concepts = append(q.Data.Concepts, m)
}

func (q *QueryOutput) SetConcepts(concepts []string) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	for _, v := range concepts {
		q.AddConcept(v)
	}

}

func (q *QueryOutput) AddConceptID(v string) {

	if q.Data == nil {
		q.Data = &QueryData{}
	}

	m := make(map[string]string)
	m["id"] = v

	q.Data.Concepts = append(q.Data.Concepts, m)
}
