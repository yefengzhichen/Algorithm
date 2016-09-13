package crackingTheCodingInterview;

import java.util.*;

import org.junit.Test;

public class longestCombineSubString {
	public int getLongest(String[] str, int n) {
        // write code here
        Map<Character, List<String>> map = new HashMap<>();
        for(int i =0 ; i < n; i++) {
            String s = str[i];
            char key = s.charAt(0);
            if(map.containsKey(key)) {
                List<String> list = map.get(key);
                list.add(s);
                map.put(key, list);
            } else {
                List<String> list = new ArrayList<>();
                list.add(s);
                map.put(key, list);
            }
        }
        Comparator<String> cmp = new Comparator<String>(){
            public int compare(String str1, String str2){                
                    Integer len1 = str1.length();
                    Integer len2 = str2.length();
                    return len1.compareTo(len2);
            }
        };
        Arrays.sort(str, cmp);
        boolean flag = false;
        int resu = 0;
        for(int i = n - 1; i >= 0; i--) {
            String one = str[i];
            int len = one.length();
            flag = getLong(one, map, 0, len);
            if(flag) {
                resu = one.length();
            	break;
            }
        }
        return resu;        
    }
    
    public boolean getLong(String str, Map<Character, List<String>> map, int index, int len) {
        if(index > len) {
            return false;
        } else if(index == len) {
            return true;
        }
        char key = str.charAt(index);
        boolean flag = false;
        if(map.containsKey(key)) {
            List<String> list = map.get(key);
            for(int i = 0; i < list.size(); i++) {
                String t = list.get(i);
                if(!t.equals(str) && isSame(str, t, index)) {
                    int mlen = t.length();
                    flag = flag || getLong(str, map, index + mlen, len);
                } 
                if(flag){
                    break;
                }
            }
        }
        return flag;
    }
    
    public boolean isSame(String str, String t, int index){
        int len = t.length();
        if(str.length() - index < len){
            return false;
        }
        if(t.equals(str.substring(index, index + len))){
            return true;
        }
        return false;
    }
	
    @Test
    public void test1(){
    	String[] strings =  {"a","b","c","ab","bc","abc"};//;["a","b","c","ab","bc","abc"],6
    	int resu = getLongest(strings,6);
    	System.out.println(resu);
    }
    //["rb","erb","rbe","e","g","mcxr"],6
    @Test
    public void test2(){
    	String[] strings =  {"rb","erb","rbe","e","g","mcxr"};//;["a","b","c","ab","bc","abc"],6
    	int resu = getLongest(strings,6);
    	System.out.println(resu);
    }
}
