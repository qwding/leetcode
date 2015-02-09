import java.util.Iterator;
import java.util.List;
import java.util.ArrayList;

public class Solution {
    public static List<String> generateParenthesis(int n) {
        List<String> list = new ArrayList<String>();
        String s = "";
        get_s(n, n, s, list);
        for(Iterator iterator = list.iterator();iterator.hasNext();){
        	System.out.println("result is : " + (String)iterator.next());
        }
        return list;
    }
    public static void get_s(int left_num,int right_num,String s, List<String> list){
    	if(left_num == 0 & right_num == 0){
    		list.add(s);
    	}//写的时候一直在想的是跳入了一个递归后，下面在写递归的时候不是已经运行完第一个递归后再运行的吗，就是第二次运行的时候是基于第一次运行的结果。但是我忽略了；一个重要问题，
    	//就是递归可以无返回值，那么不管下面怎么写，之前进入的递归都不会影响后面的，就是他们就是并行的结果了。好二
    	//但是这题的关键是剩余 的左括号数是不能大于右括号数的。
    	if(left_num > 0){
    		get_s(left_num -1, right_num,s + "(", list);
    	}
    	if(right_num > 0 & right_num > left_num){
    		get_s(left_num, right_num -1, s + ")", list);
    	}
    }
    public static void main(String[] args){
    	generateParenthesis(3);
    }
}