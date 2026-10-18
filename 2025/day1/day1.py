#! /usr/bin/env python3
# left = -
# right = +


def rotate(transform: str, start: int) -> int:
    dir = transform[0]
    num = transform[1:]

    num = int(num)
    num = num * -1 if dir == 'L' else num * 1

    pos = start + num 
    if pos < 0:
        pos = 100 + pos
    elif pos > 99:
        pos = 100 - pos

    return pos


def main():
    start = 50
    transform = 'L68'

    pos = rotate(transform, start)
    print(pos)

if __name__ == "__main__":
    main()


