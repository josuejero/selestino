from django.test import TestCase
from django.urls import reverse
from recipeservice.models import Recipe

class RecipeViewTests(TestCase):
    def setUp(self):
        self.recipe = Recipe.objects.create(title="Pasta", prep_time=10, cook_time=15)

    def test_recipe_list_view(self):
        response = self.client.get(reverse('recipe_list'))
        self.assertEqual(response.status_code, 200)
        self.assertContains(response, "Pasta")

    def test_recipe_detail_view(self):
        response = self.client.get(reverse('recipe_detail', args=[self.recipe.id]))
        self.assertEqual(response.status_code, 200)
        self.assertContains(response, "Pasta")

    def test_add_recipe_view(self):
      response = self.client.post(reverse('add_recipe'), {
          'title': 'Salad',
          'description': 'Fresh salad',
          'prep_time': 5,
          'cook_time': 0,
          'instructions': 'Mix ingredients',
          'ingredients': [self.recipe.ingredients.create(name="Lettuce", quantity="1 cup").id],
          'difficulty': 'Easy',  # Add this line
          'cuisine_type': 'American',  # Add this line
      })
      if response.status_code != 302:
          print(response.context['form'].errors)
      self.assertEqual(response.status_code, 302)  # Redirect after successful form submission
