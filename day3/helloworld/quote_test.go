package main
import "testing"
func TestHello(t *testing.T) {
	hello := "Hello, world."
	if out := helloWorld(); out != hello {
		t.Errorf("Hello() = %q, want %q", out, hello)
	}
}
func TestGlassQuote(t *testing.T) {
	glass := "I can eat glass and it doesn't hurt me."
	if out := Glass(); out != glass {
		t.Errorf("Glass() = %q, want %q", out, glass)
	}
}