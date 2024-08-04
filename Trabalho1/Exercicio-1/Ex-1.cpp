#include <iostream>
#include <fstream>
#include <cstdlib>
using namespace std;

class Sorteio{
    private:
        int N;
        int *vet;

    public:
        void setN(int N){
            this->N = N;
        }
        int getN(){
            return N;
        }
        void alocaVet(int N){
            vet = new int[N];
        }
        void setVet(int *vet){
            this->vet = vet;
        }
        int *getVet(){
            return vet;
        }
        void deleteVet(){
            delete [] vet;
        }
};

int main(void){

    Sorteio a;
    ifstream entrada; 
    ofstream saida;

    entrada.open("quermesse.in");
    saida.open("quermesse.out");

    int N;
    int teste = 1;

        if((entrada.is_open()) && (saida.is_open())){

            while((entrada >> N) && N != 0){
                saida << "Teste" << endl << teste++ << " ";
                a.setN(N);
                int valorN = a.getN();
                a.alocaVet(valorN);
                int *vet = a.getVet();
                a.setVet(vet);

                for(int i = 0; i < valorN; i++){
                    entrada >> vet[i];
                }
                
                for(int i = 0; i < valorN; i++){
                    if(vet[i] == i+1){
                        saida << vet[i] << " ";
                    }
                }

                saida << endl;
                a.deleteVet();
            }
        }
        else{
            cout << "Falha na abertura dos arquivos. Verifique o diretório e tente novamente." << endl;
            exit(-1);
        }

    entrada.close();
    saida.close();
    return 0;
}