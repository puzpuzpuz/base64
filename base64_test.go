package base64_test

import (
	crand "crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"testing"

	cristalhq "github.com/cristalhq/base64"

	. "github.com/puzpuzpuz/base64"
)

const (
	BENCH_ARRAY_LEN = 2048
	letterBytes     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func TestDecode_RandomBytes(t *testing.T) {
	for i := 1; i < 1000; i++ {
		randomBytes := make([]byte, i)
		_, err := crand.Read(randomBytes)
		if err != nil {
			t.Fatal(err)
		}
		src := make([]byte, base64.StdEncoding.EncodedLen(len(randomBytes)))
		base64.StdEncoding.Encode(src, randomBytes)
		n := base64.StdEncoding.DecodedLen(len(src))
		dst := make([]byte, n)
		dstStd := make([]byte, n)

		_, err = base64.StdEncoding.Decode(dstStd, src)
		if err != nil {
			t.Fatal()
		}
		Decode(dst, src)

		for j := 0; j < n; j++ {
			if dst[j] != dstStd[j] {
				t.Fail()
			}
		}
	}
}

func BenchmarkDecode_RandomBytes(b *testing.B) {
	randomBytes := make([]byte, BENCH_ARRAY_LEN)
	_, err := crand.Read(randomBytes)
	if err != nil {
		b.Fatal(err)
	}
	src := make([]byte, base64.StdEncoding.EncodedLen(len(randomBytes)))
	base64.StdEncoding.Encode(src, randomBytes)
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))

	b.Run("std/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = base64.StdEncoding.Decode(dst, src)
			if err != nil {
				b.Fatal()
			}
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})

	b.Run("cristalhq/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = cristalhq.StdEncoding.Decode(dst, src)
			if err != nil {
				b.Fatal()
			}
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})

	b.Run("puzpuzpuz/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Decode(dst, src)
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})
}

func BenchmarkDecode_RandomAscii(b *testing.B) {
	var err error

	src := []byte(randStringBytes(BENCH_ARRAY_LEN))
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))

	b.Run("std/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = base64.StdEncoding.Decode(dst, src)
			if err != nil {
				b.Fatal()
			}
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})

	b.Run("cristalhq/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_, err = cristalhq.StdEncoding.Decode(dst, src)
			if err != nil {
				b.Fatal()
			}
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})

	b.Run("puzpuzpuz/Decode", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Decode(dst, src)
			b.SetBytes(BENCH_ARRAY_LEN)
		}
	})
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[mrand.Intn(len(letterBytes))]
	}
	return string(b)
}
