#HEADERS="-H x-ebay-ldfi:fail=details"
HEADERS="-H x-ebay-ldfi:fail=ratings"
#HEADERS=""

#HEADERS="-H x-ebay-ldfi:fail=ratings,fail=details"

args="-H x-ebay-ldfi:"
for var in "$@"; do
    args="$args,fail=$var"
done


FILE=unit.$$.out


echo args $args

curl -vvv $args http://192.168.99.103:31380/productpage > $FILE


# check if call to ratings succeeded:
grep glyphicon-star $FILE > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS ratings
else
    echo NO RATINGS
fi

grep English $FILE > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS details
else
    echo NO DETAILS
fi

grep refreshing $FILE > /dev/null 2>&1
if [ "$?" -eq "0" ]; then
    echo PASS reviews
else
    echo NO REVIEWS
fi

rm $FILE
