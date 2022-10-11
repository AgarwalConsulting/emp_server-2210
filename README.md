# ReSTful API {REpresentational State Transfer}

CRUD {Create, Read, Update, Destroy} => Resources

HTTP Methods {POST, GET, PUT, DELETE, OPTIONS, ...}

## Employee Management Server (JSON API)

```
CRUD         Action            METHOD               URI                         Req body                  Res Body
-----------------------------------------------------------------------------------------------------------------------------
Read         Index              GET              /employees                       -                      [{...}, {...}, ...]
Read         Show               GET              /employees/{id}                  -                       {...}
Create       Create            POST              /employees                     {...}                     {id: , ...}
Update       Update             PUT              /employees/{id}                {...}                     {...}
Update       Update            PATCH             /employees/{id}             {selected attrs}             {...}
Destroy      Destroy           DELETE            /employees/{id}                  -                       {...} / - (204)
```
