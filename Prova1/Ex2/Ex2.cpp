#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

bool compara(const pair<string, int>& a, const pair<string, int>& b) {
    return a.second > b.second;
}

int main() {
    int teste = 1;
    ifstream in("estagio.in");
    if (!in.is_open()) {
        cerr << "Erro ao abrir o arquivo!" << endl;
        exit(-1);
    }

    ofstream out("estagio.out");
    vector<pair<string, int>> alunos;

    int N, md;
    string codigo;

    while (in >> N && N != 0) {
        alunos.clear(); // limpa o vetor para cada turma
        for (int i = 0; i < N; i++) {
            in >> codigo >> md;
            alunos.push_back(make_pair(codigo, md));
        }

        sort(alunos.begin(), alunos.end(), compara);

        out << "Turma " << teste++ << endl;

        int maior = alunos.front().second;

        for (const auto& aluno : alunos) {
            if (aluno.second == maior) {
                out << aluno.first << " ";
            }
        }
        out << endl << endl;
    }
    in.close();
    out.close();
}
