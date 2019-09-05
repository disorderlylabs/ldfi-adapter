import requests
import threading
import hashlib
import sys
from itertools import chain, combinations

product = "http://192.168.99.106:31380/productpage"
product_md5 = "b34651e8a6df854429c4ed038c00cc4c"


def thread_func(spec):
    resp = requests.get(product, headers={'x-ebay-ldfi': spec})
    #print(resp.status_code)
    result = hashlib.md5(resp.text.encode('utf-8'))
    if result.hexdigest() != product_md5:
        print("%s %d MISMATCH %s  %s"  % (spec, resp.status_code, result.hexdigest(), product_md5))



services = ['details', 'reviews', 'ratings']

def powerset(iterable):
    "powerset([1,2,3]) --> () (1,) (2,) (3,) (1,2) (1,3) (2,3) (1,2,3)"
    s = list(iterable)
    return chain.from_iterable(combinations(s, r) for r in range(len(s)+1))

threads = []
for item in powerset(services):
    #print(item)
    killstr = ",".join(map(lambda x: "fail="+x, item))
    x = threading.Thread(target=thread_func, args=(killstr,))
    threads.append(x)
    x.start()


threadcnt = int(sys.argv[1])
#threadcnt = 200


for i in range(threadcnt):
    x = threading.Thread(target=thread_func, args=(None,))
    threads.append(x)
    x.start()


x = threading.Thread(target=thread_func, args=("after=ratings|ratings|E|500",))
threads.append(x)
x.start()


x = threading.Thread(target=thread_func, args=("after=ratings|ratings|T|6000",))
threads.append(x)
x.start()


for t in threads:
    t.join()


