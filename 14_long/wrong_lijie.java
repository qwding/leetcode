import java.util.Arrays;
  //求的是字符串间最长的前缀
public class Solution {
    
	public static String longestCommonPrefix(String[] strs) {
        Arrays.sort(strs);
        String res = "";
        int longLength = 0;
        if(strs.length == 1){
        	return strs[0];
        }
        for(int i = 0; i < strs.length && i+1 <strs.length; ++i){
        	int tempLength = 0;
        	for(int j = 0; j < strs[i].length() && j < strs[i+1].length(); ++j){
        		if(strs[i].charAt(j) == strs[i + 1].charAt(j)){
        			tempLength++;
        		}
        	}
        	if(tempLength > longLength){
				System.out.println(strs[i] + "  " + strs[i+1] + "   " + tempLength);
				res = strs[i].substring(0,tempLength);
				longLength = tempLength;
			}
        }
        System.out.println(res);
        return res;
    }
	public static void main(String[] args){
		String[] strs = {"abcdefg","abcdefh","abcde"};
		longestCommonPrefix(strs);
	}
}