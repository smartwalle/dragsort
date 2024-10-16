package main

import (
	"fmt"
	"github.com/smartwalle/sortable"
	"sort"
)

func main() {
	var ds = &DataSource{}
	for i := 1; i <= 10; i++ {
		ds.users = append(ds.users, &User{
			ID:        int64(i),
			SortIndex: i,
		})
	}

	sortable.Sort(ds, ds.users[2], ds.users[1])
	sortable.Sort(ds, ds.users[2], ds.users[3])
	sortable.Sort(ds, ds.users[2], ds.users[4])
}

type User struct {
	ID        int64
	SortIndex int
}

func (u *User) String() string {
	return fmt.Sprintf("[%d-%d]", u.ID, u.SortIndex)
}

func (u *User) GetUniqueID() int64 {
	return u.ID
}

func (u *User) GetSortIndex() int {
	return u.SortIndex
}

func (u *User) UpdateSortIndex(sortIndex int) {
	u.SortIndex = sortIndex
}

type DataSource struct {
	users []*User
}

func (ds *DataSource) GetSortableList(minSortIndex, maxSortIndex int) ([]sortable.Element, error) {
	var elements = make([]sortable.Element, 0, len(ds.users))
	for _, u := range ds.users {
		if u.SortIndex <= maxSortIndex && u.SortIndex >= minSortIndex {
			elements = append(elements, u)
		}
	}
	return elements, nil
}

func (ds *DataSource) UpateSortableList(elements []sortable.Element) error {
	sort.SliceStable(ds.users, func(i, j int) bool {
		if ds.users[i].SortIndex < ds.users[j].SortIndex {
			return true
		}
		return false
	})
	fmt.Println("更新后的顺序：", ds.users)
	return nil
}
