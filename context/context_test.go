package context

import "testing"

// Tests if the string "test.yml" generates a correct OrbitFileMap slice.
func TestSimplePath(t *testing.T) {
	filesMap, err := getFilesMap("test.yml")
	if err != nil {
		panic(err)
	}

	if len(filesMap) != 1 {
		t.Error("Files map should have a length of one!")
	}

	if filesMap[0].Name != "default" {
		t.Error("Item of files map at index 0 should have \"default\" as name!")
	}

	if filesMap[0].Path != "test.yml" {
		t.Error("Item of files map at index 0 should have \"test.yml\" as path!")
	}
}

// Tests if the string "first,first.yml;last,last.yml" generates a correct OrbitFileMap slice.
func TestManyPaths(t *testing.T) {
	filesMap, err := getFilesMap("first,first.yml;last,last.yml")
	if err != nil {
		panic(err)
	}

	if len(filesMap) != 2 {
		t.Error("Files map should have a length of two!")
	}

	if filesMap[0].Name != "first" {
		t.Error("Item of files map at index 0 should have \"first\" as name!")
	}

	if filesMap[0].Path != "first.yml" {
		t.Error("Item of files map at index 0 should have \"first.yml\" as path!")
	}

	if filesMap[1].Name != "last" {
		t.Error("Item of files map at index 1 should have \"last\" as name!")
	}

	if filesMap[1].Path != "last.yml" {
		t.Error("Item of files map at index 0 should have \"last.yml\" as path!")
	}
}
