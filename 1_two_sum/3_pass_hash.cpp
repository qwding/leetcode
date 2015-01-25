#include <stdio.h>
#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

//此方法为先排序，头尾分别取i,j，两边和过大，则--j,两边和过小，++i,写的方法笔较笨，望优化.

class Solution {
public:
    vector<int> twoSum(vector<int> &numbers, int target) {
       unordered_map<int, int> num_idx_map;
       for(int i = 0; i < numbers.size(); ++i){
        num_idx_map[numbers[i]] = i;
       }
       for(int i = 0; i < numbers.size(); ++i){
           unordered_map<int, int>::iterator iter;
           iter = num_idx_map.find(target - numbers[i]);
           if(iter != num_idx_map.end() && i != iter -> second){
                return {i+1 , (iter -> second + 1)};
           }
        }
        return {0,0};
    }
};


int main()
{
    Solution solution;
    vector<int> numbers = {3,2,4};
    int target = 6;
    vector<int> num = solution.twoSum(numbers,target);
    cout << num[0] << endl;
    cout << num[1] << endl;
	return 0;
}
