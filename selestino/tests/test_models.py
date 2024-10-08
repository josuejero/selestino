from django.test import TestCase
from recipeservice.models import Ingredient, Recipe, Review

class IngredientModelTest(TestCase):
    def test_string_representation(self):
        ingredient = Ingredient(name="Tomato", quantity="2 cups")
        self.assertEqual(str(ingredient), "Tomato")

class RecipeModelTest(TestCase):
    def test_string_representation(self):
        recipe = Recipe(title="Pasta")
        self.assertEqual(str(recipe), "Pasta")

    def test_total_time_calculation(self):
        
        recipe = Recipe.objects.create(
            title="Pasta",
            prep_time=10,
            cook_time=20,
            description="Tasty pasta",
            cuisine_type="Italian",
            difficulty="Medium"
        )
        
        recipe.save()
        self.assertEqual(recipe.total_time, 30)

    def test_default_values(self):
        recipe = Recipe(title="Pasta")
        self.assertEqual(recipe.cuisine_type, "Unknown")
        self.assertEqual(recipe.difficulty, "Easy")

class ModelStrMethodTests(TestCase):
    def test_ingredient_str(self):
        ingredient = Ingredient.objects.create(name="Tomato", quantity="2")
        self.assertEqual(str(ingredient), "Tomato")

    def test_recipe_str(self):
        recipe = Recipe.objects.create(title="Pasta", description="Tasty pasta")
        self.assertEqual(str(recipe), "Pasta")

    def test_review_str(self):
        recipe = Recipe.objects.create(title="Pasta", description="Tasty pasta")
        review = Review.objects.create(recipe=recipe, user="John Doe", rating=5)
        self.assertEqual(str(review), "John Doe's review of Pasta")
        
class ManyToManyRelationshipTests(TestCase):
    def test_add_ingredients_to_recipe(self):
        
        tomato = Ingredient.objects.create(name="Tomato", quantity="2")
        pasta = Ingredient.objects.create(name="Pasta", quantity="200g")

        
        recipe = Recipe.objects.create(title="Pasta")
        recipe.ingredients.add(tomato, pasta)

        
        recipe = Recipe.objects.get(id=recipe.id)
        self.assertIn(tomato, recipe.ingredients.all())
        self.assertIn(pasta, recipe.ingredients.all())
        

class ForeignKeyRelationshipTests(TestCase):
    def test_review_links_to_recipe(self):
        
        recipe = Recipe.objects.create(title="Pasta", description="Tasty pasta")

        
        review = Review.objects.create(recipe=recipe, user="John Doe", rating=5)

        
        self.assertEqual(review.recipe, recipe)
