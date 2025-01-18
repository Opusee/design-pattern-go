package memento

import "testing"

func TestMemento(t *testing.T) {
	editor := NewEditor(NewDoc("《AI的觉醒》"))
	editor.Append("第一章 混沌初开")
	editor.Append("\n  正文2000字……")
	editor.Append("\n第二章 荒漠之花\n  正文3000字……")
	editor.Delete()

	//吃下后悔药，我的世界又完整了。
	editor.Undo()
}
