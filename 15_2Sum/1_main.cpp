#include <iostream>
#include <vector>
#include <unordered_map>
#include <iterator>
//费劲巴力，结果还是timeout了，真是伤心死我了！！！！

using namespace std;

class Solution {
public:
    vector<vector<int> > threeSum(vector<int> &num) {
        unordered_map<int, int> num_map;
        vector<vector<int> > result;
        int size = num.size();
        sort(num.begin(),num.end());
        printf("sizeof array: %d\n", size);
        if(size <= 2)
        	return {};
        for(int i = 0; i < size; ++i){
        	num_map[num[i]] = i;
    	}
    	for(int i = 0; i < size ; ++i){
    		for(int j = i + 1; j < size; ++j){
    			unordered_map<int, int>::iterator iter;
    			iter = num_map.find((num[i] + num[j])*-1);
    			if(iter != num_map.end() && iter -> second > j ){
    				result.push_back({num[i],num[j],iter -> first});
    				printf("value is :%d %d %d\n",num[i],num[j],iter -> first);
    				//printf("key is :%d %d %d\n",i,j,iter -> second);
    			}
    		}
    	}
    	return result;
    }
};

int main(){
	vector<int> numbers2 = {3,0,-2,-1,1,2};
	vector<int> numbers = {-1, 0, 1, 2, -4, 3,5,7};
	Solution solution;
	solution.threeSum(numbers2);
	return 0;
}