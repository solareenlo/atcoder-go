package main

import "fmt"

func main() {
    in := bufio.NewReader(os.Stdin)

    var N int
    fmt.Fscan(in,&N)
    ps:=make([]Point3D,N)
    for i := 0; i < N; i++ {
        fmt.Fscan(in,&ps[i].x[0],&ps[i].x[1],&ps[i].x[2])
    }

    var cnv Convex3D
    cnv.init(ps)

    lis:=make(map[Point3D][]Point3D)
    for i := 0; i < len(cnv.fs); i++ {
        for j := 0; j < len(cnv.fs); j++ {
            no := cross(cnv.ss[i], cnv.ss[j]).normalize();
            if (no == Point3D()) {
                continue;
            }
            lis[no].push_back(cnv.ss[i]);
            lis[no].push_back(cnv.ss[j]);
        }
    }

    ans:=make(map[int][]Point3D)

    for _, uku :=range lis {
        no := uku.first;
        v := uku.second;
        sort(all(v));
        v.erase(unique(all(v)), v.end());

        var other Point3D
        for i := 0; i < len(cnv.fs); i++ {
            if (dot(no, cnv.ss[i]) == 0) {
                continue;
            }
            if (dot(no, cnv.ss[i]) > 0) {
                other += cnv.ss[i];
            } else {
                other -= cnv.ss[i];
            }
        }

        for i := 1; i < len(v); i++ {
            no2 := cross(v[0], v[i]).normalize();
            if (no != no2) {
                v[i] *= -1;
            }
        }
        for i := 1; i < len(v); i++ {
            for j := i + 1; j < len(v); j++ {
                no2 := cross(v[i], v[j]).normalize();
                if (no != no2) {
                    swap(v[i], v[j]);
                }
            }
        }

        for i := 1; i <= len(v); i++ {
            var po, ne Point3D
            for j := 0; j < i; j++ {
                po += v[j];
            }
            for j := i; j < len(v); j++ {
                ne += v[j];
            }
            al := other + po - ne;
            ans[al.norm()].push_back(al);
            al = other - po + ne;
            ans[al.norm()].push_back(al);
        }
    }

    keys := make([]int, 0, len(ans))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    k,v:=ans[len(ans)-1]
    // vector<Point3D> v = ans.rbegin()->second;
    sort(all(v));
    v.erase(unique(all(v)), v.end());

    fmt.Printf("%.20f\n", math.Sqrt(1.0 * float64(k)) * 0.25);
    fmt.Println(len(v))
}

type Point3D struct {
    x [3]int

};

// Point3D& operator+=(const Point3D& p)
// {
//     for (int i = 0; i < 3; i++)
//         x[i] += p.x[i];
//     return *this;
// }
// Point3D& operator-=(const Point3D& p)
// {
//     for (int i = 0; i < 3; i++)
//         x[i] -= p.x[i];
//     return *this;
// }
// Point3D& operator*=(int64_t k)
// {
//     for (int i = 0; i < 3; i++)
//         x[i] *= k;
//     return *this;
// }
// Point3D& operator/=(int64_t k)
// {
//     for (int i = 0; i < 3; i++)
//         x[i] /= k;
//     return *this;
// }
// Point3D operator+(const Point3D& p) const
// {
//     return Point3D(*this) += p;
// }
// Point3D operator-(const Point3D& p) const
// {
//     return Point3D(*this) -= p;
// }
// Point3D operator*(int64_t k) const
// {
//     return Point3D(*this) *= k;
// }
// Point3D operator/(int64_t k) const
// {
//     return Point3D(*this) /= k;
// }
// int64_t& operator[](int i)
// {
//     return x[i];
// }
// int64_t operator[](int i) const
// {
//     return x[i];
// }
//
// bool operator<(const Point3D& p) const
// {
//     return x < p.x;
// }
// bool operator==(const Point3D& p) const
// {
//     return x == p.x;
// }
// bool operator!=(const Point3D& p) const
// {
//     return x != p.x;
// }
//
// int64_t gcd(int64_t a, int64_t b)
// {
//     while (b) {
//         swap(a %= b, b);
//     }
//     return a;
// }
// Point3D normalize()
// {
//     Point3D p = *this;
//     if (p == Point3D())
//         return p;
//     int64_t g = gcd(abs(p[0]), gcd(abs(p[1]), abs(p[2])));
//     return p /= g;
// }
//
// int64_t norm()
// {
//     return x[0] * x[0] + x[1] * x[1] + x[2] * x[2];
// }

