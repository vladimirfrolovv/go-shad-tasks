//go:build !solution

package hogwarts

func GetCourseList(prereqs map[string][]string) []string {
	visit, cycle := make(map[string]struct{}), make(map[string]struct{})
	res := make([]string, 0, 10)
	for k := range prereqs {
		if !dfs(k, prereqs, visit, cycle, &res) {
			panic("Cycle!!!")
		}
	}
	return res
}
func dfs(crs string, prereqs map[string][]string, visit map[string]struct{}, cycle map[string]struct{}, res *[]string) bool {
	if _, ok := cycle[crs]; ok {
		return false
	}
	if _, ok := visit[crs]; ok {
		return true
	}
	cycle[crs] = struct{}{}
	for i := 0; i < len(prereqs[crs]); i++ {
		if !dfs(prereqs[crs][i], prereqs, visit, cycle, res) {
			return false
		}
	}
	delete(cycle, crs)
	*res = append(*res, crs)
	visit[crs] = struct{}{}
	return true
}
