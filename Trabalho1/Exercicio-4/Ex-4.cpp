#include <iostream>
#include <fstream>
#include <cstdlib>
using namespace std;

class Meteoro{    
    private:
        int caiu;
    public:
    Meteoro(){
        caiu = 0;
    }
    int getCaiu(){
        return caiu;
    }
    void caiuDentro(int x, int y, int u, int v, int x1, int y1){
        if(x1 >= x && x1 <= u && y1 >= v && y1 <= y)
            caiu++;
    }

};

int main(void){
    ifstream entrada;
    ofstream saida;

    int N, X, Y, U, V, x1, y1, cont = 1;

    entrada.open("meteoros.in");
    saida.open("meteoros.out");

    if((entrada.is_open()) && (saida.is_open())){
        while(entrada >> X >> Y >> U >> V && (X!=0 || Y!=0 || U!=0 || V!=0)){
            saida << "Teste " << cont++ << endl;
            entrada >> N;
            Meteoro M;
         
        for(int i = 0; i < N; i++){ 
            entrada >> x1 >> y1;
            M.caiuDentro(X, Y, U, V, x1, y1);
        }
        
        saida << M.getCaiu() << endl << endl;
        }
    }
    else{
        cout << "Erro ao abrir/gravar arquivos!" << endl;
        exit(-1);
    }

    entrada.close();
    saida.close();
}