#include <bits/stdc++.h>
using namespace std;
#define rep(i, n) for (int i = 0; i < (n); i++)
constexpr int MX = 100005;
int n, q, a[MX], s, l, r;
int main()
{
    scanf("%d%d", &n, &q);
    rep(i, n)
    {
        scanf("%d", a + i);
        if (i)
            a[i] = gcd(a[i], a[i - 1]);
    }
    rep(i, q)
    {
        scanf("%d", &s);
        l = -1;
        r = n;
        while (l + 1 < r) {
            int m = (l + r) / 2;
            (gcd(s, a[m]) == 1 ? r : l) = m;
        }
        printf("%d\n", r == n ? gcd(s, a[n - 1]) : r + 1);
    }
}
