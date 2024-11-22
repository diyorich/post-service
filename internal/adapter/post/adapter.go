package post

type Adapter struct {
	baseURL string
}

func NewPostAdapter(baseURL string) *Adapter {
	return &Adapter{
		baseURL: baseURL,
	}
}
