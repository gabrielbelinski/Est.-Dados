#include <iostream>
#include <fstream>
using namespace std;

int hanoi(int n, int orig, int dest, int temp) {
    int mv1, mv2, mv3;
    if (n == 1) {
        return 1;
    }
    else {
        mv1 = hanoi(n-1, orig, temp, dest);
        mv2 = hanoi(1, orig, dest, temp);
        mv3 = hanoi(n-1, temp, dest, orig);
        int tmp = mv1 + mv2 + mv3;
        return tmp;
    }
}

int main() {
    int N, teste = 1;
    ifstream in("hanoi.in");
    ofstream out("hanoi.out");
    while (in >> N && N != 0) {
        out << "Teste " << teste++ << endl;
        out << hanoi(N, 1, 3, 2) << endl;
        out << endl;
    }

    in.close();
    out.close();

    return 0;
}
