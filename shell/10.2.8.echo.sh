##delete var
#MAIL=/var/spool/mail/dmtsai
#echo ${MAIL##/*/}
#echo ${MAIL%/*}
#
##replace var
#echo ${MAIL/a/A}
#echo ${MAIL//a/A}
#
#echo ${username}
#username=${username-root}
#echo ${username}
#
#username="vbird tsai"
#username=${username-root}
#echo ${username}

unset str
var=${str=newstr}
echo "var=${var}, str=${str}"