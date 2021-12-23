package pool

type Work func(input Input) Output

type Input interface {
}

type Output interface {
}
