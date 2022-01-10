package Request

type CreateBucketRequest struct {
	Name string `json:"name" validate:"required"`
}
