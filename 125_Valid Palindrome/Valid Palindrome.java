public class Solution {
    public static boolean isPalindrome(String s) {
    	if(s.length() == 0)
    		return true;
        int i = 0;
        int j = s.length() - 1;
        while(i < j){
        	if(!charRight(s.charAt(i)))
        		++i;
        	else if(!charRight(s.charAt(j))){
        		--j;
        	}
        	else{
        		System.out.println("i " + s.charAt(i) + " j " + s.charAt(j));
        		if(s.charAt(i) <= 57 || s.charAt(j) <= 57){
        			if(s.charAt(i) == s.charAt(j)){
        				i++;
        				j--;
        			}
        			else
        				return false;
        		}
        		else if((s.charAt(i) == s.charAt(j)) || ((s.charAt(i) - 32) == s.charAt(j)) || ((s.charAt(i) + 32) == s.charAt(j))){
        			i++;
        			j--;
        		}
        		else 
        			return false;
        	}
        }
        return true;
    }
    public static boolean charRight(char c){
    	if((c >= 48 && c <= 57)||( c >=65 && c <= 90) || ( c >= 97 && c <=122)){
    		return true;
    	}
    	return false;
    }
    public static void main(String[] args){
    	String s = "bc";
    	System.out.println(isPalindrome(s));
    }
}