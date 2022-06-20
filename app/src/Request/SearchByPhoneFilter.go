package Request

type SearchByPhoneFilter struct {
	Phone string `json:"phone" binding:"required"`
}
