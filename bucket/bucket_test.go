package bucket

import "testing"

func TestBucket(t *testing.T) {
	t.Run("Test Bucket func", func(t *testing.T) {

		want := "1"
		get := Bucket(1)

		if want != get {
			t.Errorf("Not correct %s, %s", want, get)
		}
	})
	t.Run("Test Bucket func", func(t *testing.T) {
		want := "2"
		get := Bucket(2)
		if want != get {
			t.Errorf("Not correct %s, %s", want, get)
		}
	})

}

func TestBucketThree(t *testing.T) {

	want := "Fizz"
	get := Bucket(3)

	if want != get {
		t.Errorf("Not correct %s, %s", want, get)
	}
}
