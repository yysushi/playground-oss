from selenium import webdriver
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By

# or webdriver.Chrome()
with webdriver.Remote(
        command_executor='http://selenium-hub:4444',
        desired_capabilities=DesiredCapabilities.CHROME) as driver:
    driver.get('http://www.python.org')
    assert "Python" in driver.title
    elem = driver.find_element(by=By.NAME, value='q')
    elem.clear()
    elem.send_keys('pycon')
    elem.send_keys(Keys.RETURN)
    assert "No results found." not in driver.page_source
