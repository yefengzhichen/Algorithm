package codingInterviewGuide;

import org.junit.Test;

public class insertNodeToSingleCircleLinkedList {
	
	public LinkedNode insertNode(LinkedNode head, int x) {
		if(head == null) {
			LinkedNode tmp = new LinkedNode(x);
			tmp.next = tmp;
			return tmp;
		}
		LinkedNode cur = head;
		LinkedNode pre = null;
		while(cur.next != head) {
			cur = cur.next;
		}
		pre = cur;  //环形链表最后一个点
		cur = head;
		while(cur.next != head) {
			if(cur.value >= x) {
				pre.next = new LinkedNode(x);
				pre.next.next = cur;
				head = (cur == head ? pre.next : head);
				return head;
			}
			pre = cur;
			cur = cur.next;
		}
		if(cur.value >= x) {
			pre.next = new LinkedNode(x);
			pre.next.next = cur;
			head = (cur == head ? pre.next : head);			
		} else {
			cur.next = new LinkedNode(x);
			cur.next.next = head;
		}
		return head;
	}
	
	@Test
	public void test() {
		LinkedNode head = new LinkedNode(1); head.next = new LinkedNode(2); head.next.next = new LinkedNode(3);
		head.next.next.next = new LinkedNode(5); head.next.next.next.next = head; 
		LinkedNode res = insertNode(head, 7);
		LinkedNode cur = res;
		while(cur.next != res) {
			System.out.print(cur.value + " ");
			cur = cur.next;
		} 
		System.out.println(cur.value);
	}
}
