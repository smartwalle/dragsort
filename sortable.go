package dragsort

type Sortable interface {
	GetUniqueID() int64

	GetSortIndex() int

	UpdateSortIndex(sortIndex int)
}
