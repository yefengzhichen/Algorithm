package crackingTheCodingInterview;

import org.junit.Assert;
import org.junit.Test;

/**
 * ���ַ���,��������Ӣ�ı�ʾ
 * @author yefengzhichen
 * 2016��8��23��
 * ��һ���Ǹ����������дһ���㷨����ӡ��������Ӣ������������һ��int x���뷵��һ��string��Ϊ��������Ӣ��������
 * ����������1234���أ�"One Thousand,Two Hundred Thirty Four"
 */
public class intToEnglishString {
	public String toString(int x) {
        // write code here
        String[] gebai = {"One", "Two", "Three","Four","Five","Six","Seven","Eight","Nine"};
        String[] tennum = {"Ten","Eleven","Twelve","Thirteen","Fourteen","Fifteen","Sixteen","Seventeen","Eighteen","Nineteen"};
        String[] shiwei = {"Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"};
        String[] thoud = {"Billion","Million","Thousand"};
        int base = (int)Math.pow(10,9);
        int thoudIndex = 0;
        String resu = "";
        while(x > 0) {
            int temp1 = x / base;
            if(temp1 > 0) {
                int bai = temp1 / 100;
                resu += ",";
                if(bai > 0) {
                    resu += (" " + gebai[bai - 1] + " Hundred");         
                }
                int temp2 = temp1 % 100;
                if(temp2 < 10) {
                    resu += (" " + gebai[temp2 - 1]);  
                } else if (temp2 < 20) {
                    resu += (" " + tennum[temp2 - 10]);  
                } else {
                    int shi = temp2 / 10;
                    resu += (" " + shiwei[shi - 2]);  
                    int ge = temp2 % 10;
                    if(ge > 0) {
                        resu += (" " + gebai[ge - 1]);  
                    }
                }
                if(thoudIndex < 3){
                	resu += (" " + thoud[thoudIndex]);  
                }               
            }
            thoudIndex ++;
            x = x % base;
            base = base / 1000;         
        }
        return resu.substring(2).replace(", ", ",");
    }
	
	@Test
	public void test1(){
		String real = "One Thousand,Two Hundred Thirty Four"; //One Thousand,Two Hundred Thrity Four
		String resu = toString(1234);
		System.out.println(resu);
		Assert.assertTrue(resu.equals(real) );
	}
	
	@Test
	public void test2(){
		String real = "One Million"; //One Thousand,Two Hundred Thrity Four 1413498
		String resu = toString(1000000);
		System.out.println(resu);
		Assert.assertTrue(resu.equals(real) );
	}
	
	@Test
	public void test3(){
		String real = "One Million,Four Hundred Thirteen Thousand,Four Hundred Ninety Eight"; //One Thousand,Two Hundred Thrity Four 1413498
		String resu = toString(1413498);
		System.out.println(resu);
		Assert.assertTrue(resu.equals(real) );
	}
}
