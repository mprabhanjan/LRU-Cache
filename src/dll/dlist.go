package dll

type Dlist_link struct {
    l_next *Dlist_link
    l_prev *Dlist_link
    user_data interface{}
}

type Dlist struct {
    head *Dlist_link
    tail *Dlist_link
    num_elements int
    num_inserts int
    num_removes int
}

func NewDlist() *Dlist {
    new_list := &Dlist {
    }
    return new_list
}

func (list *Dlist) insertAtHead(link *Dlist_link) {
    if list.head == nil {
        list.head = link
        list.tail = link
    } else {
        link.l_next = list.head
        list.head.l_prev = link
        list.head = link
    }
}

func (list *Dlist) insertAtTail(link *Dlist_link) {
    if list.tail == nil {
        list.tail = link
        list.head = link
    } else {
        link.l_prev = list.tail
        list.tail.l_next = link
        list.tail = link
    }
}

func (list *Dlist) InsertAtHead(user_data interface{}) *Dlist_link {
    new_link := &Dlist_link{user_data: user_data}
    list.insertAtHead(new_link)
    return new_link
}

func (list *Dlist) InsertAtTail(user_data interface{}) *Dlist_link {
    new_link := &Dlist_link{user_data: user_data}
    list.insertAtTail(new_link)
    return new_link
}

func (list *Dlist) RemoveLink(link *Dlist_link) {
    if list.head == link {
        list.head = link.l_next
    } else {
        link.l_prev.l_next = link.l_next
    }
    if list.tail == link {
        list.tail = link.l_prev
    } else {
        link.l_next.l_prev = link.l_prev
    }
    link.l_prev = nil
    link.l_next = nil
    link.user_data = nil
    return
}

func (list *Dlist) ReInsertAtHead(link *Dlist_link) {
    if list.head == link {
        return
    }

    list.RemoveLink(link)
    list.insertAtHead(link)
}

func (list *Dlist) ReInsertAtTail(link *Dlist_link) {
    if list.tail == link {
        return
    }

    list.RemoveLink(link)
    list.insertAtTail(link)
}

func (list *Dlist) PopHead() interface{} {
    if list.head == nil {
        return nil
    }
    user_data := list.head.user_data
    list.RemoveLink(list.head)
    return user_data
}

func (list *Dlist) PopTail() interface{} {
    if list.tail == nil {
        return nil
    }
    user_data := list.tail.user_data
    list.RemoveLink(list.tail)
    return user_data
}

type DlistIterator struct {
    dlist *Dlist
    next  *Dlist_link
    started bool
}

func NewDlistIterator(list *Dlist) *DlistIterator {
    iter := &DlistIterator{
        dlist: list,
        started: false,
        next: nil,
    }
    return iter
}

func (iter *DlistIterator) Next() interface {} {

    if iter.dlist == nil {
        return nil
    }
    if !iter.started && iter.next == nil {
        iter.started = true
        iter.next = iter.dlist.head
    } else if iter.next != nil {
        iter.next = iter.next.l_next
    } else {
        return nil
    }
    if iter.next != nil {
        return iter.next.user_data
    } else {
        return nil
    }
}

func (iter *DlistIterator) ReverseNext() interface {} {

    if iter.dlist == nil {
        return nil
    }
    if !iter.started && iter.next == nil {
        iter.started = true
        iter.next = iter.dlist.tail
    } else if iter.next != nil {
        iter.next = iter.next.l_prev
    } else {
        return nil
    }
    if iter.next != nil {
        return iter.next.user_data
    } else {
        return nil
    }
}