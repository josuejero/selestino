from django.contrib import admin
from django.apps import apps
from .models import Ingredient, Recipe, Review

class ReviewInline(admin.TabularInline):
    model = Review
    extra = 1

@admin.register(Ingredient)
class IngredientAdmin(admin.ModelAdmin):
    list_display = ['name', 'quantity']

@admin.register(Recipe)
class RecipeAdmin(admin.ModelAdmin):
    list_display = ['title', 'cuisine_type', 'difficulty', 'total_time']
    search_fields = ['title', 'cuisine_type']
    list_filter = ['difficulty', 'cuisine_type']
    readonly_fields = ['total_time']
    inlines = [ReviewInline]

@admin.register(Review)
class ReviewAdmin(admin.ModelAdmin):
    list_display = ['recipe', 'user', 'rating', 'created_at']
    list_filter = ['rating']
    search_fields = ['recipe__title', 'user']

models = apps.get_models()
for model in models:
    try:
        admin.site.register(model)
    except admin.sites.AlreadyRegistered:
        pass