type Convex3D struct {
    fs [][]Point3D
    ss []Point3D
};

/*
inside: -1
surface: 0
outside: 1
*/
func (con *Convex3D)position(p Point3D) int {
    surface := false;
    for _, f :=range fs {
        no := cross(f[1] - f[0], f[2] - f[0]);
        d := dot(no, p - f[0]);
        if (d > 0) {
            return 1;
        }
        if (d == 0) {
            surface = true;
        }
    }
    if (surface) {
        return 0;
    }
    return -1;
}

func (con *Convex3D)add(p Point3D) {
    if (position(p) <= 0) {
        return;
    }
    nfs:=make([][]Point3D,0)
    for _, f :=range fs {
        no := cross(f[1] - f[0], f[2] - f[0]).normalize();
        d := dot(no, p - f[0]);
        if (d > 0) {
            continue;
        }
        if (d < 0) {
            nfs.push_back(f);
            continue;
        }

        flag:=make([]int,len(f))
        for i := 0; i < len(f); i++ {
            no2 := cross(f[(i + 1) % f.size()] - f[i], p - f[i]).normalize();
            flag[i] = no != no2;
        }
        nf:=make([]Point3D,0)
        for i := 0; i < len(f); i++ {
            if (flag[i] == 0) {
                nf.push_back(f[i]);
            } else if (flag[(i + f.size() - 1) % f.size()] == 0) {
                nf.push_back(f[i]);
                nf.push_back(p);
            }
        }
        nfs.push_back(nf);
    }
    fs = nfs;

    type pair struct {
        x, y Point3D
    }
    s:=make(map[pair]bool)
    for _, f :=range fs {
        for i := 0; i < len(f); i++ {
            pr:=pair{ f[i], f[(i + 1) % f.size()] };
            rv:=pair{ pr.second, pr.first };
            if (pr.first == p || pr.second == p) {
                continue;
            }
            if (s.find(rv) != s.end()) {
                s.erase(rv);
            } else {
                s.insert(pr);
            }
        }
    }
    for (auto& pr : s) {
        fs.push_back({ pr.second, pr.first, p });
    }
}

Convex3D(const vector<Point3D>& ps)
{
    vector<Point3D> tetra { ps[0] };
    for (int i = 0; i < (int)ps.size(); i++) {
        if (tetra[0] == ps[i])
            continue;
        tetra.push_back(ps[i]);
        break;
    }
    for (int i = 0; i < (int)ps.size(); i++) {
        if (cross(tetra[1] - tetra[0], ps[i] - tetra[0]) == Point3D())
            continue;
        tetra.push_back(ps[i]);
        break;
    }
    for (int i = 0; i < (int)ps.size(); i++) {
        Point3D no = cross(tetra[1] - tetra[0], tetra[2] - tetra[0]);
        int64_t d = dot(no, ps[i] - tetra[0]);
        if (d == 0)
            continue;
        if (d < 0)
            swap(tetra[1], tetra[2]);
        tetra.push_back(ps[i]);
        break;
    }

    fs.push_back({ tetra[0], tetra[2], tetra[1] });
    fs.push_back({ tetra[0], tetra[1], tetra[3] });
    fs.push_back({ tetra[1], tetra[2], tetra[3] });
    fs.push_back({ tetra[2], tetra[0], tetra[3] });

    for (int i = 0; i < (int)ps.size(); i++)
        add(ps[i]);
    for (auto& f : fs) {
        Point3D s;
        for (int i = 1; i + 1 < (int)f.size(); i++)
            s += cross(f[i] - f[0], f[i + 1] - f[0]);
        ss.push_back(s);
    }
}
