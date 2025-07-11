// Code AutoGenerated; DO NOT EDIT.

package types

// InputChecklist Describes a checklist to create.
type InputChecklist struct {
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks,omitempty"`
	Title                    string               `json:"title"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
}
