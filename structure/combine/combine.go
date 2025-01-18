package combine

/**
组合模式
*/

type node struct {
	// 节点名
	Name string
}

func (n *node) Add(Node) {
	panic(n.Name + "不支持添加子节点")
}

func newNode(name string) *node {
	return &node{Name: name}
}

type Node interface {
	Ls(space ...int) // 设计为可变参是为了支持重载，参数可传可不传
	Add(child Node)
}

type Folder struct {
	*node
	// 文件夹可以包含子节点(文件或者文件夹)
	childNodes []Node
}

func NewFolder(name string) *Folder {
	// 调用父类的构造方法
	return &Folder{
		node:       newNode(name),
		childNodes: make([]Node, 0),
	}
}

func (f *Folder) Add(child Node) {
	f.childNodes = append(f.childNodes, child)
}

type File struct {
	*node
}

func NewFile(name string) *File {
	return &File{newNode(name)}
}

func (n *node) Ls(space ...int) {
	level := 0
	if space != nil {
		level = space[0]
	}
	for i := 0; i < level; i++ {
		print(" ") // 先循环输出 n 个空格
	}
	println(n.Name) // 然后再打印自己的名字
}

func (f *Folder) Ls(space ...int) {
	level := 0
	if space != nil {
		level = space[0]
	}
	// 调用父类共通的ls方法列出自己的名字。
	f.node.Ls(level)
	level++
	//之后列出的子节点前，空格数要增加一个了。
	for _, n := range f.childNodes {
		n.Ls(level) //调用子节点的ls方法。
	}
}
