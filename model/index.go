package model

// IndexQuery contains query parameters to request video information
type IndexQuery struct {
	Q  string
	Vt string
}

type IndexResponse struct {
	Status string
	Mess   string
	Q      string
	P      string
	Vid    string
	T      int
	A      string
	Kc     string
}

var (
	// baseURL contains base URL for service
	baseURL = "https://yt1s.com"

	// path.
	indexPath = "/api/ajaxSearch/index"
)

func GetBaseURL() string {
	return baseURL
}

func GetIndexPath() string {
	return indexPath
}
