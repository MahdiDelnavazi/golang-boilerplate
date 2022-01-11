package Request

type FilterBucketRequest struct {
	Name       string `json:"name"`
	Active     string `json:"active"`
	PageSize   int    `json:"pageSize"`
	PageNumber int    `json:"pageNumber"`
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
}
