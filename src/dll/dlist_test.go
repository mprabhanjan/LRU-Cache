package dll

import "testing"

func TestDlistHeadInsertion(t *testing.T) {

    elements := make([]int, 10)
    dlist := NewDlist()
    for i := 0; i < 10; i++ {
        elements[i] = 100+i
        dlist.InsertAtHead(elements[i])
    }

    iter := NewDlistIterator(dlist)
    i := len(elements)-1
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistInsertion(): Elem %d not at correct order %d!!", val, i)
        }
        i--
    }

    iter = NewDlistIterator(dlist)
    i = 0
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistInsertion() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i++
    }
}

func TestDlistTailInsertion(t *testing.T) {

    elements := make([]int, 10)
    dlist := NewDlist()
    for i := 0; i < 10; i++ {
        elements[i] = 100+i
        dlist.InsertAtTail(elements[i])
    }

    iter := NewDlistIterator(dlist)
    i := 0
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistTailInsertion(): Elem %d not at correct order %d!!", val, i)
        }
        i++
    }

    iter = NewDlistIterator(dlist)
    i = len(elements)-1
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistTailInsertion() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i--
    }
}

func TestDlistPopHead(t *testing.T) {

    elements := make([]int, 10)
    dlist := NewDlist()
    for i := 0; i < 10; i++ {
        elements[i] = 100+i
        dlist.InsertAtHead(elements[i])
    }

    val := dlist.PopHead().(int)
    if val != elements[len(elements)-1] {
        t.Errorf("TestDlistPopHead(): PopHead() expected elem %d, but got %d!!", elements[len(elements)-1], val)
    }

    iter := NewDlistIterator(dlist)
    i := len(elements)-1-1
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistPopHead(): Elem %d not at correct order %d!!", val, i)
        }
        i--
    }

    iter = NewDlistIterator(dlist)
    i = 0
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistPopHead() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i++
    }
}

func TestDlistPopTail(t *testing.T) {

    elements := make([]int, 10)
    dlist := NewDlist()
    for i := 0; i < 10; i++ {
        elements[i] = 100+i
        dlist.InsertAtTail(elements[i])
    }

    val := dlist.PopTail().(int)
    if val != elements[len(elements)-1] {
        t.Errorf("TestDlistPopTail(): PopTail() expected elem %d, but got %d!!", elements[len(elements)-1], val)
    }

    iter := NewDlistIterator(dlist)
    i := 0
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistPopTail(): Elem %d not at correct order %d!!", val, i)
        }
        i++
    }

    iter = NewDlistIterator(dlist)
    i = len(elements)-1-1
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != elements[i] {
            t.Errorf("TestDlistPopTail() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i--
    }
}
