package crackingTheCodingInterview;

import java.util.*;

import org.junit.Test;

public class subMaxAlphaMatrixs {
	public int findAlphaMatrix(String[] dic, int n) {
		// write code here
		Map<Integer, Map<Character, List<String>>> map = new HashMap<>();
		for (int i = 0; i < n; i++) {
			int key = dic[i].length();
			if (map.containsKey(key)) {
				char ch = dic[i].charAt(0);
				Map<Character, List<String>> inner = map.get(key);
				if (inner.containsKey(ch)) {
					List<String> list = inner.get(ch);
					list.add(dic[i]);
					inner.put(ch, list);
				} else {
					List<String> list = new ArrayList<>();
					list.add(dic[i]);
					inner.put(ch, list);
				}
				map.put(key, inner);
			} else {
				char ch = dic[i].charAt(0);
				Map<Character, List<String>> inner = new HashMap<>();
				List<String> list = new ArrayList<>();
				list.add(dic[i]);
				inner.put(ch, list);
				map.put(key, inner);
			}
		}
		int max = 0;
		for (Map.Entry<Integer, Map<Character, List<String>>> entry : map.entrySet()) {
			int len = entry.getKey();
			if (max < len) {
				Map<Character, List<String>> inner = entry.getValue();
				for (Map.Entry<Character, List<String>> same : inner.entrySet()) {
					boolean flag = true;
					char ch = same.getKey();
					List<String> list = same.getValue();
					for (String str : list) {
						for (int i = 1; i < str.length(); i++) {
							char first = str.charAt(i);
							if (inner.get(first) == null) {
								flag = false;
								break;
							}
						}
						if (flag) {
							max = len > max ? len : max;
							break;
						}
					}
					if (flag) {
						break;
					}
				}
			}
		}
		return max * max;
	}

	@Test
	public void test1() {
		String[] aStrings = { "aaa", "aaa", "aaa", "bb", "bb" }; // ["aaa","aaa","aaa","bb","bb"]
		int resu = findAlphaMatrix(aStrings, 5);
		System.out.println(resu);
	}

}
