package Request

type UpdateBucketRequest struct {
	Name string `json:"name" validate:"required"`
}
