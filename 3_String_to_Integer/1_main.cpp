#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int atoi(const char *str) {
    	vector<int> temp;
    	int symbol = 1;
    	int i = 0;
    	long long result = 0;
    	while(str[i] == ' '){
    		++i;
    	}
    	if(str[i] == '-'){
    		symbol = -1;
    		i++;
    	}
    	else if( str[i] == '+'){
    		symbol = 1;
    		i++;
    	}
    	printf("symbol is : %d\n",symbol);
    	while(str[i] != '\0'){
    		if ( (str[i] - '0') >= 0 && (str[i] - '0') <= 9){
    			temp.push_back( str[i] - '0');
    		}
    		else
    			break;
    		++i;
    	}
    	if(temp.empty()){
    		return 0;
    	}
    	else{
    		int n = temp.size();
    		printf("size of temp is: %d\n",n);
    		for(int k = 0; k < temp.size(); ++k){
    			result = result*10+temp[k];
    			if(symbol == 1 && result > 2147483647)
    				return 2147483647;
    			if(symbol == -1 && result >2147483648)
    				return -2147483648;
    		}	
    	}
    	result = result * symbol;
      	return result;
    }
};


int main(){
	char str[20] =" -10522545459";
	Solution solution;
	int num = solution.atoi(str);
	printf("%d\n",num);
	//int num2 = '223434535555535353552' - '11';
	//printf("string sub is %d\n",num2);
	return 0;
}