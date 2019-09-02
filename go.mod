module github.com/pangkunyi/plum

go 1.12

require (
	github.com/disintegration/imaging v1.6.1
	golang.org/x/text v0.3.2
)

replace (
	golang.org/x/image v0.0.0-20180708004352-c73c2afc3b81 => github.com/golang/image v0.0.0-20180708004352-c73c2afc3b81
	golang.org/x/text v0.3.2 => github.com/golang/text v0.3.2
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e => github.com/golang/tools v0.0.0-20180917221912-90fa682c2a6e
)
