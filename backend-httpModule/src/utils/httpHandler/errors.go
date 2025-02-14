package httpHandler

import "fmt"

var DetailsError = `{
'error':'Error in Passed Data'
}`

var Success = func(message string) string {
	ans := fmt.Sprintf(`{
	'message':'Success Process'
	'details' :%s
	}`,message)
	return ans
}