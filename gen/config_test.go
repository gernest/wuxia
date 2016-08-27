package gen

import "testing"

func TestPlan(t *testing.T) {

	strategy1 := &Strategy{
		Title:     "full_path",
		FullMatch: true,
		Pattern:   "/root/*.md",
	}

	strategy2 := &Strategy{
		Title:   "base",
		Pattern: "*.css",
	}
	p := &Plan{
		Strategies: []*Strategy{strategy1, strategy2},
	}
	txt := "/root/README.md"
	s, err := p.FindStrategy(txt)
	if err != nil {
		t.Fatal(err)
	}
	if s.Title != strategy1.Title {
		t.Errorf("expected %s got %s", strategy1.Title, s.Title)
	}

	s, err = p.FindStrategy("/root/css/style.css")
	if err != nil {
		t.Fatal(err)
	}
	if s.Title != strategy2.Title {
		t.Errorf("expected %s got %s", strategy2.Title, s.Title)
	}
}
