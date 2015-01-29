#include <iostream>

using namespace std;


class Solution {
public:
    double pow(double x, int n) {
    	double result = 1;
        if (n == 0){
        	return 1.0;
        }
        else if(n < 0){
        	result = 1.0/pow(x,-n);
        }
        result = pow(x,n/2);
        if (n % 2 != 0){
        	result *= x;
        }
        cout << "result is :" << result <<endl;
        return result;
    }
};

int main(){
	Solution solution;
	solution.pow(3.1,13);
	return 0;
}