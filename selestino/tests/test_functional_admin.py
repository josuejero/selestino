from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from django.contrib.staticfiles.testing import StaticLiveServerTestCase
from django.urls import reverse
from django.contrib.auth.models import User
import time

class AdminPageAccessTest(StaticLiveServerTestCase):
    def setUp(self):
        
        self.browser = webdriver.Chrome(executable_path='chromedriver')
        
        self.admin_user = User.objects.create_superuser('admin', 'admin@example.com', 'password')
        
        self.browser.get(f'{self.live_server_url}{reverse("admin:login")}')
        self.browser.find_element_by_name('username').send_keys('admin')
        self.browser.find_element_by_name('password').send_keys('password')
        self.browser.find_element_by_xpath('//input[@type="submit"]').click()

    def test_recipe_changelist_page(self):
        self.browser.get(f'{self.live_server_url}{reverse("admin:recipeservice_recipe_changelist")}')
        time.sleep(2)  
        self.assertIn("Select recipe to change", self.browser.page_source)

    def test_ingredient_changelist_page(self):
        self.browser.get(f'{self.live_server_url}{reverse("admin:recipeservice_ingredient_changelist")}')
        time.sleep(2)  
        self.assertIn("Select ingredient to change", self.browser.page_source)

    def tearDown(self):
        
        self.browser.quit()
