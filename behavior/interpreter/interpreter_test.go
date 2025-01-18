package interpreter

import "testing"

func TestInterpreter(t *testing.T) {
	/*
	    * BEGIN             // 脚本开始
	    * MOVE 500,600;     // 鼠标移动到坐标(500, 600)
	    *  BEGIN LOOP 5     // 开始循环5次
	    *      LEFT_CLICK;  // 循环体内单击左键
	    *      DELAY 1;     // 每次延时1秒
	    *  END;             // 循环体结束
	   * RIGHT_DOWN;       // 按下右键
	   * DELAY 7200;       // 延时2小时
	   * END;              // 脚本结束
	*/

	// 构造指令集语义树，实际情况会交给语法解析器（Evaluator or Parser）。
	sequence := NewSequence([]Expression{
		NewMove(500, 600),
		NewRepetition(
			NewSequence(
				[]Expression{NewLeftClick(), NewDelay(1)},
			),
			5,
		),
		new(RightDown),
		NewDelay(7200),
	})
	sequence.Interpret()
}
