package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	dr := NewDrivingRecorder()
	// 假设记录了12条视频
	for i := 0; i < 12; i++ {
		dr.Append(fmt.Sprintf("视频_%d", i))
	}
	//用户要获取交通事故视频，定义事故列表。
	accidents := make([]string, 0)
	//用户拿到迭代器
	it := dr.Iterator()
	for it.HasNext() { //如果还有下一条则继续迭代
		video := it.Next()
		println(video)
		//用户翻看视频发现10和8可作为证据。
		if video == "视频_10" || video == "视频_8" {
			accidents = append(accidents, video)
		}
	}
	//拿到两个视频集accidents交给交警查看。
	fmt.Printf("事故证据：%v\n", accidents)
}
