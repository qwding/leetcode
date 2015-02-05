public class Solution {
    public static int romanToInt(String s) {
        int before = 0;
        int result = 0;
        
        for(int i = s.length() - 1; i >= 0; --i){
        	int current = revertNum(s.charAt(i));
        	if((current == 1 & (before == 5 | before == 10)) | (current == 10 & (before == 50 | before == 100)) | (current == 100 & (before == 500 | before == 1000))){
        		result = result - current;
        	}
        	else{
        		result = result + current;
        	}
        	before = current;
        }
        System.out.println("at last : " + result);
        return result;
    }
    public static int revertNum(char a){
    	switch(a){
    	case 'I':
    		return 1;
    	case 'V':
    		return 5;
    	case 'X':
    		return 10;
    	case 'L':
    		return 50;
    	case 'C':
    		return 100;
    	case 'D':
    		return 500;
    	case 'M':
    		return 1000;
    		default:
    			return 0;
    	}
    }
    public static void main(String[] args){
    	romanToInt("MMMCMXCIX");
    }
}