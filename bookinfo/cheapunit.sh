#HEADERS="-H x-ebay-ldfi:fail=details"
HEADERS="-H x-ebay-ldfi:fail=ratings"
#HEADERS=""

curl -vvv $HEADERS http://192.168.99.103:31380/productpage > unit.out


# check if call to ratings succeeded:
grep glyphicon-star unit.out > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS ratings
else
    echo NO RATINGS
fi

grep English unit.out > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS details
else
    echo NO DETAILS
fi

grep refreshing unit.out > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS reviews
else
    echo NO REVIEWS
fi

