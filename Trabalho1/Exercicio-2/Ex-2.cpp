#include <iostream>
#include <fstream>
#include <cstdlib>

using namespace std;

class Valores{
    private:
        int N;
        int vet[4];
    public:
        void setN(int N){
            this->N = N;
        }
        int getN(){
            return this->N;
        }
        void setVet(int N){
            this->vet[0] = N/50;
            N = N%50;
            this->vet[1] = N/10;
            N = N%10;
            this->vet[2] = N/5;
            N = N%5;
            this->vet[3] = N/1;
        }
        int getVet(int i){
            return this->vet[i];
        }
        void deleteVet(){
            for(int i = 0; i < 4; i++){
                this->vet[i] = 0;
            }
        }

};

int main(){

    ifstream fin;
    ofstream fout;
    Valores v;

    fin.open("bit.in");
    fout.open("bit.out");

    if((fin.is_open()) && (fout.is_open())){

        int entrada;
        int teste = 1;

        while((fin >> entrada) && entrada != 0){

            fout << "Teste " << teste++ << endl;

            v.setN(entrada);
            v.setVet(v.getN());

            for(int i = 0; i < 4; i++){
                
                fout << v.getVet(i);
                if(i < 4-1){
                    fout << ", ";
                }
            }

            fout << endl << endl;

            v.deleteVet();
            
        }
    }
    else{
        cout << "Não foi possivel abrir os arquivos. Verifique os diretorios." << endl;
        exit(-1);
    }

    fin.close();
    fout.close();
    return 0;

}