from django.shortcuts import render, redirect
from .models import Recipe
from .forms import RecipeForm

def recipe_list(request):
    recipes = Recipe.objects.all()
    return render(request, 'recipeservice/recipe_list.html', {'recipes': recipes})

def recipe_detail(request, id):
    recipe = Recipe.objects.get(id=id)
    return render(request, 'recipeservice/recipe_detail.html', {'recipe': recipe})

def add_recipe(request):
    if request.method == 'POST':
        form = RecipeForm(request.POST)
        if form.is_valid():
            form.save()
            return redirect('recipe_list')
    else:
        form = RecipeForm()
    return render(request, 'recipeservice/recipe_form.html', {'form': form})

