# go-crud-todo


type Todo struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"` // bson just for mongoDb to tell it is id
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
