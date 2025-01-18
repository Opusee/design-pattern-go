package bridge

import "testing"

func TestBridge(t *testing.T) {
	//白色画笔对应的所有形状
	NewWhitePen(new(CircleRuler)).Draw()
	NewWhitePen(new(SquareRuler)).Draw()
	NewWhitePen(new(TriangleRuler)).Draw()

	//黑色画笔对应的所有形状
	NewBlackPen(new(CircleRuler)).Draw()
	NewBlackPen(new(SquareRuler)).Draw()
	NewBlackPen(new(TriangleRuler)).Draw()
}
