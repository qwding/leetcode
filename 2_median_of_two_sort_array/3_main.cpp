#include <iostream>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(int A[], int m, int B[], int n) {
       
            
    }
};

/*
    http://blog.csdn.net/yutianzuijin/article/details/11499917/
    这个讲了迭代，这次编写总结几个问题，关于++a和a++与？：关系。
    ？之前的判断会在将该加的++a加上。在之后里的元素的a++里边，只对执行的语句进行加减。就是符合判断条件的：前或后语句。
    弄了好久，有点晕.....不想弄了
*/

int main(){
    int A[5] = {1,5,7,11,15};
    int B[3] = {2,6,12};
    Solution solution;
    double result = solution.findMedianSortedArrays(A,5,B,3);
    printf("value is :%f\n", result);
    return 0;


}