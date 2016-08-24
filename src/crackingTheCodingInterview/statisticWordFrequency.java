package crackingTheCodingInterview;

import org.junit.Assert;
import org.junit.Test;

/**
 * KMP 算法的两种实现方式，未优化next数组
 * @author yefengzhichen
 * 2016年8月24日
 */
public class statisticWordFrequency {
	
	public int kmp2(char[] str, int m, char[] word, int n) {
		int[] next = next2(word, n);
		int resu = -1;
		int k = 0;
		int i = 0;
		while(k < n && i < m){
			if(k == -1 || str[i] == word[k]){
				k++;				
				i++;	
			} else {
				k = next[k];
			}
		}
		if(k >= n){
			resu = i - n;
		}
		return resu;
	}
	
	public int[] next2(char[] a, int n){
		int[] next = new int[n];
		next[0] = -1;
		int k = -1;
		int i = 0;
		while(i < n - 1){
			if(k == -1 || a[i] == a[k]){
				k++;				
				i++;	
				next[i] = k; 
			} else {
				k = next[k];
			}
		}
		return next;
	}
	
	public int kmp1(char[] str, int m, char[] word, int n) {
		int[] next = next1(word, n);
		int k = 0;
		int resu = -1;
		for(int i = 0; i < m; i++) {
			while(k > -1 && word[k] != str[i]){
				k = next[k];
			}
			++k;
			if(k == n){
				resu = i - n + 1;
				break;
			}
		}
		return resu;
	}
	
	public int[] next1(char[] a, int n){
		int[] next = new int[n];
		next[0] = -1;
		int k = -1;
		for(int i = 1; i < n; i++) {
			while(k > -1 && a[k] != a[i - 1]){
				k = next[k];
			}
			next[i] = ++k;
		}
		return next;
	}
	
	@Test
	public void testKmp2(){
		String string = "accaaabaabcacaabc";
		int len1 =string.length();
		String word = "abaabcac";
		int len2 = word.length();
		int resu = kmp2(string.toCharArray(), len1, word.toCharArray(), len2);
		System.out.println(resu);
//		Assert.assertTrue(resu == 5);
	}
	
	@Test
	public void test2Next(){
		String string = "abaabc";
		char[] ch = string.toCharArray();
		int[] real = {-1, 0, 0, 1, 1, 2};
		int[] resu = next2(ch, 6);
		for(int i : resu){
			System.out.print(i + " ");
		}
		System.out.println();
//		Assert.assertArrayEquals(real, resu);
	}
	
	@Test
	public void testKmp1(){
		String string = "accaaabaabcacaabc";
		int len1 =string.length();
		String word = "abaabcac";
		int len2 = word.length();
		int resu = kmp1(string.toCharArray(), len1, word.toCharArray(), len2);
		System.out.println(resu);
		Assert.assertTrue(resu == 5);
	}
	
	@Test
	public void test1(){
		String string = "abaabc";
		char[] ch = string.toCharArray();
		int[] real = {-1, 0, 0, 1, 1, 2};
		int[] resu = next1(ch, 6);
		Assert.assertArrayEquals(real, resu);
		for(int i : resu){
			System.out.print(i + " ");
		}
		System.out.println();
	}
	
	@Test
	public void test2(){
		String string = "ababaca";
		char[] ch = string.toCharArray();
		int[] real = {-1, 0, 0, 1, 2, 3, 0};
		int[] resu = next1(ch, 7);
		Assert.assertArrayEquals(real, resu);
		for(int i : resu){
			System.out.print(i + " ");
		}
		System.out.println();
	}
	
	@Test
	public void test3(){
		String string = "abaabcac";
		char[] ch = string.toCharArray();
		int[] real = {-1, 0, 0, 1, 1, 2, 0, 1};
		int[] resu = next1(ch, 8);
		for(int i : resu){
			System.out.print(i + " ");
		}
		System.out.println();
		Assert.assertArrayEquals(real, resu);
		
	}
}
