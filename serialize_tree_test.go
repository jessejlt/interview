package interview_test

import (
	"testing"

	"github.com/jessejlt/interview"
)

func TestSerializeBinaryTree(t *testing.T) {

	tree := interview.NewBinaryTree(10)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(1)

	have := interview.SerializeBinaryTree(tree)
	expect := "10,3,5,1"

	if have != expect {
		t.Fatalf("Have=%s, Expect=%s", have, expect)
	}
}

func TestDeserializeBinaryTree(t *testing.T) {

	expect := "10,3,5,1"
	tree, err := interview.DeserializeBinaryTree(expect)
	if err != nil {
		t.Fatal(err)
	}

	have := interview.SerializeBinaryTree(tree)
	if have != expect {
		t.Fatalf("Have=%s, Expect=%s", have, expect)
	}
}
