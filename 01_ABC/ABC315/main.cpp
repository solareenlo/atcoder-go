#include "atcoder/convolution"
#include <bits/stdc++.h>
using namespace std;
using Z = atcoder::modint998244353;
using atcoder::convolution;
const int N = 200005;
Z f[N], g[N];
int n, a[N];
void solve(int l, int r)
{
	if (l == r) {
		if (l) {
			g[l] = g[l - 1] + f[l];
			f[l] = g[l] * a[l];
		}
		return;
	}
	int mid = (l + r) >> 1;
	solve(l, mid);
	if (!l) {
		vector<Z> tt = convolution(vector<Z>(f, f + mid + 1), vector<Z>(f + l, f + mid + 1));
		for (int i = mid + 1; i <= r; i++) {
			f[i] += tt[i - 1];
		}
	} else {
		vector<Z> tt = convolution(vector<Z>(f, f + r - l + 1), vector<Z>(f + l, f + mid + 1));
		for (int i = mid + 1; i <= r; i++) {
			f[i] += tt[i - l - 1] * 2;
		}
	}
	solve(mid + 1, r);
}
int main()
{
	ios::sync_with_stdio(false);
	cin >> n;
	f[0] = 1;
	for (int i = 1; i <= n; i++) {
		cin >> a[i];
	}
	solve(0, n);
	for (int i = 1; i <= n; i++) {
		cout << f[i].val() << ' ';
	}
	return 0;
}
