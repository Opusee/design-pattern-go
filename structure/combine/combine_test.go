package combine

import "testing"

func TestCombine(t *testing.T) {
	driverD := NewFolder("D盘")

	doc := NewFolder("文档")
	doc.Add(NewFile("简历.doc"))
	doc.Add(NewFile("项目介绍.ppt"))

	driverD.Add(doc)

	music := NewFolder("音乐")

	jay := NewFolder("周杰伦")
	jay.Add(NewFile("双节棍.mp3"))
	jay.Add(NewFile("告白气球.mp3"))
	jay.Add(NewFile("听妈妈的话.mp3"))

	jack := NewFolder("张学友")
	jack.Add(NewFile("吻别.mp3"))
	jack.Add(NewFile("一千个伤心的理由.mp3"))

	music.Add(jay)
	music.Add(jack)

	driverD.Add(music)

	driverD.Ls()
}
