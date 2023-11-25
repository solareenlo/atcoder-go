package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var pi float64 = math.Atan(1.0) * 4.0

type vec3 struct {
	x, y, z float64
}

func (v vec3) mulF(c float64) vec3 {
	return vec3{v.x * c, v.y * c, v.z * c}
}

func (v vec3) divF(c float64) vec3 {
	return v.mulF(1.0 / c)
}

func (l vec3) plus(r vec3) vec3 {
	return vec3{l.x + r.x, l.y + r.y, l.z + r.z}
}

func (l vec3) minus(r vec3) vec3 {
	return vec3{l.x - r.x, l.y - r.y, l.z - r.z}
}

func (l vec3) mul(r vec3) vec3 {
	return vec3{l.x * r.x, l.y * r.y, l.z * r.z}
}

func (l vec3) dot(r vec3) float64 {
	return l.x*r.x + l.y*r.y + l.z*r.z
}

func (l vec3) cross(r vec3) vec3 {
	return vec3{l.y*r.z - l.z*r.y, l.z*r.x - l.x*r.z, l.x*r.y - l.y*r.x}
}

func (v vec3) len() float64 {
	return math.Sqrt(v.dot(v))
}

func (v *vec3) normalize() float64 {
	t := v.len()
	*v = v.divF(t)
	return t
}

func (l vec3) angle(r vec3) float64 {
	return math.Acos(l.dot(r) / (l.len() * r.len()))
}

type quat struct {
	w float64
	v vec3
}

func (l quat) mul(r quat) quat {
	tmp0 := r.v.mulF(l.w)
	tmp1 := l.v.mulF(r.w)
	tmp2 := l.v.cross(r.v)
	return quat{l.w*r.w - l.v.dot(r.v), tmp0.plus(tmp1).plus(tmp2)}
}

type Dat struct {
	angle float64
	imos  int
	id    int
}

