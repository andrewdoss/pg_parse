# pg_parse

Thin wrapper of [pg_query_go](https://github.com/pganalyze/pg_query_go) for convenient command line access to Postgres query parse trees.

## Usage

Either enter a SQL query string as a single command:

```bash
go run main.go "SELECT 1;"
```

 Or pipe in on stdin. 

```bash
go run main.go < query_string.sql
```

Example: Simple Select

```sql
SELECT foo, bar
FROM baz
WHERE foo > 1;
```

```bash
go run main.go < simple_select.sql | jq
```

Output: (very verbose, you can use jq to filter arbitrarily for clarity)

```json
"SelectStmt": {
    "targetList": [
    {
        "ResTarget": {
        "val": {
            "ColumnRef": {
            "fields": [
                {
                "String": {
                    "str": "foo"
                }
                }
            ],
            "location": 7
            }
        },
        "location": 7
        }
    }
    ],
    "fromClause": [
    {
        "RangeVar": {
        "relname": "bar",
        "inh": true,
        "relpersistence": "p",
        "location": 16
        }
    }
    ],
    "whereClause": {
    "A_Expr": {
        "kind": 0,
        "name": [
        {
            "String": {
            "str": ">"
            }
        }
        ],
        "lexpr": {
        "ColumnRef": {
            "fields": [
            {
                "String": {
                "str": "foo"
                }
            }
            ],
            "location": 26
        }
        },
        "rexpr": {
        "A_Const": {
            "val": {
            "Integer": {
                "ival": 1
            }
            },
            "location": 32
        }
        },
        "location": 30
    }
    },
    "op": 0
}
```

Example: Nested Select (emphasizes tree representation and composability of parsed statements)

```sql
SELECT t.foo
FROM (SELECT foo from bar) t
WHERE foo > 1;
```

```bash
go run main.go < nested_select.sql | jq
```

Output: 
```json
"SelectStmt": {
    "targetList": [
    {
        "ResTarget": {
        "val": {
            "ColumnRef": {
            "fields": [
                {
                "String": {
                    "str": "t"
                }
                },
                {
                "String": {
                    "str": "foo"
                }
                }
            ],
            "location": 7
            }
        },
        "location": 7
        }
    }
    ],
    "fromClause": [
    {
        "RangeSubselect": {
        "subquery": {
            "SelectStmt": {
            "targetList": [
                {
                "ResTarget": {
                    "val": {
                    "ColumnRef": {
                        "fields": [
                        {
                            "String": {
                            "str": "foo"
                            }
                        }
                        ],
                        "location": 26
                    }
                    },
                    "location": 26
                }
                }
            ],
            "fromClause": [
                {
                "RangeVar": {
                    "relname": "bar",
                    "inh": true,
                    "relpersistence": "p",
                    "location": 35
                }
                }
            ],
            "op": 0
            }
        },
        "alias": {
            "Alias": {
            "aliasname": "t"
            }
        }
        }
    }
    ],
    "whereClause": {
    "A_Expr": {
        "kind": 0,
        "name": [
        {
            "String": {
            "str": ">"
            }
        }
        ],
        "lexpr": {
        "ColumnRef": {
            "fields": [
            {
                "String": {
                "str": "foo"
                }
            }
            ],
            "location": 48
        }
        },
        "rexpr": {
        "A_Const": {
            "val": {
            "Integer": {
                "ival": 1
            }
            },
            "location": 54
        }
        },
        "location": 52
    }
    },
    "op": 0
}
```