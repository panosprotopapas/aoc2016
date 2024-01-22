import re

print(
    "Part 1:",
    sum(
        list(
            map(
                lambda x: True
                if x[0] + x[1] > x[2] and x[0] + x[2] > x[1] and x[1] + x[2] > x[0]
                else False,
                [
                    [int(i) for i in re.findall(r"\d+", l.strip())]
                    for l in open("./day03/input.txt", "r").readlines()
                ],
            )
        )
    ),
)

print(
    "Part 2:",
    [
        sum(
            [
                sum(
                    [
                        True
                        if (
                            input[i][j] + input[i + 1][j] > input[i + 2][j]
                            and input[i][j] + input[i + 2][j] > input[i + 1][j]
                            and input[i + 1][j] + input[i + 2][j] > input[i][j]
                        )
                        else False
                        for j in range(3)
                    ]
                )
                for i in range(
                    0,
                    len(input),
                    3,
                )
            ]
        )
        for input in [
            [
                [int(i) for i in re.findall(r"\d+", l.strip())]
                for l in open("./day03/input.txt", "r").readlines()
            ]
        ]
    ][0],
)
