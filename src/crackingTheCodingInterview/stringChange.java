package crackingTheCodingInterview;

import java.util.*;

import org.junit.Test;


/**
 * 字符串转换，从一个字符串到另一个字符串，中间状态为给定字典中的。
 * @author yefengzhichen
 * 2016年8月25日
 */
public class stringChange {
	public int countChanges(String[] dic, int n, String s, String t) {
		// write code here
		boolean[] access = new boolean[n];
		List<String> list = new ArrayList<>();
		list.add(s);
		boolean flag = false;
		int resu = 0;
		//和原始串相同
		for (int j = 0; j < n; j++) {
			if (dic[j].equals(s)) {
				access[j] = true;
			}
		}
		while (!flag) {
			int lastSize = list.size();
			List<String> newList = new ArrayList<>();
			for (int i = 0; i < lastSize; i++) {
				String inter = list.get(i);
				if (inter.equals(t)) {
					flag = true;
					break;
				}
				for (int j = 0; j < n; j++) {
					if (!access[j] && isDisOne(dic[j], inter)) {
						access[j] = true;
						newList.add(dic[j]);
					}
				}
			}
			if(!flag){
				if(newList.size() == 0) {
					resu = -1;
					flag = true;
				} else {
					resu++;
					list = newList;
				}
			}			
		}
		return resu;
	}

	public boolean isDisOne(String str1, String str2) {
		int len = str1.length();
		int resu = 0;
		for (int i = 0; i < len; i++) {
			if (str1.charAt(i) != str2.charAt(i)) {
				resu++;
			}
		}
		if (resu == 1) {
			return true;
		}
		return false;
	}
	
	@Test
	public void test1() {
		String[] dic = {"abc","adc","bdc","aaa"}; //["abc","adc","bdc","aaa”],4,”abc","bdc"
		int resu = countChanges(dic,4,"abc","bdc");
		System.out.println(resu);
	}
	
}
