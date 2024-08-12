from django.test import TestCase
from django.contrib import admin
from recipeservice.models import Ingredient, Recipe, Review
from recipeservice.admin import IngredientAdmin, RecipeAdmin, ReviewAdmin, ReviewInline


class AdminSiteTest(TestCase):
    def test_models_registered(self):
        self.assertIn(Ingredient, admin.site._registry)
        self.assertIn(Recipe, admin.site._registry)
        self.assertIn(Review, admin.site._registry)

    def test_custom_admin_class(self):
        self.assertIsInstance(admin.site._registry[Ingredient], IngredientAdmin)
        self.assertIsInstance(admin.site._registry[Recipe], RecipeAdmin)
        self.assertIsInstance(admin.site._registry[Review], ReviewAdmin)

class RecipeAdminTest(TestCase):
    def test_list_display(self):
        admin_class = admin.site._registry[Recipe]
        self.assertEqual(admin_class.list_display, ['title', 'cuisine_type', 'difficulty', 'total_time'])

    def test_search_fields(self):
        admin_class = admin.site._registry[Recipe]
        self.assertEqual(admin_class.search_fields, ['title', 'cuisine_type'])

    def test_list_filter(self):
        admin_class = admin.site._registry[Recipe]
        self.assertEqual(admin_class.list_filter, ['difficulty', 'cuisine_type'])

class ReviewInlineTest(TestCase):
    def test_inline(self):
        admin_class = admin.site._registry[Recipe]
        print(admin_class.inlines)  # Add this line for debugging
        self.assertTrue(ReviewInline in admin_class.inlines)
