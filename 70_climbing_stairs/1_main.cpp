#include <iostream>

using namespace std;

class Solution {
public:
    int climbStairs(int n) {
    	/*
    	if (n < 0)
    		return 0;
    	if (n <= 1)
    		return 1;
    	else
    		return climbStairs(n - 1) + climbStairs(n - 2); //超时!!!
    		*/
    	if(n <= 2){
    		return n;
    	}
    	int before_1 = 1;
    	int before_2 = 2;
    	int current = 0;
    	for(int i = 2; i < n; ++i){
    		current = before_1 + before_2;
    		before_1 = before_2;
    		before_2 = current;
    	}
    	return current;

    }
};

int main(){
	Solution solution;
	cout <<"last result: " <<solution.climbStairs(1) << endl;
	return 0;
}