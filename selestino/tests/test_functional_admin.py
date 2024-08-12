from django.urls import reverse
from django.contrib.auth.models import User
from django.test import TestCase

class AdminPageAccessTest(TestCase):
    def setUp(self):
        self.user = User.objects.create_superuser('admin', '[email protected]', 'password')
        self.client.force_login(self.user)

    def test_recipe_changelist_page(self):
        url = reverse('admin:recipeservice_recipe_changelist')
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)

    def test_ingredient_changelist_page(self):
        url = reverse('admin:recipeservice_ingredient_changelist')
        response = self.client.get(url)
        self.assertEqual(response.status_code, 200)
