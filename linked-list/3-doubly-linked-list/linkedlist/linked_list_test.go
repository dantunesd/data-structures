package linkedlist

import (
	"data-structures-algorithms/linked-list/3-doubly-linked-list/node"
	"reflect"
	"testing"
)

func TestLinkedList_GetNodesData(t *testing.T) {
	var nodesData []int

	nodeTwo := node.NewNode(2)

	nodeOne := node.NewNode(1)
	nodeOne.SetNext(nodeTwo)

	nodeZero := node.NewNode(0)
	nodeZero.SetNext(nodeOne)

	nodeZero.SetPrevious(nil)
	nodeOne.SetPrevious(nodeZero)
	nodeTwo.SetPrevious(nodeOne)

	type fields struct {
		head *node.Node
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "should return an empty list when head is nil",
			fields: fields{
				head: nil,
			},
			want: nodesData,
		},
		{
			name: "should return a list with only the head data",
			fields: fields{
				head: node.NewNode(0),
			},
			want: []int{0},
		},
		{
			name: "should return a list with 3 datas after pushed",
			fields: fields{
				head: nodeZero,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLinkedList(tt.fields.head)
			if got := l.GetNodesData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.GetNodesData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Push_OnAnEmptyList(t *testing.T) {
	linkedList := NewLinkedList(nil)
	linkedList.Push(0)

	want := []int{0}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Push() = %v, want %v", got, want)
	}
}

func TestLinkedList_Push_OnAListWithHeadOnly(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(1))
	linkedList.Push(0)

	want := []int{0, 1}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Push() = %v, want %v", got, want)
	}
}

func TestLinkedList_Push_OnAListWithManyItems(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(1))
	linkedList.Push(0)
	linkedList.Push(-1)

	want := []int{-1, 0, 1}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Push() = %v, want %v", got, want)
	}
}

func TestLinkedList_Append_OnAnEmptyList(t *testing.T) {
	linkedList := NewLinkedList(nil)
	linkedList.Append(0)

	want := []int{0}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Append() = %v, want %v", got, want)
	}
}

func TestLinkedList_Append_OnAListWithHeadOnly(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(1))
	linkedList.Append(2)

	want := []int{1, 2}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Append() = %v, want %v", got, want)
	}
}

func TestLinkedList_Append_OnAListWithManyItems(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(1))
	linkedList.Append(2)
	linkedList.Append(3)

	want := []int{1, 2, 3}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Append() = %v, want %v", got, want)
	}
}

func TestLinkedList_DeleteFirstOcurrency_OnAnEmptyList(t *testing.T) {
	linkedList := NewLinkedList(nil)
	linkedList.DeleteFirstOcurrency(0)

	var want []int
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_DeleteFirstOcurrency_DeleteHeadOnAListWithHeadOnly(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(0))
	linkedList.DeleteFirstOcurrency(0)

	var want []int
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_DeleteFirstOcurrency_DeleteHeadOnAListWithManyItems(t *testing.T) {
	tail := node.NewNode(1)
	head := node.NewNode(0)
	head.SetNext(tail)

	head.SetPrevious(nil)
	tail.SetPrevious(head)

	linkedList := NewLinkedList(head)
	linkedList.DeleteFirstOcurrency(0)

	want := []int{1}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_DeleteFirstOcurrency_DeleteAMidNodeOnAListWithManyItems(t *testing.T) {
	tail := node.NewNode(2)

	mid := node.NewNode(1)
	mid.SetNext(tail)

	head := node.NewNode(0)
	head.SetNext(mid)

	head.SetPrevious(nil)
	mid.SetPrevious(head)
	tail.SetPrevious(mid)

	linkedList := NewLinkedList(head)
	linkedList.DeleteFirstOcurrency(1)

	want := []int{0, 2}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_DeleteFirstOcurrency_DeleteTheLastNodeOnAListWithManyItems(t *testing.T) {
	tail := node.NewNode(3)

	midTwo := node.NewNode(2)
	midTwo.SetNext(tail)

	mid := node.NewNode(1)
	mid.SetNext(midTwo)

	head := node.NewNode(0)
	head.SetNext(mid)

	head.SetPrevious(nil)
	mid.SetPrevious(head)
	midTwo.SetPrevious(mid)
	tail.SetPrevious(tail)

	linkedList := NewLinkedList(head)
	linkedList.DeleteFirstOcurrency(1)

	want := []int{0, 2, 3}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_Reverse_OnAnEmptyList(t *testing.T) {
	linkedList := NewLinkedList(nil)
	linkedList.Reverse()

	var want []int
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_Reverse_OnAnListWithHeadOnly(t *testing.T) {
	linkedList := NewLinkedList(node.NewNode(1))
	linkedList.Reverse()

	want := []int{1}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.DeleteFirstOcurrency() = %v, want %v", got, want)
	}
}

func TestLinkedList_Reverse_OnAListWithManyItems(t *testing.T) {
	tail := node.NewNode(2)

	mid := node.NewNode(1)
	mid.SetNext(tail)

	head := node.NewNode(0)
	head.SetNext(mid)

	head.SetPrevious(nil)
	mid.SetPrevious(head)
	tail.SetPrevious(mid)

	linkedList := NewLinkedList(head)

	linkedList.Reverse()

	want := []int{2, 1, 0}
	got := linkedList.GetNodesData()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LinkedList.Reverse() = %v, want %v", got, want)
	}
}
