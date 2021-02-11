package docbase

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}
