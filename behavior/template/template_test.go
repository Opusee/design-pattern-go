package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	human := NewHuman()
	human.FeedMilk()
	human.Move()

	whale := NewWhale()
	whale.FeedMilk()
	whale.Move()

	bat := NewBat()
	bat.FeedMilk()
	bat.Move()

	println("-------------")

	NewAutoTestPM().Kickoff()
}
