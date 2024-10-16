package dragsort

type DataSource interface {
	GetSortable(uniqueID int64) (Sortable, error)

	GetSortableList(minSortIndex, maxSortIndex int) ([]Sortable, error)

	UpateSortableList(elements []Sortable) error
}
