## admiral view group

view existing group

### Synopsis

view existing group by substring of group name or view all records when no argument passedpass the flag `-j,--json` to view the group in json structure with group variables

```
admiral view group ['group name'] [flags]
```

### Examples

```
admiral view group
admiral view group group1
admiral view group group1 -j
```

### Options

```
  -h, --help   help for group
  -j, --json   view in json format (present vars)
```

### SEE ALSO

* [admiral view](admiral_view.md)	 - view existing record

###### Auto generated by spf13/cobra on 10-Nov-2020