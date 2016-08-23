package crackingTheCodingInterview;

import org.junit.Assert;
import org.junit.Test;

/**��С��������
 * @author yefengzhichen
 * 2016��8��23��
 * ��һ���������飬���дһ���������ҳ�����m��n��ֻҪ��m��n֮���Ԫ���ź������������������ġ�
 * ע�⣺n-mӦ��ԽСԽ�ã�Ҳ����˵���ҳ������������������.
 * ����һ��int����A������Ĵ�Сn���뷵��һ����Ԫ�飬�����������е������յ㡣
 * (ԭ����λ�ô�0��ʼ���,��ԭ�������򣬷���[0,0])����֤A��Ԫ�ؾ�Ϊ��������
 * ����������[1,4,6,5,9,10],6���أ�[2,3]
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
