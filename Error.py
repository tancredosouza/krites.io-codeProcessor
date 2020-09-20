from enum import Enum


class Error(Enum):
    NO_ERROR = 0
    COMPILE = 1
    EXECUTION = 2
    TIMEOUT = 3
    WRONG_ANSWER = 4
