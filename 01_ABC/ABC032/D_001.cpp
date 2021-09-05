#include <iostream>
#include <vector>
#include <map>

int main() {
    int n, ww;
    std::cin >> n >> ww;
    std::vector<int> v(n), w(n);
    for (int i=0; i<n; i++) {
        std::cin >> v[i] >> w[i];
    }
    std::map<int, int64_t> dp;
    dp[0] = 0;

    for (int i=0; i<n; i++) {
        std::map<int, int64_t> tmp;
        for (auto& j : dp) {
            if ((int64_t)j.first + w[i] <= ww) {
                tmp[j.first+w[i]] = std::max(tmp[j.first+w[i]], (int64_t)j.second+v[i]);
            } else {
                break;
            }
        }
        for (auto& j : tmp) {
            dp[j.first] = std::max(dp[j.first], j.second);
        }
        auto itr = dp.begin();
        int64_t m = -1;
        while (itr != dp.end()) {
            if ((*itr).second <= m) {
                dp.erase(itr++);
            } else {
                m = (*itr).second;
                itr++;
            }
        }
    }

    int64_t res = 0;
    for (auto& i : dp) {
        res = std::max(res, i.second);
    }
    std::cout << res << std::endl;
    return 0;
}
