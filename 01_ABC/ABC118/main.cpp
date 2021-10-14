#include<bits/stdc++.h>
#define int long long
using namespace std;
void chmax(string &A,string B){if(A.size()<B.size()||A<B)A=B;}
int N,M,d[]={0,2,5,5,4,5,6,3,7,6};
string dp[10100];

signed main(){
  cin>>N>>M;
  vector<int> A(M);
  for(int &i:A)cin>>i;
  for(int i=0;i<N;i++){
    if(i&&dp[i]=="")
		continue;
    for(int j:A) {
		chmax(dp[i+d[j]],(char)(j+'0')+dp[i]);
	}
	cout << dp[i] << endl;
  }
  cout<<dp[N]<<endl;
}
