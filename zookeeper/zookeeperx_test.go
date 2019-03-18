package zookeeper

import "testing"

func TestDelete(t *testing.T) {
	Delete("/parent/path/to/delete")
}

func TestImportFromFile(t *testing.T) {
	ImportFromFile("/tmp/zk-test.txt", "")
	Export("127.0.0.1", "/k0", "/tmp/zk-test-export.txt")
}

func TestReImportWithParent(t *testing.T) {
	root := "/k0"
	Delete(root)
	ImportFromFile("/tmp/zk-test.txt", "")
	Export("127.0.0.1", "/k0", "/tmp/zk-test-export.txt")
}

func TestExport(t *testing.T) {
	Export("127.0.0.1", "/k0", "/tmp/zk-test-export.txt")
}
