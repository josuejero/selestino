from django.test import TestCase
from django.urls import reverse
from .models import Recipe, Ingredient
from .forms import RecipeForm

class RecipeModelTest(TestCase):
    def setUp(self):
        self.ingredient1 = Ingredient.objects.create(name='Tomato', quantity='2')
        self.ingredient2 = Ingredient.objects.create(name='Onion', quantity='1')
        self.recipe = Recipe.objects.create(
            title='Tomato Soup',
            description='A delicious tomato soup',
            instructions='Boil the ingredients...',
            prep_time=10,
            cook_time=20
        )
        self.recipe.ingredients.add(self.ingredient1, self.ingredient2)

    def test_ingredient_creation(self):
        ingredient = Ingredient.objects.get(name='Tomato')
        self.assertEqual(ingredient.name, 'Tomato')

    def test_recipe_creation(self):
        self.assertEqual(self.recipe.title, 'Tomato Soup')
        self.assertEqual(self.recipe.total_time, 30)
        self.assertIn(self.ingredient1, self.recipe.ingredients.all())

    def test_recipe_list_view(self):
        response = self.client.get(reverse('recipe_list'))
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, 'recipeservice/recipe_list.html')
        self.assertContains(response, 'Tomato Soup')

    def test_recipe_detail_view(self):
        response = self.client.get(reverse('recipe_detail', args=[self.recipe.id]))
        self.assertEqual(response.status_code, 200)
        self.assertTemplateUsed(response, 'recipeservice/recipe_detail.html')
        self.assertContains(response, 'Tomato Soup')

    def test_add_recipe_form(self):
        form_data = {
            'title': 'Onion Soup',
            'description': 'A delicious onion soup',
            'ingredients': [self.ingredient2.id],
            'instructions': 'Boil the ingredients...',
            'prep_time': 10,
            'cook_time': 20,
            'difficulty': 'Easy',
            'cuisine_type': 'Unknown'
        }
        form = RecipeForm(data=form_data)
        self.assertTrue(form.is_valid())
        form.save()
        self.assertEqual(Recipe.objects.count(), 2)

    def test_recipe_missing_ingredient(self):
        form_data = {
            'title': 'Incomplete Soup',
            'description': 'Missing ingredients',
            'instructions': 'Boil the ingredients...',
            'prep_time': 10,
            'cook_time': 20
        }
        form = RecipeForm(data=form_data)
        self.assertFalse(form.is_valid())
