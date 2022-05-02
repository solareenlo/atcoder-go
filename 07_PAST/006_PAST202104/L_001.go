package main

func main() {
	var N,M int
	fmt.Scan(&N,&M)
	X  := make([]int, 50)
	Y  := make([]int, 50)
	for i := 0; i < N; i++ {
		fmt.Scan(&X[i],&Y[i])
  }
  CX  := make([]int, 8)
  CY  := make([]int, 8)
  R  := make([]int, 8)
  for i := 0; i < M; i++ {
	  fmt.Scan(&CX[i],&CY[i],&R[i])
  }

  ans := 1e150;
  for k := 0; k < 1 << M; k++ {
    vector<pair<double, pair<int, int>>> E;
    atcoder::dsu P(N + M);
    for (int i = 0; i < N; i++) {
      for (int j = i + 1; j < N; j++) {
        E.push_back( make_pair(hypot(X[i] - X[j], Y[i] - Y[j]), make_pair(i, j)));
      }
	}
    for (int j = 0; j < M; j++) {
      if (k >> j & 1) {
        for (int i = 0; i < N; i++) {
          double d = hypot(X[i] - CX[j], Y[i] - CY[j]) - R[j];
          if (d < 0) {
            d = -d;
		  }
          E.push_back(make_pair(d, make_pair(i, N + j)));
        }
	  }
	}
    for (int i = 0; i < M; i++) {
      if (k >> i & 1) {
        for (int j = i + 1; j < M; j++) {
          if (k >> j & 1) {
            double d = hypot(CX[i] - CX[j], CY[i] - CY[j]);
            if (d > R[i] + R[j]) {
              d -= R[i] + R[j];
			} else if (d < abs(R[i] - R[j])) {
              d = abs(R[i] - R[j]) - d;
			} else {
              d = 0;
			}
            E.push_back(make_pair(d, make_pair(N + i, N + j)));
          }
		}
	  }
	}
    sort(E.begin(), E.end());
    double now = 0;
    for (pair<double, pair<int, int>> e : E) {
      int u = e.second.first, v = e.second.second;
      if (!P.same(u, v)) {
        now += e.first;
	  }
      P.merge(u, v);
    }
    ans = min(ans, now);
  }
  cout << fixed << setprecision(16) << ans << endl;
  return 0;
}
