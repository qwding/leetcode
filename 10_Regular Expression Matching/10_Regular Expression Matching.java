public class Solution {
    public boolean isMatch(String s, String p) {
         int sLen = s.length();
         int pLen = p.length();
         if (s == null)
             return p == null;
         if(pLen == 0){
        	 return sLen == 0;
         }
         if(pLen == 1){
        	 if(p.equals(s) || (p.equals(".") && s.length() == 1)){   //开始这里的p.equals(".") 写成了p.equals(',')结果一直不过。很惨，终于知道了java里字符和string用引号的严重
        		 return true;
        	 }	 
        	 else{
        		 return false;
        	 }
        		
         }
         if(p.charAt(1) != '*'){
        	 if (s.length() > 0
                     && (p.charAt(0) == s.charAt(0) || p.charAt(0) == '.')) {
        		 return isMatch(s.substring(1), p.substring(1));
        	 }
        	 return false;
         }
         else{
        	 while(s.length() > 0 && (s.charAt(0) == p.charAt(0) | p.charAt(0) == '.')){
        		 if(isMatch(s, p.substring(2)))
        			 return true;
        		 s = s.substring(1);
        	 }
        	 return isMatch(s, p.substring(2));
         }
    }
    public static void main(String[] args) {
        System.out.println(new Solution().isMatch("aa", "a"));
        System.out.println(new Solution().isMatch("aa", "aa"));
        System.out.println(new Solution().isMatch("aaa", "aa"));
        System.out.println(new Solution().isMatch("aa", "a*"));
        System.out.println(new Solution().isMatch("aa", ".*"));
        System.out.println(new Solution().isMatch("ab", ".*"));
        System.out.println(new Solution().isMatch("aab", "c*a*b"));
        
        System.out.println(new Solution().isMatch("", ""));
        System.out.println(new Solution().isMatch("abcdeff", ".*"));
        System.out.println(new Solution().isMatch("ab", ".*c"));
       System.out.println(new Solution().isMatch("bb", ".bab"));

   }

}