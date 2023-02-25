### tcp

#### Conn
```md
Client          Server
        syn
------------------>

        syn-ack
<------------------

        ack
------------------>
```

#### Closed Port
```md
Client          Server
        syn
-------------------->

        rst
<-------------------
```

#### Filtered Port
```md
Client          Server
        syn
-------------------->

        Timeout
```