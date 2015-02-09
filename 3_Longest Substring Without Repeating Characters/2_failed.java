import java.util.HashMap;

public class Solution {
    public static int lengthOfLongestSubstring(String s) {

    	int length = s.length();
    	if(length <= 1){
    		System.out.println("res: " + length);
    		return length;
    	}
        HashMap<Character, Integer> map = new HashMap<Character, Integer>(); 
        int res = 0;
        int cur_2 = 0;
        int cur_1 = 0;
        int cur = 0;
        int cur_start = 0;
        for(int i = 0; i < length; ++i){
        	if(i == length -1 & !map.containsKey(s.charAt(i))){
        		System.out.println("end length : " + length);
        		System.out.println("end i : " + i);
        		System.out.println("end cur start : " + cur_start);
        		cur = length - cur_start;
        		System.out.println("end cur : " + cur);
        		res = res > cur? res : cur;
        		System.out.println("end res: " + res);
        	}
        	if(map.containsKey(s.charAt(i))){
        		System.out.println("map : " + map.get(s.charAt(i)));
        		cur_2 = i- map.get(s.charAt(i)) + 1;
        		cur_1 = map.get(s.charAt(i)) - cur_start + 1;
        		System.out.println("map cur_2 : " + cur_2);
        		System.out.println("map cur_1 : " + cur_1);
        		cur = cur_1 > cur_2 ? cur_1 : cur_2;
        		cur_start = map.get(s.charAt(i)) + 1;
        		map.clear();
        		
        		res = res > cur? res : cur;
        		System.out.println("map : " + res);
        	}
        	else{
        		map.put(s.charAt(i), i);
        	}
        }
        	
        System.out.println("res: " + res);
        return res;
    }
    public static void main(String[] args){
    	String s = "abcabcbb";
    	lengthOfLongestSubstring(s);
    }
}