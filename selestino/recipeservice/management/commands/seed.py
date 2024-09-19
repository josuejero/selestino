from django.core.management.base import BaseCommand
from recipeservice.models import Ingredient, Recipe

class Command(BaseCommand):
    def handle(self, *args, **kwargs):
        ingredient1 = Ingredient.objects.create(name='Tomato', quantity='2')
        ingredient2 = Ingredient.objects.create(name='Onion', quantity='1')
        recipe = Recipe.objects.create(
            title='Tomato Soup', 
            description='A delicious tomato soup', 
            instructions='Boil the ingredients...',
            prep_time=10,  
            cook_time=20   
        )
        recipe.ingredients.add(ingredient1, ingredient2)
