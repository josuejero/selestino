from django.test import TestCase
from django.urls import reverse
from recipeservice.models import Recipe

class RecipeViewTests(TestCase):
    def setUp(self):
        self.recipe = Recipe.objects.create(title="Pasta", prep_time=10, cook_time=15)

    def test_recipe_list_view(self):
        response = self.client.get(reverse('recipe_list'), follow=True)
        print(f"Recipe list view status code: {response.status_code}")
        self.assertEqual(response.status_code, 200)

    def test_recipe_detail_view(self):
        response = self.client.get(reverse('recipe_detail', kwargs={'id': self.recipe.id}), follow=True)
        print(f"Recipe detail view status code: {response.status_code}")
        self.assertEqual(response.status_code, 200)

    def test_add_recipe_view(self):
        response = self.client.get(reverse('add_recipe'), follow=True)
        print(f"Add recipe view status code: {response.status_code}")
        self.assertEqual(response.status_code, 200)
