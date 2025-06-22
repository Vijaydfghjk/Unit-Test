package strings


func Reverse(s string)string{ 
        temp := ""
	 for i:= len(s)-1; i>=0;i--{

		   temp += string(s[i])
	 }
return temp
}