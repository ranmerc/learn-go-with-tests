package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "wrong://website"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://www.example.com",
		"wrong://website",
	}

	want := map[string]bool{
		"https://www.google.com":  true,
		"https://www.example.com": true,
		"wrong://website":         false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)

	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
