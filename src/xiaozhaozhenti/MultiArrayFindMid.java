package xiaozhaozhenti;

/**
 * 求两个等长度有序数组的中位数
 * @author yefengzhichen
 * 2016年8月22日
 */
public class MultiArrayFindMid {
	public static int getMid(int[] arr1, int[] arr2, int n) {
		int index1 = 0, index2 = 0;
		int pos = 0;
		int resu1 = -1, resu2 = -1;
		while(pos <= n) {
			if(index1 < n && (arr1[index1] <= arr2[index2] || index2 >= n)) {
				resu1 = resu2;
				resu2 = arr1[index1++];
			} else {
				resu1 = resu2;
				resu2 = arr2[index2++];
			}
			pos++;
		}
		return (resu1 + resu2)/2;
	}
	
	public static void main(String[] args){
	    int a1[] = {1, 12, 15, 26, 38};
	    int a2[] = {2, 13, 17, 30, 45};
	    int resu = getMid(a1, a2, 5);
	    System.out.println(resu);
	}
	
}
