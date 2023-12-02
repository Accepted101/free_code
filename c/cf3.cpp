#include <iostream>
#include <queue>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    
    int t;
    scanf("%d",&t);
    while(t--) {
        vector<int> a;
        int n;
        scanf("%d",&n);
        for(int i = 0,j;i < n;++i) {
            scanf("%d",&j);
            a.push_back(j);
        }
        vector<int>b = a;
        sort(b.begin(),b.end());
        bool ok = true;
        int ans = -1;
        for(int i = 0;i < a.size();++i) {
            if(a[i] == b[0]) {
                for(int j = i+1;j < a.size();++j) {
                    if(a[j] < a[j-1]) {
                        ok = false;
                        break;
                    }
                }
                ans = i;
                break;
            }
        }
        if (ok) {
            printf("%d\n",ans);
        } else {
            printf("-1\n");
        }
    }
}
