#include <iostream>

using namespace std;


class Solution {
public:
    double pow(double x, int n) {
        if (n == 0){
        	return 1.0;
        }
        else if(n < 0){
        	double half = pow(x,-n/2);
            if(n%2 == 0) return 1.0/(half * half);
            else return 1.0/(half * half *x);
        }
        else{
            double result = pow(x,n/2);
            if (n % 2 != 0){
                cout << "result is :" << result * result * x <<endl;
                return result * result * x;
            }
            else{
                cout << "result is :" << result * result <<endl;
                return result * result;
            }
        }
 

    }
};

int main(){
	Solution solution;
	solution.pow(1.0000,-2147483648);
	return 0;
}