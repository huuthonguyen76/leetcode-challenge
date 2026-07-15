# https://leetcode.com/problems/print-in-order/

import threading

from typing import Callable


class Foo:
    def __init__(self):
        self.lock_1 = threading.Lock()
        self.lock_2 = threading.Lock()
        self.lock_3 = threading.Lock()

        print("herw")
        self.lock_2.acquire()
        print("there")
        self.lock_2.acquire()
        print("dasd")
        self.lock_3.acquire()


    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()

        self.lock_2.release()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        with self.lock_2:
            # printSecond() outputs "second". Do not change or remove this line.
            printSecond()

            self.lock_3.release()


    def third(self, printThird: 'Callable[[], None]') -> None:
        with self.lock_3:
            # printThird() outputs "third". Do not change or remove this line.
            printThird()
        


def printFirst():
    print("first", end="")

def printSecond():
    print("second", end="")

def printThird():
    print("third", end="")

if __name__ == "__main__":
    foo = Foo()
    # The following test will start all three in random order to see if they are printed in order.
    output = []

    def my_printFirst():
        output.append("first")

    def my_printSecond():
        output.append("second")

    def my_printThird():
        output.append("third")

    t1 = threading.Thread(target=foo.first, args=(my_printFirst,))
    t2 = threading.Thread(target=foo.second, args=(my_printSecond,))
    t3 = threading.Thread(target=foo.third, args=(my_printThird,))

    # Start out of order
    t3.start()
    t2.start()
    t1.start()

    t1.join()
    t2.join()
    t3.join()

    assert output == ["first", "second", "third"], f"Output was {output}"
        

