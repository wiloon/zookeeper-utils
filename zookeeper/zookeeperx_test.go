package zookeeper

import "testing"

var root = "/k0"

// import
func TestImportFromFile(t *testing.T) {
	ImportFromFile("127.0.0.1", "/tmp/foo.txt", "")
}

// export
func TestExport(t *testing.T) {
	Export("127.0.0.1", "", "/tmp/foo.txt")
}

func TestDelete(t *testing.T) {
	Delete("127.0.0.1", "")
}

func TestImportExport(t *testing.T) {
	ImportFromFile("127.0.0.1", "/tmp/foo.txt", "")
	Export("127.0.0.1", "/k0", "/tmp/foo.txt")
}

func TestReImportWithParent(t *testing.T) {
	root := "/k0"
	Delete("127.0.0.1", root)
	ImportFromFile("127.0.0.1", "/tmp/foo.txt", "")
	Export("127.0.0.1", "/k0", "/tmp/foo.txt")
}
