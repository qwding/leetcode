public class Solution {
    public static int maxArea(int[] height) {
        int max_area = 0;
        int i = 0;
        int j = height.length - 1;
        while(i < j){
        	if(height[i] >= height[j]){
            	int temp_max_area = (j -i) * height[j];
            	if(temp_max_area > max_area){
            		max_area = temp_max_area;
            	}
            	--j;
        	}
        	else{
        		int temp_max_area = (j -i) * height[i];
        		if(temp_max_area > max_area){
            		max_area = temp_max_area;
            	}
        		++i;
        	}
        }
        
        System.out.println(max_area);
        return max_area;
    }
    public static void main(String args[]){
    	int[] array =  {1,3,6,4,5};
    	maxArea(array);
    }
}