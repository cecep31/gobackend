package posts

import "testing"

func TestPost(t *testing.T) {

	// TODO
	got := 8 + 2
	want := 10

	if condition := got == want; !condition {
		t.Error("got", got, "want", want)
	}
}
