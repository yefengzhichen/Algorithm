package codingInterviewGuide;

class TreeNode {
	int value;
	TreeNode left;
	TreeNode right;
	
	TreeNode(int value) {
		this.value = value;
	}	
}


public class sortedBinaryTreeToDoubleLinkedList {

	public static void main(String[] args) {
		TreeNode root = new TreeNode(4); root.left = new TreeNode(2); root.right = new TreeNode(6);
		root.left.left = new TreeNode(1); root.left.right = new TreeNode(3);
		root.right.left = new TreeNode(5); root.right.right = new TreeNode(7);
		TreeNode cur = root;
		TreeNode res = null;
		while(cur.left != null) {
			cur = cur.left;
		}
		res = cur;
		TreeNode pre = null;
		transfer(root, pre);
		//打印转换后的数据
		cur = res;
		while(cur.right != null) {
			System.out.print(cur.value + " ");
			cur = cur.right;
		}
		System.out.println(cur.value);
		while(cur.left != null) {
			System.out.print(cur.value + " ");
			cur = cur.left;
		}
		System.out.println(cur.value);
	}
	
	public static TreeNode transfer(TreeNode root, TreeNode pre) {
		if(root == null) {
			return pre;
		}
		pre = transfer(root.left, pre);
		if(pre == null) {
			pre = root;
			root.left = null;
		} else {
			pre.right = root;
			root.left = pre;
			pre = root;
		}				
		pre = transfer(root.right, pre);
		return pre;
	}
	
}
