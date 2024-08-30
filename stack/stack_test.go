package stack

import "testing"

func TestNewWithoutSize(t *testing.T) {
	stack := New[int]()
	if len(stack.items) != 0 {
		t.Errorf("Expected stack size 0, got %d", len(stack.items))
	}
}

func TestNewWithSize(t *testing.T) {
	stack := New[int](10)
	if len(stack.items) != 10 {
		t.Errorf("Expected stack size 10, got %d", len(stack.items))
	}
}

func TestMultipleSizes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for multiple size arguments, but did not panic")
		}
	}()
	New[int](10, 20)
}

func TestPush(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Expected stack size 2, got %d", stack.Size())
	}

	if stack.Top() != 2 {
		t.Errorf("Expected top element 2, got %d", stack.Top())
	}
}

func TestPop(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Pop()

	if stack.Size() != 1 {
		t.Errorf("Expected stack size 1 after pop, got %d", stack.Size())
	}

	if stack.Top() != 1 {
		t.Errorf("Expected top element 1 after pop, got %d", stack.Top())
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping from an empty stack, but did not panic")
		}
	}()
	stack.Pop()
}

func TestTopEmptyStack(t *testing.T) {
	stack := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when accessing top element of an empty stack, but did not panic")
		}
	}()
	stack.Top()
}

func TestIsEmpty(t *testing.T) {
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Expected stack to be not empty")
	}
}

func TestClear(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Clear()

	if stack.Size() != 0 {
		t.Errorf("Expected stack size 0 after clear, got %d", stack.Size())
	}
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after clear")
	}
}
