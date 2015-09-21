package cmd

import (
	"github.com/mabetle/mcore/mrun"
)

func AddFunc(fn func(), args ... string ){
	mrun.AddFunc(fn, args ... )
}

func AddRunner(r mrun.Runner, args ... string){
	mrun.AddRunner(r, args ... )
}

func Main(){
	mrun.Main()
}


