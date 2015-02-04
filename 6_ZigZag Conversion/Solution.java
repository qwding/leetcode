
public class Solution {
	public static String convert(String s, int nRows){
		int length = s.length();
		if(nRows == 1 | nRows > length){
			return s;
		}
		System.out.println("length is : " + length);
		int block = nRows + nRows - 2;		
		String str = "";
		int row = 1;
		while(row <= nRows){
			int num = 0;
			while(num < length){
				System.out.println("num : " + num + "  row: " + row);
				if((num + row - 1) < length ){
					str =str + s.charAt(num + row - 1); 
				}
				if(row != 1 && row != nRows &&(num + block + 1 - row) < length){
					str = str + s.charAt(num + block + 1 - row);
				}
				num = num + block;
			}
			row++;
		}
		System.out.println("result:" + str);
		return str;
	}
	public static void main(String args[]){
		String str = "rmbjwpglakqvchvzvshicnqdluqgwqdnceyywglwqetunotigasjqjoddgkzw";
		convert(str, 10);
		System.out.println("mani string:" + str);
	}
}
