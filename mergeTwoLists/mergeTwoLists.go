package mergetwolists

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = new(ListNode) // 前置虚拟节点法
	cur := ret              // 定义一个保存当前节点的变量
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next // 当前节点向后移
	}
	// 追加没遍历到的链表
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return ret.Next
}
