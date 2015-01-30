#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    void setZeroes(vector<vector<int> > &matrix) {
    	int m = matrix.size();
    	int n = matrix[0].size();
    	int column[n];
    	for(int i = 0; i < n; i++)
    		column[i] = 1;
    	int row[m];
    	for(int i = 0; i < m; i++)
    		row[i] = 1;
    	//cout << " m: "<<m<<endl;
    	//cout << "n: " << n << endl;
        for(int i = 0; i < m; ++i){
        	for(int j = 0; j < n; ++j){
        		//cout << matrix[i][j] << " ";
        		if(matrix[i][j] == 0){
        			column[j] = 0;
        			row[i] = 0;
        		}
        	}
        	//cout << endl;
        }
        /*
        for(int i = 0; i < n; ++i)
        	cout << column[i] << " " <<endl;
        for(int i = 0; i < m; ++i)
        	cout << row[i] <<" " << endl;
        	*/
        for(int i = 0; i < n; ++i){
        	if(column[i] == 0){
        		for(int j =0; j < m; ++j){
        			matrix[j][i] = 0;
        		}
        	}
        }
        for(int i = 0; i < m; ++i){
        	if(row[i] == 0){
        		for(int j =0; j < n; ++j){
        			matrix[i][j] = 0;
        		}
        	}
        }
        /*
        for(int i = 0; i < m; ++i){
        	for(int j = 0; j < n; ++j){
        		cout << matrix[i][j] << " ";
        	}
        	cout << endl;
        }
        */
    }

};

int main(){
	/*
	vector<vector<int> > num = { {0,0,0,5},
								 {4,3,1,4},
								 {0,1,1,4},
								 {1,2,1,3},
								 {0,0,1,1} };
*/
	vector<vector<int> > num ={{1}};				 
	Solution solution;
	solution.setZeroes(num);
	return 0;
}