package crackingTheCodingInterview;

import org.junit.Assert;
import org.junit.Test;

/**最小调整有序
 * @author yefengzhichen
 * 2016年8月23日
 * 有一个整数数组，请编写一个函数，找出索引m和n，只要将m和n之间的元素排好序，整个数组就是有序的。
 * 注意：n-m应该越小越好，也就是说，找出符合条件的最短序列.
 * 给定一个int数组A和数组的大小n，请返回一个二元组，代表所求序列的起点和终点。
 * (原序列位置从0开始标号,若原序列有序，返回[0,0])。保证A中元素均为正整数。
 * 测试样例：[1,4,6,5,9,10],6返回：[2,3]
 */
public class minAdjustRangeToSorted {
	public int[] findSegment(int[] A, int n) {
        // write code here
        int[] resu=new int[2];
        if(A.length<2||n<2){
            return resu;
        }
        int max=A[0];
        int high = 0;
        for(int i = 1; i < n; i++) {
            if(A[i] < max) {
                high = i;
            } else {
                max = A[i];
            }
        }
        int min = A[n - 1];
        int low = n - 1;
        for (int i = n - 2; i >= 0; i--) {
            if(A[i] > min) {
                low = i;
            } else {
                min = A[i];
            }
        }
        if(low < high) {
            resu[0] = low;
            resu[1] = high;
        } 
        return resu;       
    }
	
	@Test
	public void test1(){
		int a[] = {1, 12, 15, 26, 38};
	    int[] resu = findSegment(a, 5);
	    Assert.assertTrue(resu[0] == 0 || resu[1] == 0);
	}
	
	@Test
	public void test2(){
		int a[] = {1,4,6,5,9,10};
	    int[] resu = findSegment(a, 6);
	    Assert.assertTrue(resu[0] == 2 || resu[1] == 3);
	}
	
	@Test
	public void test3(){
		int a[] = {-1,9,2,10,0,8,12};
	    int[] resu = findSegment(a, 7);
	    Assert.assertTrue(resu[0] == 1 || resu[1] == 5);
	}
}
