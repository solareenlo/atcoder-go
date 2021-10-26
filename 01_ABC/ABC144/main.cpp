#include <bits/stdc++.h>
using namespace std;

int main(){
  int n , m;
cin >> n >> m;
  vector<int> g[n];
  for(int i = 0; i < m; i++){
int a, b;
    cin >> a >> b;
    a--, b--;
    g[a].push_back(b);
  }
  double ans=1e9;
  for(int u = 0; u < n; u++){
vector<double> dp(n);
  for(int i = n-1; i >= 0; i--){
    double s = 0 ,m = 0;
     int c = 0;
    for(int j: g[i]){
c++;
      s += dp[j]; 
      m  = max(m, dp[j]);
}
    if(!c){
dp [i]  = 0;
      continue;
    }
    
    if(u == i){
if(c > 1){
c--;
  s -= m;
}
}
dp[i] = s / c + 1;
}
      ans = min( ans, dp[0]);
  }
  printf("%.10lf\n", ans);
}
