import java.util.ArrayList;

public class Solution {
    public static int lengthOfLongestSubstring(String s) {

    	int length = s.length();
    	if(length <= 1){
    		System.out.println("res: " + length);
    		return length;
    	}
        ArrayList<Character> list = new ArrayList(); 
        int res = 0;
        int cur = 0;
        for(int i = 0; i < length; ++i){
        	list.clear();
        	list.add(s.charAt(i));
        	cur = 1;
        	for(int j = i + 1; j < length; j++){
        		if(!list.contains(s.charAt(j))){
        			list.add(s.charAt(j));
        			cur++;
        		}
        		else{
        			res = res > cur? res:cur;
        			
        		}
        	}
        }
        System.out.println("res: " + res);
        return res;
    }
    public static void main(String[] args){
    	String s = "";
    	lengthOfLongestSubstring(s);
    }
}