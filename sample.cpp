#include <iostream>

using namespace std;

int fib(int x) {
    if (x < 2) return x+1;

    return fib(x -1) + fib(x-2);
}

int main() {
    int n;
    cin >> n;
    while (n--) {
        int x;
        cin >> x;
        cout << fib(x) << endl;
    }
    return 0;
}