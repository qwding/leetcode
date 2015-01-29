#include <iostream>
#include <stack>
#include <string>

using namespace std;

class Solution {
public:
    bool isValid(string s) {
       stack<char> bra_stack;
       int s_len = s.length();
       for(int i = 0; i < s_len; ++i){
       		if(bra_stack.empty() || (bra_stack.top() == '[' && s[i] != ']') || (bra_stack.top() == '{' && s[i] != '}') || (bra_stack.top() == '(' && s[i] != ')')){
       			bra_stack.push(s[i]);
       		}
       		else 
       			bra_stack.pop();
       }
       int result =  bra_stack.empty();
       cout << result <<endl;
       return result;
    }
};

int main(){
	string str = "[]{}({}))";
	Solution solution;

	return solution.isValid(str);

}