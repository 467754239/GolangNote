#!/usr/bin/env python

import random

unival = random.uniform(0, 10)
print "rpc.call_error %f" % unival
print "rpc.call_total %f" % random.randint(100, 200)
