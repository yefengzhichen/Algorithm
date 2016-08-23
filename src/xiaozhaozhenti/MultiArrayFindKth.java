package xiaozhaozhenti;

import org.junit.Assert;
import org.junit.Test;


/**
 * ��������������ĵ�K������
 * ʱ�临�Ӷ�Ϊlog(m + n), �ȹ鲢����ʽ�úܶ�
 * @author yefengzhichen
 * 2016��8��23��
 */
public class MultiArrayFindKth {

	public int getKth(int[] a, int i, int m, int[] b, int j, int n, int k){
		if (m - i > n -j) {
			return getKth(b, j, n, a, i, m, k);
		}
		if (m == i) {
			return b[j + k - 1];
		}
		if (k == 1) {
			return Math.min(a[i], b[j]);
		}
		int pa = Math.min(k/2, m - i);
		int pb = k - pa;
		if (a[i + pa - 1] == b[j + pa - 1]) {
			return a[i + pa - 1];
		} else if (a[i + pa - 1] < b[j + pa - 1]) {
			return getKth(a, i + pa, m, b, j, n, k - pa);
		} else {
			return getKth(a, i, m, b, j + pb, n, k - pb);
		}
	} 
	
	@Test
	public void test1(){
		int a1[] = {1, 12, 15, 26, 38};
	    int a2[] = {2, 13, 17, 30, 45};
	    int resu = getKth(a1, 0, 5, a2, 0, 5, 5);
	    Assert.assertTrue(resu == 15);
	}
	
	@Test
	public void test2(){
		int a1[] = {1, 12, 15, 26, 38};
	    int a2[] = {2, 13, 17, 30, 45};
	    int resu = getKth(a1, 0, 5, a2, 0, 5, 1);
	    Assert.assertTrue(resu == 1);
	}
	
	@Test
	public void test3(){
		int a1[] = {1, 12, 15, 26, 38};
	    int a2[] = {2, 13, 17, 30, 45};
	    int resu = getKth(a1, 0, 5, a2, 0, 5, 9);
	    Assert.assertTrue(resu == 38);
	}
	
	@Test
	public void test4(){
		int a1[] = {1, 12, 15, 26, 38};
	    int a2[] = {2, 13, 17, 30, 45};
	    int resu = getKth(a1, 0, 5, a2, 0, 5, 10);
	    Assert.assertTrue(resu == 45);
	}
}
