from django import forms
from .models import Recipe, Ingredient, Review  

class RecipeForm(forms.ModelForm):
    class Meta:
        model = Recipe
        fields = ['title', 'description', 'ingredients', 'instructions', 'prep_time', 'cook_time', 'cuisine_type', 'difficulty']
        widgets = {
            'ingredients': forms.CheckboxSelectMultiple(attrs={'class': 'form-control'}),
        }

class IngredientForm(forms.ModelForm):
    class Meta:
        model = Ingredient
        fields = ['name', 'quantity']

class ReviewForm(forms.ModelForm):
    class Meta:
        model = Review
        fields = ['user', 'rating', 'comment']
