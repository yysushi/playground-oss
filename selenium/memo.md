# Note for Selenium

## Components

- browser
- web driver
- selenium

test client <-> selenium <-> web driver <-> browser <-> app

## TODO

- record
- grid
- capture screen
- pytest integration
- how to track python package and docker image mappings

https://www.selenium.dev/selenium/docs/api/py/api.html#webdriver-chrome

https://github.com/SeleniumHQ/selenium/tree/trunk/py

https://github.com/pytest-dev/pytest-selenium

https://github.com/SeleniumHQ/docker-selenium#debugging

https://github.com/danielkaiser/python-chromedriver-binary

## Command

```shell-session
$ docker-compose up --rm
$ docker build -t sample .
$ docker run -v $(pwd):/app --net selenium_default sample
```

`vnc://localhost:6900`
