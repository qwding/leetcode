import java.util.ArrayList;
import java.util.List;

public class Solution {
     static char[][] num = {{' '},
			{'@'},
			{'a','b','c'},
			{'d','e','f'},
			{'g','h','i'},
			{'j','k','l'},
			{'m','n','o'},
			{'p','q','r','s'},
			{'t','u','v'},
			{'w','x','y','z'}};
    public static List<String> letterCombinations(String digits) {
    	List<String> list = new ArrayList<String>();
    	
    	if(digits.length() == 0){
    		list.add("");
    		return list;
    	}
        
        String s = new String("");
        get_res(digits, 0, s, list);
        System.out.println(list);
        return list;
    }
    
    public static void get_res(String digits,int d_idx, String s, List<String> list){
    	for(int i = 0; i < num[digits.charAt(d_idx) - '0'].length;++i){
    		if(d_idx < digits.length() - 1){
    			get_res(digits,d_idx + 1, s+num[digits.charAt(d_idx) - '0'][i],list);
    		}
    		else if(d_idx == digits.length() - 1){
    			list.add(s + num[digits.charAt(d_idx) - '0'][i]);
    		}
    		else{
    			System.out.println("come into unknow place");
    		}
    	}
    }
    public static void main(String[] args){
    	String s = new String("");
    	letterCombinations(s);
    	
    }
}