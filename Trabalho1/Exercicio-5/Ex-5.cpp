#include <iostream>
#include <fstream>
using namespace std;

class Sorveteiro{
    public:
        int U, V;
        Sorveteiro(int U, int V) {
            this->U = U;
            this->V = V;
        }
};

int main(void)
{
    ifstream fin("sorvete.in");
    ofstream fout("sorvete.out");

    int P, S, U, V, cont = 1;

    while(fin >> P >> S && (P != 0 && S != 0)){
        fout << "Teste " << cont++ << endl;
        Sorveteiro *sorveteiros[S];

        for(int i = 0; i < S; i++){
            fin >> U >> V;
            sorveteiros[i] = new Sorveteiro(U, V);
        }

        int inicio = -1, fim = -1;

        for(int i = 0; i <= P; i++){
            bool atendido = false;
                for(int j = 0; j < S; j++){
                    if(i >= sorveteiros[j]->U && i < sorveteiros[j]->V){
                        atendido = true;
                        break;
                    }
                }
                    if(atendido && inicio == -1){
                        inicio = i;
                    }   
                    else if(!atendido && inicio != -1){
                        fim = i;
                        fout << inicio << " " << fim << endl;
                        inicio = -1;
                        fim = -1;
                    }
        }

        if(inicio != -1){
            fout << inicio << " " << P << endl;
        }

        fout << endl;

        for(int i = 0; i < S; i++){
            delete sorveteiros[i];
        }
    }

    fin.close();
    fout.close();
}