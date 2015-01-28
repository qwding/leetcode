#include <iostream>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(int A[], int m, int B[], int n) {
        int k = m+n;
        if(k <= 4){
            printf("k little than %d\n", k);
        }
        else{
            k = (k + 1)/2;
            k = (k + 1)/2;
            int a_idx = -1;
            int b_idx = -1;
            while(k>=1){
                if(A[k + a_idx] <= B[k + b_idx]){
                    a_idx = a_idx + k;
                    printf("this circle is a :%f and k is%d\n", (double)A[a_idx],k);
                }
                else{
                    b_idx = b_idx + k;
                    printf("this circle is b : %f and k is%d\n", (double)B[b_idx],k);
                }
                if (k == 1)
                    break;
                k = (k + 1)/2;
            }
            if ((m+n)%2 == 1){
                return (A[a_idx + 1] < B[b_idx + 1]) ? A[a_idx + 1] : B[b_idx + 1];
            }
            else{
                printf("a_idx is : %d, b_idx is : %d\n", a_idx,b_idx);
                double temp_mid_little = (A[++a_idx] < B[++b_idx]) ? A[a_idx++] : B[b_idx++];
                 printf("a_idx is : %d, b_idx is : %d\n", a_idx,b_idx);
                double temp_mid_big = (A[a_idx] < B[b_idx]) ? A[a_idx ] : B[b_idx];
                 printf("a_idx is : %d, b_idx is : %d\n", a_idx,b_idx);
                printf("littel is %f,biger is %f\n", temp_mid_little,temp_mid_big);
                return (temp_mid_big +temp_mid_little)/2.0;
            }
        }
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