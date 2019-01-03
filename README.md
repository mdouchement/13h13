# 13h13

**routes_exporter** for Prometheus. This a proxy that count number of requests for a given route and HTTP status code.

```
# HELP http_route_count Number of requests for a given route and HTTP status code.
# TYPE http_route_count counter
http_route_count{route="/",service="shiori",status="301"} 1
http_route_count{route="/css/fontawesome.css",service="shiori",status="200"} 2
http_route_count{route="/css/fonts/fa-regular-400.woff2",service="shiori",status="200"} 1
http_route_count{route="/css/fonts/source-sans-pro-v11-latin-300.woff2",service="shiori",status="200"} 1
http_route_count{route="/css/fonts/source-sans-pro-v11-latin-600.woff2",service="shiori",status="200"} 1
http_route_count{route="/css/fonts/source-sans-pro-v11-latin-regular.woff2",service="shiori",status="200"} 1
http_route_count{route="/css/source-sans-pro.css",service="shiori",status="200"} 2
http_route_count{route="/css/stylesheet.css",service="shiori",status="200"} 2
http_route_count{route="/js/axios.js",service="shiori",status="200"} 2
http_route_count{route="/js/js-cookie.js",service="shiori",status="200"} 2
http_route_count{route="/js/vue.js",service="shiori",status="200"} 2
http_route_count{route="/login",service="shiori",status="200"} 2
```

## Usage

Take a look to the `docker-compose.yml` file.

## License

**MIT**
