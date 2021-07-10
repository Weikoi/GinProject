package main

type Option func(*Options)

type Options struct {
	Name string
	Age  int
}

var (
	DefaultName = "defaultName"
	DefaultAge  = 10
)

// 如果没有对指定成员变量初始化，就使用默认参数
func NewOptions(opts ...Option) *Options {

	// 初始化默认值
	options := &Options{
		Name: DefaultName,
		Age:  DefaultAge,
	}

	for _, o := range opts {
		o(options)
	}
	return options
}

// 初始化name
func Name(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

// 初始化age
func Age(age int) Option {
	return func(o *Options) {
		o.Age = age
	}
}
