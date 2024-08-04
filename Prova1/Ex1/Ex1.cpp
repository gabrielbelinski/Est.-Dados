#include <iostream>
using namespace std;

int main(){
    
    int N, J, Z, teste = 1;

    while(cin >> N && N != 0){

        int v[N];
    
        for(int i = 0; i < N; i++){
            cin >> J >> Z;
            v[i] = J - Z;
            if(i > 0){
                v[i] = v[i] + v[i-1];
            }
        }

        cout << "Teste " << teste << endl;

        for(int i = 0; i < N; i++){
            cout << v[i] << endl;
        }
        teste++;
        cout << endl;
    }
}