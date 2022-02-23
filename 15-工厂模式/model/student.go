package model

type student struct {
	Name  string
	score float64
}

//因为 student 结构体首字母是小写，因此只能子啊model使用
//我们通过工厂模式来解决
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

//如果score为小写，那么在其他包我们不可以访问，我们可以提供一个方法
func (s *student) GetScore() float64 {
	return s.score
}
