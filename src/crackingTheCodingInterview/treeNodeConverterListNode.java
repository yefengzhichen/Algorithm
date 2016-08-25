package crackingTheCodingInterview;

import org.junit.Test;

class ListNode {
	int val;
	ListNode next = null;

	ListNode(int val) {
		this.val = val;
	}
}

class TreeNode {
	int val = 0;
	TreeNode left = null;
	TreeNode right = null;

	public TreeNode(int val) {
		this.val = val;
	}
}


/**
 * 二叉树转换为链表
 * @author yefengzhichen
 * 2016年8月25日
 */
public class treeNodeConverterListNode {
	
	ListNode resu = new ListNode(0);
	ListNode last = resu;
	public ListNode treeToList(TreeNode root) {
        // write code here
        if (root == null) {
            return null;
        }    
        transfer(root);
        return resu.next;
    }
    
    public void transfer(TreeNode root) {
        if(root != null) {
            transfer(root.left);
            last.next = new ListNode(root.val);
            last = last.next;
            transfer(root.right);
        }
    }

	@Test
	public void test1() {
		TreeNode root = new TreeNode(4);
		root.left = new TreeNode(2);
		root.right = new TreeNode(6);
		root.left.left = new TreeNode(1);
		root.left.right = new TreeNode(3);
		root.right.left = new TreeNode(5);
		root.right.right = new TreeNode(7);
		ListNode resu = treeToList(root);
		System.out.println(resu.val);
	}
}
