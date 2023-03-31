#include <iostream>

int main() {
    int n;
    std::cin >> n;

    int a[n];
    for (int i = 0; i < n; i++) {
        std::cin >> a[i];
        a[i]--;
    }

    int c, k, res = n/2;
    for (int i = 0; i < n; i++) {
        c = 1;
        k = a[i];
        while (k != i) {
            k = a[k];
            c++;
        }
        if (c%2 == 1) {
            res = -1;
            break;
        }
    }
    std::cout << res << std::endl;
    return 0;
}
