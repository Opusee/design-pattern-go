package chain

/**
责任链模式
*/

type Approver interface {
	Approve(amount int) // 抽象审批方法由具体审批人子类实现
	SetNextApprover(nextApprover Approver) Approver
}

// Approver 审批人抽象
type approver struct {
	name         string   // 抽象出审批人的姓名。
	nextApprover Approver // 下一个审批人，更高级别领导。
}

func NewApprover(name string) *approver {
	return &approver{name: name}
}

func (a *approver) SetNextApprover(nextApprover Approver) Approver {
	a.nextApprover = nextApprover
	return a.nextApprover // 返回下个审批人，链式编程。
}

type Staff struct {
	*approver
}

func NewStaff(name string) *Staff {
	return &Staff{NewApprover(name)}
}

func (s *Staff) Approve(amount int) {
	if amount <= 1000 {
		println("审批通过。【员工：", s.name, "】")
	} else {
		println("无权审批，升级处理。【员工：", s.name, "】")
		s.nextApprover.Approve(amount)
	}
}

type Manager struct {
	*approver
}

func NewManager(name string) *Manager {
	return &Manager{NewApprover(name)}
}

func (m *Manager) Approve(amount int) {
	if amount <= 5000 {
		println("审批通过。【经理：", m.name, "】")
	} else {
		println("无权审批，升级处理。【员工：", m.name, "】")
		m.nextApprover.Approve(amount)
	}
}

type CEO struct {
	*approver
}

func NewCEO(name string) *CEO {
	return &CEO{NewApprover(name)}
}

func (c *CEO) Approve(amount int) {
	if amount <= 10000 {
		println("审批通过。【CEO：", c.name, "】")
	} else {
		println("驳回申请。【CEO：", c.name, "】")
	}
}
