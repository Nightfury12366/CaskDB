Caskdb is a fast, embeddable and easy to maintain K-V database engine, which is based on bitcask model and implemented
by golang. It currently supports five data structures: string, list, hash, set and Zset. Support client connection.

The data types and command operations supported by caskdb are as follows:

- String
    - Set
    - MSet
    - SetNx
    - MSetNx
    - Get
    - MGet
    - GetSet
    - Remove
    - SLen

- Hash
    - HSet
    - HSetNx
    - HGet
    - HGetAll
    - HDel
    - HLen
    - HExist

- List
    - LPush
    - LRPush
    - LPop
    - LRPop
    - LInsert
    - LRInsert
    - LSet
    - LRem
    - LLen
    - LIndex
    - LRange
    - LExist

- Set
    - SAdd
    - SRem
    - SMove
    - SUnion
    - SDiff
    - SScan
    - SCard
    - SIsMember

- ZSet
    - ZAdd
    - ZRem
    - ZScoreRange
    - ZScore
    - ZCard
    - ZIsMember