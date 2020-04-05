package store

// should we have an auto-genrated types file?

type Widget struct {
	WidgetID    string `json:"widget_id" db:"widget_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type Tag struct {
	TagID       string `json:"tag_id" db:"tag_id"`
	Name        string `json:"name" db:"name"`
	Slug        string `json:"slug" db:"slug"`
	Description string `json:"description" db:"description"`
}

type WidgetTag struct {
	WidgetTagID string `json:"widget_tag_id" db:"widget_tag_id"`
	WidgetID    string `json:"widget_id" db:"widget_id"`
	TagID       string `json:"tag_id" db:"tag_id"`
}

func (o *Widget) FieldList() []string {
	return []string{"widget_id", "name", "description"}
}

func (o *Widget) TableName() string {
	return "widget"
}

func (o *Tag) FieldList() []string {
	return []string{"tag_id", "name", "slug", "description"}
}

func (o *WidgetTag) FieldList() []string {
	return []string{"widget_tag_id", "widget_id", "tag_id"}
}
