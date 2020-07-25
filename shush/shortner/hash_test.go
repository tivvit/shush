package shortner

import (
	"testing"
)

func TestHashShortner(t *testing.T) {
	testHash(t, "http://example.com", "md5", 5, "a9b9f")
	testHash(t, "http://example.com", "md5", 30, "a9b9f04336ce0181a08e774e01113b")
	testHash(t, "http://example.com", "md5", 50, "a9b9f04336ce0181a08e774e01113b31")
	testHash(t, "http://example.com", "sha1", 50, "89dce6a446a69d6b9bdc01ac75251e4c322bcdff")
	testHash(t, "http://example.com", "sha256", 50, "f0e6a6a97042a4f1f1c87f5f7d44315b2d852c2df5c7991cc6")
	testHash(t, "http://example.com", "sha512", 50, "75584225a2f4e84caa1d830ff6195cdaf0f667d6b0bf92a7fc")
	testHash(t, "http://example.com", "fnv32", 50, "687474703a2f2f6578616d706c652e636f6d811c9dc5")
	testHash(t, "http://example.com", "fnv32a", 50, "687474703a2f2f6578616d706c652e636f6d811c9dc5")
	testHash(t, "http://example.com", "fnv64", 50, "687474703a2f2f6578616d706c652e636f6dcbf29ce4842223")
	testHash(t, "http://example.com", "fnv64a", 50, "687474703a2f2f6578616d706c652e636f6dcbf29ce4842223")
	testHash(t, "http://example.com", "fnv128", 50, "687474703a2f2f6578616d706c652e636f6d6c62272e07bb01")
	testHash(t, "http://example.com", "fnv128a", 50, "687474703a2f2f6578616d706c652e636f6d6c62272e07bb01")
	testHash(t, "http://example.com", "adler32", 50, "687474703a2f2f6578616d706c652e636f6d00000001")
	testHash(t, "http://example.com", "crc32ieee", 50, "687474703a2f2f6578616d706c652e636f6d00000000")
	testHash(t, "http://example.com", "crc64iso", 50, "687474703a2f2f6578616d706c652e636f6d00000000000000")
	testHash(t, "http://example.com", "crc64ecma", 50, "687474703a2f2f6578616d706c652e636f6d00000000000000")
	_, err := Hash([]byte("http://example.com"), "abc", 50)
	if err == nil {
		t.Error("Accepted unknown hash func")
	}
}

func testHash(t *testing.T, url string, hashFunc string, ln int, result string) {
	h, err := Hash([]byte(url), hashFunc, ln)
	if err != nil {
		t.Error(err)
	}
	//t.Logf("%s hash of %s is %s", hashFunc, url, h)
	if h != result {
		t.Errorf("%s is incorrect", hashFunc)
	}
}
