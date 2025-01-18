package iterator

/**
迭代器模式
*/

type DrivingRecorder struct {
	index   int        // 当前记录位置
	records [10]string // 假设只能记录10条视频
}

func NewDrivingRecorder() *DrivingRecorder {
	return &DrivingRecorder{
		index: -1,
	}
}

func (d *DrivingRecorder) Append(record string) {
	if d.index == 9 { // 循环覆盖
		d.index = 0
	} else {
		d.index++
	}
	d.records[d.index] = record
}

func (d *DrivingRecorder) Iterator() Iterator[string] {
	iterator := Iterator[string]{}
	cursor := d.index // 迭代器游标，不染指原始游标。
	loopCount := 0
	iterator.Next = func() string {
		i := cursor // 记录即将返回的游标位置
		if cursor == 0 {
			cursor = 9
		} else {
			cursor--
		}
		loopCount++
		return d.records[i]
	}
	iterator.HasNext = func() bool {
		return loopCount < len(d.records)
	}
	return iterator
}

type Iterator[T any] struct {
	Next    func() T    //返回下一个元素
	HasNext func() bool //是否还有下一个元素
}
