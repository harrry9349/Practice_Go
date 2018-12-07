package animals

import(
	"testing"
)

func TestElephantFeed(t *testing.T){
	expect := "Glass"
	actual := ElephantFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

func TestMonkeyFeed(t *testing.T){
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

func TestRabbitFeed(t *testing.T){
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

func TestLionFeed(t *testing.T){
	expect := "Meat"
	actual := LionFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}

func TestDolphinFeed(t *testing.T){
	expect := "Fish"
	actual := DolphinFeed()

	if expect != actual{
		t.Errorf("%s != %s",expect,actual)
	}
}