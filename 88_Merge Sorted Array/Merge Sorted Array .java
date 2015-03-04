public class Solution {
    public static void merge(int A[], int m, int B[], int n) {
        if(n == 0){
        	return;
        }
    	if(m == 0){
        	System.arraycopy( B, 0, A, 0,n);
        	return;
        }
    	int i = 0;
    	int j = 0;
    	while(i < m && j < n){
    		if(A[i] <= B[j]){
    			i++;
    		}
    		else if(A[i] > B[j]){
    			for(int k = m; k > i; --k){
    				A[k] = A[k -1];
    			}
    			A[i] = B[j];
    			j++;
    			m++;
    		}
    		if(j == n){
    			return;
    		}
    		if(i == m){
    			System.out.println("come in");
    			System.arraycopy(B,j,A,m,n -j);
    			return;
    		}
    	}
    }
    public static void main(String[] args){
    	int[] arrayA = new int[100];
    	//arrayA[0] = 1;
    	int[] arrayB = new int[10];
    	arrayB[0] = 2;
    	merge(arrayA,0,arrayB,1);
    	for(int i : arrayA){
    		System.out.println( i );
    	}
    }
}