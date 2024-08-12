import pytest
from django.contrib import admin
from recipeservice.models import Ingredient, Recipe, Review

@pytest.mark.django_db
def test_admin_registration():
    assert Ingredient in admin.site._registry
    assert Recipe in admin.site._registry
    assert Review in admin.site._registry
