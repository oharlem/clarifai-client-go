package clarifai

type SearchService struct {
	InputObject  *InputObject
	OutputObject *OutputObject
	QueryObject  *QueryObject
	Session      *Session
	Page         int
	PerPage      int
}

type Hit struct {
	Score float64 `json:"score"`
	Input *Input  `json:"input,omitempty"`
}

type SearchResponse struct {
	Status ServiceStatus `json:"status"`
	Hits   []Hit         `json:"hits"`
}

func (svc *SearchService) GetPage() int {
	return svc.Page
}

func (svc *SearchService) GetPerPage() int {
	return svc.PerPage
}

func (svc *SearchService) GetSession() *Session {
	return svc.Session
}

func (svc *SearchService) GetInputObject() *InputObject {
	return svc.InputObject
}

func NewSearchService(s *Session) *SearchService {
	return &SearchService{
		Session:     s,
		InputObject: &InputObject{},
	}
}

//
//
// API
//
//

// Call sends a search query request to Clarifai search service.
func (svc *SearchService) Call(q *Query) (*SearchResponse, error) {
	out := &SearchResponse{}
	PP(q)
	payload, err := PrepPayload(q)
	if err != nil {
		return out, err
	}

	err = PostCall(svc, GetURI(svc, "searches"), payload, out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// AddImagesToIndex sends images to a search index.
func (svc *SearchService) AddImagesToIndex() (*AddImagesResponse, error) {
	out := &AddImagesResponse{}

	payload, err := GetPayload(svc)
	if err != nil {
		return out, err
	}

	err = PostCall(svc, GetURI(svc, "inputs"), payload, out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// AddInput adds input to the predict request.
// Limited to 128 inputs per call as per https://developer-preview.clarifai.com/guide/inputs#inputs.
func (svc *SearchService) AddInput(i *Input) error {
	if svc.InputObject.GetInputsQty() >= InputLimit {
		return ErrInputLimitReached
	}
	svc.InputObject.AddInput(i)
	return nil
}

func (svc *SearchService) AddOutput(i *Input) error {
	if svc.InputObject.GetInputsQty() >= InputLimit {
		return ErrInputLimitReached
	}
	svc.InputObject.AddInput(i)
	return nil
}

func (svc *SearchService) WithPagination(page, perPage int) {
	svc.Page = page
	svc.PerPage = perPage
}

// AddImagesResponse is a response, specific to procedure of adding images to a search index.
type AddImagesResponse struct {
	Status *ServiceStatus `json:"status,omitempty"`
	Inputs []*Input       `json:"inputs,omitempty"`
}

// SearchByPredictedConcepts searches by predictions received for previously supplied images.
func (svc *SearchService) SearchByPredictedConcepts(c []string) (*SearchResponse, error) {

	q := NewQuery()

	q.AddOutputConcepts(c)

	return svc.Call(q)
}

func (q *Query) AddOutputConcepts(c []string) {

	for _, v := range c {
		qo := QueryOutput{}
		qo.AddConcept(v)

		qf := QueryFragment{}
		qf.SetOutput(&qo)

		q.AddAndCondition(&qf)
	}
}

// SearchByUserSuppliedConcepts searches by concepts that were previously added along with images.
func (svc *SearchService) SearchByUserSuppliedConcepts(c []string) (*SearchResponse, error) {

	q := NewQuery()

	q.AddInputConcepts(c)

	return svc.Call(q)
}

func (q *Query) AddInputConcepts(c []string) {

	for _, v := range c {
		qi := QueryInput{}
		qi.AddConcept(v)

		qf := QueryFragment{}
		qf.SetInput(&qi)

		q.AddAndCondition(&qf)
	}
}

// ReverseImageSearch search by variable number of image URLs or bytes. Returns images from your search index that are visually similar to the ones provided.
func (svc *SearchService) ReverseImageSearch(inputs ...*Input) (*SearchResponse, error) {

	q := NewQuery()

	for _, i := range inputs {
		qo := QueryOutput{}
		qo.SetInput(i)

		qf := QueryFragment{}
		qf.SetOutput(&qo)

		q.AddAndCondition(&qf)
	}

	return svc.Call(q)
}

// SearchByCustomMetadata searches by metadata, that was previously added with images.
func (svc *SearchService) SearchByCustomMetadata(m interface{}) (*SearchResponse, error) {

	q := NewQuery()

	qo := QueryInput{}
	qo.SetMetadata(m)

	qf := QueryFragment{}
	qf.SetInput(&qo)

	q.AddAndCondition(&qf)

	return svc.Call(q)
}
