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
    	}//д��ʱ��һֱ�������������һ���ݹ��������д�ݹ��ʱ�����Ѿ��������һ���ݹ�������е��𣬾��ǵڶ������е�ʱ���ǻ��ڵ�һ�����еĽ���������Һ����ˣ�һ����Ҫ���⣬
    	//���ǵݹ�����޷���ֵ����ô����������ôд��֮ǰ����ĵݹ鶼����Ӱ�����ģ��������Ǿ��ǲ��еĽ���ˡ��ö�
    	//��������Ĺؼ���ʣ�� �����������ǲ��ܴ������������ġ�
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