class Solution {
public:
    double findMedianSortedArrays(int A[], int m, int B[], int n) {
        double C[m+n];
        int a_idx = 0;
        int b_idx = 0;
        for(int i = 0; i < m+n; ++i){
            if (A[a_idx] <= B[b_idx]){
                C[i] = A[a_idx];
                a_idx++;
                if(a_idx >= m){
                    while(b_idx < n){
                        ++i;
                        C[i] = B[b_idx];
                        b_idx++;
                    }
                    break;
                }
            }
            else if(A[a_idx] > B[b_idx]){
                C[i] = B[b_idx];
                b_idx++;
                if(b_idx >= n ){
                    while(a_idx < m){
                        ++i;
                        C[i] = A[a_idx];
                        a_idx++;
                    }
                    break;
                }
            }

        }

        for(int i = 0; i < m+n ; ++i)
        	printf("order_c is %f\n",C[i]);
        if ((m+n)%2 == 1){
        	return (double)C[(m+n)/2];
        }
        else{

        	return (C[(m+n)/2] + C[(m+n)/2 - 1])/2.0;
        }
    }
};