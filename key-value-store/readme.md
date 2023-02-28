Requirements:
1. GET, SET, UNSET - retrieve/set/unset key accordingly

2. TRASACTIONS:
- 2A. Interpretation 1: I may have misinterpreted these requirements and I am still fuzzy on some of them.
    - BEGIN: beginning of a transaction, commnd stored in command cache
    - GET, SET, UNSET: commands executed and commands are stored in a command cache
    - ROLLBACK: Any commands till the occurance of the most recent BEGIN need to be undone, which means any commands between the second most recent BEGIN to the most recent BEGIN are re-executed (I'm still fuzzy on how far back we need to go in case of nested transactions). 
    - COMMIT: ?? Using the tmpCache implementation is a good idea
    - END: Delete all cached commands

- 2B. Interpretation 2: A simple implementation where you use to a temporary cache until a transaction is committed and you write to the permanent cache on COMMIT or END
    - Code: db-tmpcache.go
    - BEGIN: Beginning of a transaction, NO OP
    - GET: gets values from permanent KVStore
    - SET, UNSET: commands executed and results are written to a temporary cache `tmpCache`
    - ROLLBACK: Any commands since the most recent COMMIT are undone. This can be accomplished by maintaining a temporay cache map 'tmpCache'. Any commands are executed and values are added to temporary cache. This happens until a COMMIT or ROLLBACK is executed. Once a ROLLBACK is executed, the temporary cache is wiped out and only the changes since the last COMMIT, i.e., permanent KVStore is preserved. 
    - COMMIT: On commit, any values in the temporary cache gets merged to the original cache.
    - END: Delete any temporary cache values
