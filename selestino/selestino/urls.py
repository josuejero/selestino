from django.urls import path
from . import views
import logging

logger = logging.getLogger(__name__)

urlpatterns = [
    path('recipes/', views.recipe_list, name='recipe_list'),
    path('recipes/<int:id>/', views.recipe_detail, name='recipe_detail'),
    path('recipes/add/', views.add_recipe, name='add_recipe'),
    path('recipes/<int:id>/edit/', views.edit_recipe, name='edit_recipe'),
    path('recipes/<int:id>/delete/', views.delete_recipe, name='delete_recipe'),
    path('recipes/<int:recipe_id>/reviews/', views.add_review, name='add_review'),
]
