package crackingTheCodingInterview;

import java.util.*;

import org.junit.Test;

class Node {
    Character ch;
    Map<Character, Node> map;
    
    public Node(Character ch) {
        this.ch = ch;
        map = new HashMap<Character, Node>();
    }
    
    public boolean containsKey(Character ch){
        return map.containsKey(ch);
    }
    
    public Node get(Character ch){
        return map.get(ch);
    }
    
    public Node put(Character ch, Node node){
        return map.put(ch, node);
    }
}

public class multiStringIsSubString { 
    public boolean[] chkSubStr(String[] p, int n, String s) {
        // write code here
        Node root = new Node('0');
        List<Node> list = new ArrayList<>();
        int len = s.length();
        for(int i = 0; i < len; i++) {
            int oldSize= list.size();
            char key = s.charAt(i);
            List<Node> newList = new ArrayList<>();
            for(int j = 0; j < oldSize; j++) {
               Node node = list.get(j);
               if(node.containsKey(key)) {
                   newList.add(node.get(key));
               } else {
                   Node child = new Node(key);
                   node.put(key, child);
                   newList.add(node.get(key));
               }
            }
            if(root.containsKey(key)){
                newList.add(root.get(key));
            } else {
                Node child = new Node(key);
                root.put(key, child);
                newList.add(root.get(key));
            }
            list = newList;
        }
        boolean[] resu = new boolean[n];
        for(int i = 0; i < n; i++) {
            String str = p[i];
            Node node = root;
            boolean flag = true;
            for(int j = 0; j < str.length(); j++) {
                char key = str.charAt(j);
                if(node.containsKey(key)) {
                    node = node.get(key);
                } else {
                    flag = false;
                    break;
                }
            }
            resu[i] = flag;
        }
        return resu;
    }
    
    @Test
    public void test1(){
    	String[] strings = {"a","b","c","d"};//["a","b","c","d"],4,"abc"
    	boolean[] resu = chkSubStr(strings,4,"abc");
    	for(int i = 0; i < resu.length; i++) {
    		System.out.print(resu[i]+ " ");
    	}
    }
    
}