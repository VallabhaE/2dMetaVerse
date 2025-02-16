package httpHandler

import "fmt"

var DetailsError = `{
'error':'Error in Passed Data'
}`

var Success = func(message string) string {
	ans := fmt.Sprintf(`{
	'message':'Success Process'
	'details' :%s
	}`, message)
	return ans
}

var SendMapDetailsToUser = func(space string, spaceElements string, Elements string) string {
	ans := fmt.Sprintf(`{
		'message':'Success Process'
		'details' :{
			'space': %s,
			'spaceElements':%s,
			'Elements':%s
		}
		}`, space,spaceElements,Elements)


	return ans
}
