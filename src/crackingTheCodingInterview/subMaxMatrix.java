package crackingTheCodingInterview;
import java.util.*;

import org.junit.Test;

/**
 * 最大的子矩阵之和
 * @author yefengzhichen
 * 2016年8月25日
 */
public class subMaxMatrix {
    public int sumOfSubMatrix(int[][] mat, int n) {
        // write code here
        
        int max = Integer.MIN_VALUE;   
        for(int i = 0; i < n; i++) {
            int[] a = new int[n];
            for(int j = i; j < n; j++) {
                int sum = 0;
                for(int k = 0; k < n; k++) {
                    a[k] += mat[j][k];
                    sum += a[k];
                    max = sum > max ? sum : max;
                    sum = sum < 0 ? 0 : sum;
                }           
            }
        }
        return max;
    }
    
    @Test
    public void test1() {
    	int[][] a = {{1,2,-3},{3,4,-5},{-5,-6,-7}}; //[[1,2,-3],[3,4,-5],[-5,-6,-7]],3
    	int resu = sumOfSubMatrix(a, 3);
    	System.out.println(resu);
    }
    
}