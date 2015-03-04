import java.util.Arrays;

public class Solution {
    
	public static String longestCommonPrefix(String[] strs) {
        if(strs.length == 0){
        	return "";
        }
		if(strs.length == 1){
        	return strs[0];
        }
        Arrays.sort(strs);
        int longLength = strs[0].length();
        int temp_length = 0;
        for(int i = 0; i < strs[0].length() && i < strs[strs.length - 1].length(); ++ i){
        	if(strs[0].charAt(i) == strs[strs.length - 1].charAt(i)){
        		temp_length++;
        	}
        	if(strs[0].charAt(i) != strs[strs.length - 1].charAt(i)){
        		if(temp_length < longLength){
        		//	System.out.println(strs[0].substring(0,temp_length) + " " + temp_length);
        			return strs[0].substring(0,temp_length);
        		}
        		
        	}
        }
       // System.out.println(strs[0]);
        return strs[0];
    }
}