package dragsort

import "math"

func Sort(dataSource DataSource, sourceID int64, targetID int64) error {
	if dataSource == nil || sourceID < 1 || targetID < 1 || sourceID == targetID {
		return nil
	}

	source, err := dataSource.GetSortable(sourceID)
	if err != nil {
		return err
	}
	if source == nil || source.GetUniqueID() < 1 {
		return nil
	}

	target, err := dataSource.GetSortable(targetID)
	if err != nil {
		return err
	}
	if target == nil || target.GetUniqueID() < 1 {
		return nil
	}

	var sourceSortIndex = source.GetSortIndex()
	var targetSortIndex = target.GetSortIndex()

	// 取出 source、target 及两者之间的所有数据
	var minSortIndex = int(math.Min(float64(sourceSortIndex), float64(targetSortIndex)))
	var maxSortIndex = int(math.Max(float64(sourceSortIndex), float64(targetSortIndex)))
	elements, err := dataSource.GetSortableList(minSortIndex, maxSortIndex)
	if err != nil {
		return err
	}

	for _, ele := range elements {
		if ele.GetUniqueID() == source.GetUniqueID() {
			ele.UpdateSortIndex(targetSortIndex)
			continue
		}

		if sourceSortIndex > targetSortIndex {
			// 往前移动，其它元素的排序 +1
			ele.UpdateSortIndex(ele.GetSortIndex() + 1)
		} else {
			// 往后移动，其实元素的排序 -1
			ele.UpdateSortIndex(ele.GetSortIndex() - 1)
		}
	}

	err = dataSource.UpateSortableList(elements)
	if err != nil {
		return err
	}

	return nil
}
