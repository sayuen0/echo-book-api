package request_model

// request models
type UserGetRequest struct {
	Id int64 `json:"id" param:"id" validate:"required"`
}

type UserCreateRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UserUpdateRequest struct {
	id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UserDeleteRequest struct {
	Id int64 `json:"id" validate:"required"`
}

