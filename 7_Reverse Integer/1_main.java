
public class Solution {
	public static int reverse(int x){
		int result = 0;
		int resource = x;
		long first_result = 0;
		while(x != 0){
			//int temp = 
			first_result = first_result * 10 + x % 10;;
			x = x / 10;
		}
		if ((resource > 0 & first_result > 2147483647 ) | (resource < 0 & first_result < -2147483648)){
				return 0;	
		}
		//���������Ҫע�����Ҫ��101100��������reverse��Ҫע������ȷ��ʾ������reverse���ܹ��硣
		result = (int)first_result;
		return result;
	}
	public static void main(String[] args){
		int xyz = reverse(-1534236469);
		System.out.println("at last " + xyz);
		
	}
}