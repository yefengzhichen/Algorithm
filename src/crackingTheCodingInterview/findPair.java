package crackingTheCodingInterview;

import java.util.*;
import org.junit.Assert;
import org.junit.Test;

/**
 * 整数对查找
 * @author yefengzhichen
 * 2016年8月24日
 * 请设计一个高效算法，找出数组中两数之和为指定值的所有整数对。
 * 给定一个int数组A和数组大小n以及需查找的和sum，请返回和为sum的整数对的个数。保证数组大小小于等于3000.
 * 测试样例：[1,2,3,4,5],5,6 返回：2
 */
public class findPair {
    public int countPairs(int[] A, int n, int sum) {
        // write code here
        Map<Integer, Integer> map = new HashMap<>();
        int resu = 0;
        for(int i = 0; i < n; i++) {
            Integer value = map.get(A[i]);
            value = value == null ? 1 : value + 1;
            map.put(A[i], value);
        }
        for(int i = 0; i < n; i++) {
            int key = sum - A[i];
            if(key == A[i]) {
                if(map.get(key) > 1){
                    int temp = map.get(key);
                    resu = resu + (temp - 1) * temp / 2; 
                    map.put(key, 0);
                }
            } else {
                if(map.get(key) !=null && map.get(key) > 0) {
                    int temp1 = map.get(key);
                    int temp2 = map.get(A[i]);
                    resu += (temp1 * temp2); 
                    map.put(A[i], 0);
                    map.put(key, 0);
                }
            }
        }
        return resu;
    }
    
    @Test
    public void test1(){
    	int[] a = {1,2,3,4,5};
    	int resu = countPairs(a, 5, 6);
    	System.out.println(resu);
    	Assert.assertTrue(resu == 2);
    }
    
    @Test
    public void test2(){
    	int[] a = {11,7,7,6,14,2,14,15,2,1,2,12,13,9,8,15,13,8,10,11,14,10,2,9,4,9,3,7,6,10,15,4,7,6,15,3,9,13,5,2,6,10,10,1,12,4,3,3,8,8,1,4,7,11,13,5,13,15,4,3,1,11,6,11,9,9,11,15,12,10,13,3,11,4,8,9,7,3,13,9,11,3,2,11,10,1,4,2,3,3,14,11,5,10,1,14,8,1,11,3,1,9,14,6,1,7,15,10,14,6,4,12,11};
    	int resu = countPairs(a, 113, 16);
    	System.out.println(resu);
    	Assert.assertTrue(resu == 398);
    }
}
