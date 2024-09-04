from django.urls import reverse
from django.contrib.auth.models import User
from django.test import TestCase

class AdminPageAccessTest(TestCase):
    def setUp(self):
        # Create a superuser
        self.admin_user = User.objects.create_superuser('admin', 'admin@example.com', 'password')
        self.client.login(username='admin', password='password')

    def test_recipe_changelist_page(self):
        response = self.client.get(reverse('admin:recipeservice_recipe_changelist'), follow=True)
        print(f"Recipe changelist status code: {response.status_code}, final destination: {response.redirect_chain}")
        self.assertEqual(response.status_code, 200)

    def test_ingredient_changelist_page(self):
        response = self.client.get(reverse('admin:recipeservice_ingredient_changelist'), follow=True)
        print(f"Ingredient changelist status code: {response.status_code}, final destination: {response.redirect_chain}")
        self.assertEqual(response.status_code, 200)
