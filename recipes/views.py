from django.shortcuts import render
from .models import Recipe

def home(request):
    return render(request, 'recipes/home.html')

def search_recipes(request):
    if request.method == "POST":
        ingredients = request.POST.getlist('ingredients')
        recipes = Recipe.objects.filter(ingredients__name__in=ingredients).distinct()
        return render(request, 'recipes/search_results.html', {'recipes': recipes})
    return render(request, 'recipes/search.html')
