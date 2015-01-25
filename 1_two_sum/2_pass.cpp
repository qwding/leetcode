#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

//此方法为先排序，头尾分别取i,j，两边和过大，则--j,两边和过小，++i,写的方法笔较笨，望优化.

class Solution {
public:
    vector<int> twoSum(vector<int> &numbers, int target) {
        int size = numbers.size();
        
        vector<int> temp_array = numbers;
        sort(numbers.begin(),numbers.end());
        int i = 0;
        int j = size - 1;
        while(i < j){
            while(i < j && numbers[i] + numbers[j] > target){
                --j;
            }
            while(i < j && numbers[i] + numbers[j] <target){
                ++i;
            }
            if(numbers[i] + numbers[j] == target){
                int return_i = -1;
                int return_j = -1;
                for(int index = 0; index < temp_array.size(); ++index){
                    if(temp_array[index] == numbers[i] && return_i == -1)
                        return_i = index + 1;
                    else if(temp_array[index] == numbers[j] && return_j == -1)
                        return_j = index + 1;
                }
               return (return_i < return_j) ? {return_i,return_j} : {return_j,return_i};
            }
        }
        return {0,0};
    }
};


int main()
{
    Solution solution;
    vector<int> numbers = {0,4,3,0};
    int target = 0;
    vector<int> num = solution.twoSum(numbers,target);
    cout << num[0] << endl;
    cout << num[1] << endl;
	return 0;
}
