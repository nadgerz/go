package main

func TestRacer(t *testingT) {
	slowURL = "http://www.facebook.com"
	fastURL = "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
