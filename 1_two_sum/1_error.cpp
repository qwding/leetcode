#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int> &numbers, int target) {
    	for(int i = numbers.size()-1; i >= 0 ; i--){
        	int sub = target - numbers[i];
            printf("sub num is %d\n",sub);
        	if (sub > 0)
        		for(int j = 0; j < i; j++){
        			if (sub == numbers[j]){
        				return {j+1,i+1};
                    }
        	}
    	}
        return {0,0};
    }
};
//zhushi

int main()
{
    Solution solution;
    vector<int> numbers = {1,3,2,11,5,15};
    int target = 9;
    vector<int> num = solution.twoSum(numbers,target);
    cout << num[0] << endl;
    cout << num[1] << endl;
	return 0;
}
