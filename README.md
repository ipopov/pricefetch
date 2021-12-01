Building

```
docker build --target bin --output .  .
```

Crontab entry, to fetch data at 8pm every weekday:

```
0 20 * * 1-5 /path/to/.../bin/pricefetch --config /path/to/.../config.json >> /path/to/.../db
```
