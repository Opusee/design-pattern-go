package memento

/**
备忘录模式
*/

type Doc struct {
	title string // 文章标题
	body  string //文章内容
}

// NewDoc 新建文档先命名
func NewDoc(title string) *Doc {
	return &Doc{
		title: title,
		body:  "",
	}
}

func (d *Doc) SetTitle(title string) {
	d.title = title
}

func (d *Doc) GetTitle() string {
	return d.title
}

func (d *Doc) SetBody(body string) {
	d.body = body
}

func (d *Doc) GetBody() string {
	return d.body
}

// CreateHistory 创建历史记录
func (d *Doc) CreateHistory() *History {
	return NewHistory(d.body)
}

// RestoreHistory 恢复历史记录
func (d *Doc) RestoreHistory(history *History) {
	d.body = history.GetBody()
}

// Editor 编辑器
type Editor struct {
	Doc             *Doc       //文档引用
	HistoryRecords  []*History // 历史记录列表
	HistoryPosition int        // 历史记录当前位置
}

func NewEditor(doc *Doc) *Editor {
	println("<<<打开文档", doc.GetTitle())
	e := &Editor{
		Doc:             doc,                 // 注入文档
		HistoryRecords:  make([]*History, 0), // 初始化历史记录
		HistoryPosition: -1,
	}
	e.Backup() // 保存一份历史记录
	e.Show()
	return e
}

func (e *Editor) Append(txt string) {
	doc := e.Doc
	e.Doc.SetBody(doc.GetBody() + txt)
	e.Backup()
	e.Show()
}

func (e *Editor) Save() {
	println("<<<存盘操作")
}

func (e *Editor) Delete() {
	println("<<<删除操作")
	e.Doc.SetBody("")
	e.Backup()
	e.Show()
}

func (e *Editor) Backup() {
	e.HistoryRecords = append(e.HistoryRecords, e.Doc.CreateHistory())
	e.HistoryPosition++
}

// Show 显示当前文本内容
func (e *Editor) Show() {
	println(e.Doc.GetBody())
	println("文章结束>>>\n")
}

// Undo 撤销操作：如按下Ctr+Z，回到过去。
func (e *Editor) Undo() {
	println(">>>撤销操作")
	if e.HistoryPosition == 0 {
		return //到头了，不能再撤销了。
	}
	e.HistoryPosition-- //历史记录位置回滚一笔
	history := e.HistoryRecords[e.HistoryPosition]
	e.Doc.RestoreHistory(history) //取出历史记录并恢复至文档
	e.Show()
}

// History 用于备忘文章内容
type History struct {
	body string
}

func NewHistory(body string) *History {
	return &History{body: body}
}

func (h *History) GetBody() string {
	return h.body
}
