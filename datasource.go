package sortable

type DataSource interface {
	GetSortableList(minSortIndex, maxSortIndex int) ([]Element, error)

	UpateSortableList(elements []Element) error
}