func rot(p, axis vec3, angle float64) vec3 {
	ct := math.Cos(angle * 0.5)
	st := math.Sin(angle * 0.5)
	q := quat{ct, axis.mulF(st)}
	iq := quat{ct, axis.mulF(-st)}
	pp := quat{0.0, p}
	res := q.mul(pp).mul(iq)
	return res.v
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var p, p2, bp, cp, np [555]vec3
	var r2 [555]float64
	var graph [555][]Dat

	var solve func()
	solve = func() {
		var n, r int
		fmt.Fscan(in, &n, &r)
		for i := 0; i < n; i++ {
			var x, y, z float64
			fmt.Fscan(in, &x, &y, &z)
			p[i] = vec3{x, y, z}
		}
		pert_axis := vec3{2521.0, 98000.0, 9414.0}
		pert_axis.normalize()
		pert_angle := 123162.0
		q := quat{math.Cos(pert_angle * 0.5), pert_axis.mulF(math.Sin(pert_angle * 0.5))}
		iq := quat{math.Cos(pert_angle * 0.5), pert_axis.mulF(-math.Sin(pert_angle * 0.5))}
		for i := 0; i < n; i++ {
			p[i] = q.mul(quat{0.0, p[i]}).mul(iq).v
		}
		for i := 0; i < n; i++ {
			p2[i] = vec3{-p[i].x, -p[i].y, -p[i].z}
			Len := p2[i].normalize()
			r2[i] = math.Asin(float64(r) / Len)
		}
		valid := make([]bool, n)
		for i := range valid {
			valid[i] = true
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					dist := p2[i].angle(p2[j])
					if dist <= r2[j]-r2[i] {
						valid[i] = false
						break
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			if valid[i] {
				up := vec3{0.0, 1.0, 0.0}
				ax := up.cross(p2[i])
				ax.normalize()
				bp[i] = rot(p2[i], ax, r2[i])
			}
		}
		for i := 0; i < n; i++ {
			if valid[i] {
				ax := p2[i]
				pa := bp[i]
				pb := rot(pa, ax, 2.0*pi/3.0)
				pc := rot(pa, ax, -2.0*pi/3.0)
				cp[i] = pa.plus(pb).plus(pc).divF(3.0)
				np[i] = (pb.minus(pa)).cross(pc.minus(pa))
				np[i].normalize()
			}
		}
		UFp := NewDsu(n)
		for i := 0; i < n; i++ {
			graph[i] = make([]Dat, 0)
		}
		it := 0
		for i := 0; i < n; i++ {
			if valid[i] {
				var projQ func(int) (quat, quat)
				projQ = func(idx int) (quat, quat) {
					var q quat
					ez := vec3{0.0, 0.0, 1.0}
					ax := np[idx].cross(ez)
					ax.normalize()
					theta := np[idx].angle(ez)
					sth := math.Sin(theta / 2.0)
					cth := math.Cos(theta / 2.0)
					q.v = ax.mulF(sth)
					q.w = cth
					iq := quat{w: q.w, v: vec3{q.v.x, q.v.y, q.v.z}}
					iq.v = iq.v.mulF(-1.0)
					return q, iq
				}
				var rotQ func(quat, quat, vec3) vec3
				rotQ = func(q, iq quat, v vec3) vec3 {
					return (q.mul(quat{0.0, v}).mul(iq)).v
				}
				q, iq := projQ(i)

				for j := 0; j < i; j++ {
					if valid[j] && i != j {
						dist := p2[i].angle(p2[j])
						if math.IsNaN(dist) || dist > r2[i]+r2[j] {
							continue
						}
						UFp.Merge(i, j)
						crossline_v := np[i].cross(np[j])
						var crossline_p vec3
						crossline_p.z = 0.0
						npjx_npix := np[j].x / np[i].x
						crossline_p.y = (cp[j].dot(np[j]) - cp[i].dot(np[i])*npjx_npix) / (np[j].y - np[i].y*npjx_npix)
						crossline_p.x = (cp[i].dot(np[i]) - crossline_p.y*np[i].y) / np[i].x
						line_p := rotQ(q, iq, crossline_p.minus(cp[i]))
						line_v := rotQ(q, iq, crossline_v)
						radius := (bp[i].minus(cp[i])).len()
						other := rotQ(q, iq, cp[j].minus(cp[i]))

						ka := line_v.dot(line_v)
						kb := line_v.dot(line_p)
						kc := line_p.dot(line_p) - radius*radius
						kd := kb*kb - ka*kc
						kds := math.Sqrt(kd)
						k1 := (-kb + kds) / ka
						k2 := (-kb - kds) / ka
						cross1 := line_p.plus(line_v.mulF(k1))
						cross2 := line_p.plus(line_v.mulF(k2))
						angle1 := math.Atan2(cross1.y, cross1.x)
						angle2 := math.Atan2(cross2.y, cross2.x)
						angle_o := math.Atan2(other.y, other.x)
						if angle1 < 0.0 {
							angle1 += 2.0 * pi
						}
						if angle2 < 0.0 {
							angle2 += 2.0 * pi
						}
						if angle_o < 0.0 {
							angle_o += 2.0 * pi
						}
						if angle2 < angle1 {
							angle2, angle1 = angle1, angle2
							cross1, cross2 = cross2, cross1
							k1, k2 = k2, k1
						}
						if angle_o < angle1 || angle_o > angle2 {
							angle1 += 2.0 * pi
							angle2, angle1 = angle1, angle2
							cross1, cross2 = cross2, cross1
							k1, k2 = k2, k1
						}
						id1 := it
						it++
						id2 := it
						it++
						graph[i] = append(graph[i], Dat{angle1, +1, id1})
						graph[i] = append(graph[i], Dat{angle2, -1, id2})
						jq, jiq := projQ(j)
						crossj1 := rotQ(jq, jiq, rotQ(iq, q, cross1).plus(cp[i]).minus(cp[j]))
						crossj2 := rotQ(jq, jiq, rotQ(iq, q, cross2).plus(cp[i]).minus(cp[j]))
						other = rotQ(jq, jiq, cp[i].minus(cp[j]))
						angle1 = math.Atan2(crossj1.y, crossj1.x)
						angle2 = math.Atan2(crossj2.y, crossj2.x)
						angle_o = math.Atan2(other.y, other.x)
						if angle1 < 0.0 {
							angle1 += 2.0 * pi
						}
						if angle2 < 0.0 {
							angle2 += 2.0 * pi
						}
						if angle_o < 0.0 {
							angle_o += 2.0 * pi
						}
						if angle1 < angle2 {
							if angle_o < angle1 || angle_o > angle2 {
								graph[j] = append(graph[j], Dat{angle2, +1, id2})
								graph[j] = append(graph[j], Dat{angle1 + 2.0*pi, -1, id1})
							} else {
								graph[j] = append(graph[j], Dat{angle1, +1, id1})
								graph[j] = append(graph[j], Dat{angle2, -1, id2})
							}
						} else {
							if angle_o < angle2 || angle_o > angle1 {
								graph[j] = append(graph[j], Dat{angle1, +1, id1})
								graph[j] = append(graph[j], Dat{angle2 + 2.0*pi, -1, id2})
							} else {
								graph[j] = append(graph[j], Dat{angle2, +1, id2})
								graph[j] = append(graph[j], Dat{angle1, -1, id1})
							}
						}
					}
				}
			}
		}
		UFe := NewDsu(it)
		for i := 0; i < n; i++ {
			if valid[i] {
				m := len(graph[i])
				sort.Slice(graph[i], func(a, b int) bool {
					return graph[i][a].angle < graph[i][b].angle
				})
				for j := 0; j < m-1; j++ {
					graph[i][j+1].imos += graph[i][j].imos
				}
				half_sum := 0
				for j := 0; j < m; j++ {
					if graph[i][j].angle < 2.0*pi {
						half_sum = graph[i][j].imos
					}
				}
				sort.Slice(graph[i], func(a, b int) bool {
					return math.Mod(graph[i][a].angle, 2.0*pi) < math.Mod(graph[i][b].angle, 2.0*pi)
				})
				first_sum := 0
				for j := 0; j < m; j++ {
					if graph[i][j].angle < 2.0*pi {
						first_sum = graph[i][j].imos
						graph[i][j].imos += half_sum
					} else {
						half_sum = graph[i][j].imos
						graph[i][j].imos += first_sum
					}
				}
				for j := 0; j < m; j++ {
					if graph[i][j].imos == 0 {
						id1 := graph[i][j].id
						id2 := graph[i][(j+1)%m].id
						UFe.Merge(id1, id2)
					}
				}
			}
		}
		ans := 1
		for i := 0; i < n; i++ {
			if valid[i] && UFp.Leader(i) == i {
				if UFp.Size(i) > 1 {
					ans -= 1
				}
			}
		}
		for i := 0; i < it; i++ {
			if UFe.Leader(i) == i {
				if UFe.Size(i) > 1 {
					ans += 1
				}
			}
		}
		fmt.Fprintln(out, ans)
	}

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		solve()
	}
}

