from django.shortcuts import render, redirect
from .models import Recipe
from .forms import RecipeForm
import logging
from django.http import HttpResponse

logger = logging.getLogger(__name__)

def recipe_list(request):
    logger.debug(f"Accessing recipe_list. Method: {request.method}, Path: {request.path}")
    recipes = Recipe.objects.all()
    return render(request, 'recipeservice/recipe_list.html', {'recipes': recipes})

def recipe_detail(request, id):
    logger.debug(f"Accessing recipe_detail for ID {id}. Method: {request.method}, Path: {request.path}")
    recipe = Recipe.objects.get(id=id)
    return render(request, 'recipeservice/recipe_detail.html', {'recipe': recipe})

def add_recipe(request):
    logger.debug(f"Accessing add_recipe. Method: {request.method}, Path: {request.path}")
    form = RecipeForm()
    if request.method == 'POST':
        form = RecipeForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('recipe_list')
    return render(request, 'recipeservice/recipe_form.html', {'form': form})


