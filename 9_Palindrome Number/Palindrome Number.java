public class Solution {
	public static boolean isPalindrome(int x){
		int before = x;
		long after = 0;
		while(x > 0){
			after = after * 10 + x%10;
			x = x / 10;
		}
		if(after == before){
			System.out.println("yes");
		}
		else
			System.out.println("no");
		return after == before;
	}
	public static void main(String[] args){
		isPalindrome(1123211);
		
	}
}