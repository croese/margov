package margov

import "testing"

func TestNewGeneratorEmptyString(t *testing.T) {
	_, err := NewGenerator("")
	if err == nil {
		t.Error(`NewGenerator("") should be an error`)
	}

	if err.Error() != "input string can't be empty" {
		t.Error("error message should be 'input string can't be empty'")
	}
}

func TestNewGenerator(t *testing.T) {
	gen, err := NewGenerator("the cat and the dog")
	if err != nil {
		t.Error(`NewGenerator("the cat and the dog") should succeed`)
	}

	if gen == nil {
		t.Error(`NewGenerator("the cat and the dog") should return a non-nil Generator`)
	}
}

func TestLen(t *testing.T) {
	gen, _ := NewGenerator("the cat and the dog")
	if gen.Len() != 3 {
		t.Errorf(`%d != 4`, gen.Len())
	}
}

func TestFollowerEntries(t *testing.T) {
	gen, _ := NewGenerator("the cat and the dog")
	if len(gen.followers["the"]) != 2 {
		t.Errorf("the should have 2 followers")
	}

	if len(gen.followers["cat"]) != 1 {
		t.Errorf("cat should have 1 follower")
	}

	if len(gen.followers["and"]) != 1 {
		t.Errorf("and should have 1 follower")
	}

	if gen.followers["dog"] != nil {
		t.Errorf("dog should not be in the followers collection")
	}

	if gen.followers["cat"]["and"] != 1 {
		t.Errorf("and should be one of cat's followers")
	}
}
