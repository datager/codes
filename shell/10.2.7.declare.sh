#sum=100+300+500
#echo ${sum}
#
#declare -i sum=100+300+500
#echo ${sum}
#
#declare -x sum
#export | grep sum
#
##declare -r sum
##sum=tesgting
#
#echo "aaa" ${sum}
#declare +x sum
#declare -p sum
#export | grep sum
#echo "bbb" ${sum}

var[1]="small min"
var[2]="big min"
var[3]="nice min"
echo "${var[1]},${var[2]},${var[3]}"