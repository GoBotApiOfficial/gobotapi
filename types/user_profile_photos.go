package types

type UserProfilePhotos struct {
	Photos [][]PhotoSize `json:"photos,omitempty"`
	TotalCount int `json:"total_count"`
}