type dsu struct {
	n            int
	parentOrSize []int
}

func NewDsu(n int) *dsu {
	d := new(dsu)
	d.n = n
	d.parentOrSize = make([]int, d.n)
	for i := range d.parentOrSize {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *dsu) Merge(a, b int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *dsu) Same(a, b int) bool {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if !(0 <= b && b < d.n) {
		panic("")
	}
	return d.Leader(a) == d.Leader(b)
}

func (d *dsu) Leader(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *dsu) Size(a int) int {
	if !(0 <= a && a < d.n) {
		panic("")
	}
	return -d.parentOrSize[d.Leader(a)]
}

func (d *dsu) Groups() [][]int {
	leaderBuf := make([]int, d.n)
	groupSize := make([]int, d.n)
	for i := 0; i < d.n; i++ {
		leaderBuf[i] = d.Leader(i)
		groupSize[leaderBuf[i]]++
	}
	result := make([][]int, d.n)
	for i := 0; i < d.n; i++ {
		result[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < d.n; i++ {
		result[leaderBuf[i]] = append(result[leaderBuf[i]], i)
	}
	eraseEmpty := func(a [][]int) [][]int {
		result := make([][]int, 0, len(a))
		for i := range a {
			if len(a[i]) != 0 {
				result = append(result, a[i])
			}
		}
		return result
	}
	return eraseEmpty(result)
}
