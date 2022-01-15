#include<iostream>

void updateMatrix(int matrix[][3], int m,int n) {
    bool r1 = false;
    bool c1 = false;

    for(int i=0;i<m;i++) {
        if(matrix[i][0] == 0) {
            r1 = true;
            break;
        }
    }
    for(int i=0;i<n;i++) {
        if(matrix[0][i] == 0) {
            c1 = true;
            break;
        }
    }

    for(int i=0;i<m;i++) {
        r1 = false;
        for(int j=0;j<n;j++) {
            if(matrix[i][j] == 0) {
                r1 = true;
                break;
            }
        }
        matrix[i][0] = r1 ? 0 : matrix[i][0];
    }

    for(int i=0;i<n;i++) {
        c1 = false;
        for(int j=0;j<m;j++) {
            if(matrix[i][j] == 0) {
                c1 = true;
                break;
            }
        }
        matrix[0][i] = c1 ? 0 : matrix[0][i];
    }

    for(int i=0;i<m;i++) {
        for(int j=0;j<n;j++) {
            if(matrix[i][0] == 0 || matrix[0][j] == 0) {
                matrix[i][j] = 0;
            }
        }
    }
}

void printMatrix(int matrix[][3], int m, int n) {
    for(int i=0;i<m;i++) {
        for(int j=0;j<n;j++)
            std::cout << matrix[i][j] << "\t";
        std::cout<<std::endl;
    }
}

int main() {
    std::cout << "hello, world" << std:: endl;
    int matrix[3][3] = {
        {1, 2, 3},
        {4, 0, 5},
        {6, 7, 8}
    };
    updateMatrix(matrix, 3, 3);
    printMatrix(matrix, 3, 3);
    return 0;
}