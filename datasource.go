package dragsort

type DataSource interface {
	GetSortableList(minSortIndex, maxSortIndex int) ([]Sortable, error)

	UpateSortableList(elements []Sortable) error
}
