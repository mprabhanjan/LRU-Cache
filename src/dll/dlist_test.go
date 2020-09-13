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
func TestDlistResinsertAtHead(t *testing.T) {

    elements := make([]int, 100000)
    dlist := NewDlist()
    var link *Dlist_link
    for i := 0; i < 100000; i++ {
        elements[i] = 100+i
        if i == 50000 {
            link = dlist.InsertAtTail(elements[i])
        } else {
            dlist.InsertAtTail(elements[i])
        }
    }

    dlist.ReInsertAtHead(link)
    newElems := elements[50000:50001]
    newElems = append(newElems, elements[:50000]...)
    newElems = append(newElems, elements[50001:]...)

    iter := NewDlistIterator(dlist)
    i := 0
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != newElems[i] {
            t.Errorf("TestDlistReinsertAtHead(): Elem %d not at correct order %d!!", val, i)
        }
        i++
    }

    iter = NewDlistIterator(dlist)
    i = len(elements)-1
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != newElems[i] {
            t.Errorf("TestDlistReinsertAtHead() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i--
    }
}
func TestDlistResinsertAtTail(t *testing.T) {

    elements := make([]int, 100000000)
    dlist := NewDlist()
    var link *Dlist_link
    for i := 0; i < 100000000; i++ {
        elements[i] = 100+i
        if i == 99999 {
            link = dlist.InsertAtTail(elements[i])
        } else {
            dlist.InsertAtTail(elements[i])
        }
    }

    dlist.ReInsertAtTail(link)
    newElems := elements[:99999]
    newElems = append(newElems, elements[99999+1:]...)
    newElems = append(newElems, elements[99999])

    iter := NewDlistIterator(dlist)
    i := 0
    for el := iter.Next(); el != nil; el = iter.Next() {
        val := el.(int)
        if val != newElems[i] {
            t.Errorf("TestDlistReinsertAtHead(): Elem %d not at correct order %d!!", val, i)
        }
        i++
    }

    iter = NewDlistIterator(dlist)
    i = len(elements)-1
    for el := iter.ReverseNext(); el != nil; el = iter.ReverseNext() {
        val := el.(int)
        if val != newElems[i] {
            t.Errorf("TestDlistReinsertAtHead() - Reversewalk: Elem %d not at correct order %d!!", val, i)
        }
        i--
    }
}