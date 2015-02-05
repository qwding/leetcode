public class Solution {
    public static String intToRoman(int num) {
       
    	String[][] graph  = {{"I","II","III","IV","V","VI","VII","VIII","IX"},
    						   {"X","XX","XXX","XL","L","LX","LXX","LXXX","XC"},
    						   {"C","CC","CCC","CD","D","DC","DCC","DCCC","CM"},
    						   {"M","MM","MMM"} };
    	int sub = 10000;
    	String res_str = "";
    	while(sub > 1){
    		int temp = num % sub;
    		System.out.println("sub is: " + sub + " temp before is : " + temp);
    		temp = temp/(sub/10);
    		int idx_i = -1;
    		if(temp > 0){
    			switch(sub/10){
    				case 1000:
    					idx_i = 3;
    					break;
    				case 100:
    					idx_i = 2;
    					break;
    				case 10:
    					idx_i = 1;
    					break;
    				case 1:
    					idx_i = 0;
    					break;
    				default:
    					System.out.println("come in to idx");
    			}
    			System.out.println("idx_i is :" + idx_i + " wei shu is " + temp);
    			System.out.println("graph is : " + graph[idx_i][temp-1]);
    			res_str += graph[idx_i][temp - 1];	
    		}
    		sub = sub / 10;
    	}
    	System.out.println("result is: " + res_str);
    	return res_str;
    }
    public static void main(String[] args){
    	intToRoman(1613);
    }
}