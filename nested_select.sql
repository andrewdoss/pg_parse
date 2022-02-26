SELECT t.foo
FROM (SELECT foo from bar) t
WHERE foo > 1;