package codingInterviewGuide;

import org.junit.Test;

class LinkedNode {
	int value;
	LinkedNode next;
	public LinkedNode(int value) {
		this.value = value;
	}	
	
	@Override
	public String toString(){
		String string = "" + value;
		LinkedNode node = next;
		while(node != null) {
			string += (" " + node.value);
			node = node.next;
		}
		return string;
	}	
}

public class selectSortOfSingleLinkedList {
	
	public LinkedNode selectSort(LinkedNode head) {
		if (head == null || head.next == null) {
			return head;
		}
		LinkedNode res = null;
		LinkedNode pre = null;
		LinkedNode cur = null;
		LinkedNode root = head;
		while (root != null) {
			cur = root;
			LinkedNode min = root;
			LinkedNode minPre = null;
			while (cur.next != null) {
				if (cur.next.value < min.value) {
					min = cur.next;
					minPre = cur;					
				}
				cur = cur.next;
			}
			if (pre == null) {
				res = pre = min;
			} else {
				pre.next = min;
				pre = min;
			}
			if (minPre == null) {
				root = root.next;
			} else {
				minPre.next = min.next;
			}
			min.next = null;
		}
		return res;	
	}
	
	@Test
	public void test() {
		LinkedNode head = new LinkedNode(1); head.next = new LinkedNode(2); head.next.next = new LinkedNode(0);
		head.next.next.next = new LinkedNode(3); head.next.next.next.next = new LinkedNode(5); 
		LinkedNode res = selectSort(head);
		System.out.println(res);
	}
	
}
